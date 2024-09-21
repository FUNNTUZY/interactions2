package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"interactions/api/proto"
	"interactions/internal/config"
	"interactions/internal/infrastructure/db"
	"interactions/internal/infrastructure/logger"
	"interactions/internal/infrastructure/metrics"
	"interactions/internal/infrastructure/middleware"
	"interactions/internal/infrastructure/repository"
	interfaces "interactions/internal/interfaces/grpc"
	"interactions/internal/usecase"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"github.com/kelseyhightower/envconfig"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/reflection"
)

func main() {

	// err := godotenv.Load("C:/Users/AlexK192/Desktop/interactions/.env")
	// if err != nil {
	// 	log.Printf("Ошибка при загрузке .env файла: %v", err)
	// }

	var cfg config.Config
	err := envconfig.Process("", &cfg)
	if err != nil {
		log.Fatal().Err(err).Msg("Ошибка загрузки конфигурации")
	}

	// Инициализация логера
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	logger.InitLogger(cfg.Logger)

	// Подключение к базе данных
	dbConn, err := db.NewPostgresDB(cfg.Database)
	if err != nil {
		log.Fatal().Err(err).Msg("Ошибка подключения к базе данных")
	}
	defer func() {
		if err := dbConn.Close(); err != nil {
			log.Error().Err(err).Msg("Ошибка при закрытии соединения с базой данных")
		}
	}()
	//Выполнение миграций
	// **Получение *sql.DB из *bun.DB**
	sqlDB := dbConn.DB

	// **Указание пути к миграциям**
	migrationsPath := "./migrations"

	// **Запуск миграций**
	err = db.RunMigrations(sqlDB, migrationsPath)
	if err != nil {
		log.Fatal().Err(err).Msg("Ошибка применения миграций")
	}

	// Инициализация репозитория и usecase
	interactionRepo := repository.NewInteractionRepositoryImpl(dbConn)
	interactionUsecase := usecase.NewInteractionUsecase(interactionRepo)

	// Инициализация метрик
	metrics.InitMetrics(cfg.Metrics.PrometheusPort)

	// Создание контекста с поддержкой завершения
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	// Обработка сигналов завершения
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c
		cancel()
	}()

	// Запуск серверов
	err = runServers(ctx, &cfg, interactionUsecase)
	if err != nil {
		log.Fatal().Err(err).Msg("Ошибка запуска серверов")
	}
}

func runServers(ctx context.Context, cfg *config.Config, interactionUsecase usecase.InteractionUsecase) error {
	// Инициализация gRPC-сервера
	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			middleware.LoggingInterceptor,
			middleware.MetricsMiddleware(metrics.ServerMetrics),
		),
		grpc.ChainStreamInterceptor(
			middleware.LoggingInterceptorStream,
			middleware.MetricsMiddlewareStream(metrics.ServerMetrics),
		),
	)

	// Создание и регистрация сервиса InteractionService
	interactionServiceServer := interfaces.NewInteractionServiceServerImpl(interactionUsecase)
	proto.RegisterInteractionServiceServer(grpcServer, interactionServiceServer)

	// Включение рефлексии для gRPC
	reflection.Register(grpcServer)
	grpcAddress := fmt.Sprintf("%s:%d", cfg.Server.GRPCAddress, cfg.Server.GRPCPort)

	// Создание gRPC-листенера
	grpcListener, err := net.Listen("tcp", grpcAddress)
	if err != nil {
		return err
	}

	// Запуск gRPC-сервера в отдельной горутине
	go func() {
		log.Info().Msgf("gRPC сервер запущен на %s", cfg.Server.GRPCAddress)
		if err := grpcServer.Serve(grpcListener); err != nil {
			log.Error().Err(err).Msg("Ошибка работы gRPC сервера")
		}
	}()

	// Настройка Mux для grpc-gateway
	gwMux := runtime.NewServeMux()
	grpcGatewayAddress := fmt.Sprintf("%s:%d", cfg.Server.GRPCAddress, cfg.Server.GRPCPort)
	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	err = proto.RegisterInteractionServiceHandlerFromEndpoint(ctx, gwMux, grpcGatewayAddress, opts)
	if err != nil {
		return err
	}

	httpAddress := fmt.Sprintf("%s:%d", cfg.Server.HTTPPAddress, cfg.Server.HTTPPort)

	// Настройка HTTP-сервера
	httpServer := &http.Server{
		Addr:    httpAddress,
		Handler: gwMux,
	}

	// Запуск HTTP-сервера в отдельной горутине
	go func() {
		log.Info().Msgf("HTTP сервер запущен на %s", cfg.Server.HTTPPAddress)
		if err := httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Error().Err(err).Msg("Ошибка работы HTTP сервера")
		}
	}()

	<-ctx.Done()
	log.Info().Msg("Завершение работы серверов...")

	// Graceful shutdown для gRPC и HTTP серверов
	grpcServer.GracefulStop()
	if err := httpServer.Shutdown(ctx); err != nil {
		log.Error().Err(err).Msg("Ошибка при завершении работы HTTP сервера")
	}

	return nil
}
