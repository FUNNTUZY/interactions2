package middleware

import (
	"context"
	"runtime/debug"

	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// PanicRecoveryInterceptor перехватывает панику и восстанавливает сервер
func PanicRecoveryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	defer func() {
		if r := recover(); r != nil {
			log.Error().Str("method", info.FullMethod).Msgf("Паника перехвачена: %v\n%s", r, debug.Stack())
		}
	}()

	h, err := handler(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "внутренняя ошибка сервера")
	}
	return h, nil
}
