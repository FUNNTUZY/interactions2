package middleware

import (
	"context"
	"time"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// LoggingInterceptor - перехватчик для логгирования запросов gRPC
func LoggingInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (resp interface{}, err error) {
	start := time.Now()
	log.Info().Msgf("gRPC call: %s", info.FullMethod)

	// Вызов реального обработчика
	resp, err = handler(ctx, req)

	// Логгирование результата
	log.Info().
		Str("method", info.FullMethod).
		Dur("duration", time.Since(start)).
		Err(err).
		Msgf("gRPC response status: %s", status.Code(err))

	return resp, err
}

func LoggingInterceptorStream(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	start := time.Now()
	log.Info().Msgf("gRPC stream call: %s", info.FullMethod)

	// Вызов реального обработчика
	err := handler(srv, ss)

	// Логгирование результата
	log.Info().
		Str("method", info.FullMethod).
		Dur("duration", time.Since(start)).
		Err(err).
		Msgf("gRPC stream response status: %s", status.Code(err))

	return err
}
