package metrics

import (
	"fmt"
	"net/http"

	promgrpc "github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/rs/zerolog/log"
)

var (
	// Общие метрики для gRPC запросов
	RequestCounter = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "grpc_requests_total",
			Help: "Total number of gRPC requests",
		},
		[]string{"method", "code"},
	)

	RequestDuration = prometheus.NewHistogramVec(
		prometheus.HistogramOpts{
			Name: "grpc_request_duration_seconds",
			Help: "Duration of gRPC requests",
		},
		[]string{"method"},
	)

	// Метрики gRPC-сервера
	ServerMetrics *promgrpc.ServerMetrics
)

// InitMetrics инициализирует метрики для gRPC и HTTP-сервера
func InitMetrics(prometheusPort int) {
	// Инициализация метрик для gRPC-сервера
	ServerMetrics = promgrpc.NewServerMetrics()

	// Регистрация пользовательских метрик
	prometheus.MustRegister(RequestCounter)
	prometheus.MustRegister(RequestDuration)

	// Запуск HTTP-сервера для Prometheus
	http.Handle("/metrics", promhttp.Handler())
	go func() {
		log.Info().Msgf("Prometheus metrics available at :%d/metrics", prometheusPort)
		err := http.ListenAndServe(fmt.Sprintf(":%d", prometheusPort), nil)
		if err != nil {
			log.Fatal().Err(err).Msg("Failed to start Prometheus metrics server")
		}
	}()
}
