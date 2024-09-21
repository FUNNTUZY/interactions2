package middleware

import (
	"context"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/status"
)

// ErrorInterceptor перехватывает ошибки и логирует их
func ErrorInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	h, err := handler(ctx, req)

	// Если произошла ошибка, логируем её и возвращаем статус
	if err != nil {
		st, _ := status.FromError(err)
		log.Error().Err(err).Str("method", info.FullMethod).Msgf("gRPC ошибка: %s", st.Message())
		return nil, err
	}

	return h, nil
}
