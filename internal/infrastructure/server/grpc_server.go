package server

import (
	"context"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"

	"interactions/api/proto"
	"interactions/internal/infrastructure/middleware"

	grpc_prometheus "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type GrpcServer struct {
	server *grpc.Server
}

// NewGrpcServer создает и настраивает gRPC сервер с мидлвейрами
func NewGrpcServer(uc proto.InteractionServiceServer) *GrpcServer {
	// Инициализация метрик
	grpcMetrics := grpc_prometheus.NewServerMetrics()

	// Настройка мидлвейров
	opts := []grpc.ServerOption{
		grpc.ChainUnaryInterceptor(
			middleware.LoggingInterceptor,        // Ваш логгинг middleware
			grpcMetrics.UnaryServerInterceptor(), // Метрики Prometheus
		),
		grpc.ChainStreamInterceptor(
			middleware.LoggingInterceptorStream,   // Логгирование для stream
			grpcMetrics.StreamServerInterceptor(), // Метрики для stream
		),
	}

	// Создание gRPC сервера
	server := grpc.NewServer(opts...)

	// Регистрация сервисов
	proto.RegisterInteractionServiceServer(server, uc)
	reflection.Register(server)
	// Регистрация метрик Prometheus
	grpcMetrics.InitializeMetrics(server)

	// Регистрация рефлексии для gRPC (для gRPC-инструментов, например `grpcurl`)
	reflection.Register(server)

	return &GrpcServer{server: server}
}

// Run запускает gRPC сервер и обрабатывает завершение работы
func (s *GrpcServer) Run(ctx context.Context, addr string) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("Ошибка при запуске сервера: %w", err)
	}

	// Graceful shutdown
	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, os.Interrupt, syscall.SIGTERM)
		<-c

		fmt.Println("Выключение сервера...")
		s.server.GracefulStop()
	}()

	fmt.Printf("gRPC сервер запущен на %s\n", addr)
	return s.server.Serve(lis)
}
