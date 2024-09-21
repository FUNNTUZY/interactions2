package interfaces

import (
	"context"
	"interactions/api/proto"
	"interactions/internal/usecase"
)

type InteractionServiceServerImpl struct {
	proto.UnimplementedInteractionServiceServer
	Usecase usecase.InteractionUsecase
}

// NewInteractionServiceServerImpl создает новый сервер взаимодействий
func NewInteractionServiceServerImpl(usecase usecase.InteractionUsecase) proto.InteractionServiceServer {
	return &InteractionServiceServerImpl{Usecase: usecase}
}

func (s *InteractionServiceServerImpl) AddInteraction(ctx context.Context, req *proto.AddInteractionRequest) (*proto.InteractionResponse, error) {
	return s.Usecase.AddInteraction(ctx, req)
}

func (s *InteractionServiceServerImpl) GetInteraction(ctx context.Context, req *proto.GetInteractionRequest) (*proto.GetInteractionResponse, error) {
	return s.Usecase.GetInteraction(ctx, req)
}
