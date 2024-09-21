package middleware

import (
	"context"

	"github.com/grpc-ecosystem/go-grpc-middleware/providers/prometheus" // Используем новый пакет
	"google.golang.org/grpc"
)

// MetricsMiddleware - middleware для метрик Prometheus (unary)
func MetricsMiddleware(metrics *prometheus.ServerMetrics) grpc.UnaryServerInterceptor {
	return func(
		ctx context.Context,
		req interface{},
		info *grpc.UnaryServerInfo,
		handler grpc.UnaryHandler,
	) (interface{}, error) {
		// Используем interceptor метрик
		return metrics.UnaryServerInterceptor()(ctx, req, info, handler)
	}
}

// MetricsMiddlewareStream - middleware для метрик Prometheus (stream)
func MetricsMiddlewareStream(metrics *prometheus.ServerMetrics) grpc.StreamServerInterceptor {
	return func(
		srv interface{},
		ss grpc.ServerStream,
		info *grpc.StreamServerInfo,
		handler grpc.StreamHandler,
	) error {
		// Используем interceptor метрик для потоков
		return metrics.StreamServerInterceptor()(srv, ss, info, handler)
	}
}
