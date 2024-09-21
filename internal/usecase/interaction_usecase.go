package usecase

import (
	"context"
	"interactions/api/proto"
	"interactions/internal/domain"
	"interactions/internal/domain/entity"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

// InteractionUsecase - интерфейс, описывающий бизнес-логику работы с взаимодействиями
type InteractionUsecase interface {
	AddInteraction(ctx context.Context, req *proto.AddInteractionRequest) (*proto.InteractionResponse, error)
	GetInteraction(ctx context.Context, req *proto.GetInteractionRequest) (*proto.GetInteractionResponse, error)
}

// InteractionUsecaseImpl - структура, реализующая бизнес-логику взаимодействий
type InteractionUsecaseImpl struct {
	repo domain.InteractionRepository
}

// NewInteractionUsecase создает новый экземпляр InteractionUsecase
func NewInteractionUsecase(repo domain.InteractionRepository) InteractionUsecase {
	return &InteractionUsecaseImpl{repo: repo}
}

// AddInteraction создает новое взаимодействие
func (uc *InteractionUsecaseImpl) AddInteraction(ctx context.Context, req *proto.AddInteractionRequest) (*proto.InteractionResponse, error) {
	if req.UserId == "" || req.AdId == "" || req.SellerId == "" {
		return &proto.InteractionResponse{
			Success: false,
		}, nil
	}

	// Преобразуем enum в строку
	interactionType := req.Type.String()
	UUID := uuid.NewString()
	interaction := entity.Interaction{
		ID:              UUID, // Если ID генерируется базой данных, можно оставить пустым
		UserID:          req.UserId,
		AdID:            req.AdId,
		SellerID:        req.SellerId,
		InteractionType: interactionType,
		CreatedAt:       time.Now(),
	}

	err := uc.repo.CreateInteraction(ctx, interaction)
	if err != nil {
		log.Error().Err(err).Msg("Ошибка при добавлении взаимодействия")
		return &proto.InteractionResponse{
			Success: false,
		}, nil
	}

	return &proto.InteractionResponse{
		Success: true,
	}, nil
}

// GetInteraction возвращает взаимодействие по идентификаторам
func (uc *InteractionUsecaseImpl) GetInteraction(ctx context.Context, req *proto.GetInteractionRequest) (*proto.GetInteractionResponse, error) {
	interactions, err := uc.repo.GetInteractions(ctx, req.UserId, req.AdId, req.SellerId)
	if err != nil {
		log.Error().Err(err).Msg("Ошибка при получении взаимодействий")
		return &proto.GetInteractionResponse{}, nil
	}

	var types []proto.InteractionType

	for _, interaction := range interactions {
		switch interaction.InteractionType {
		case "message_sent":
			types = append(types, proto.InteractionType_message_sent)
		case "phone_revealed":
			types = append(types, proto.InteractionType_phone_revealed)
		}
	}

	return &proto.GetInteractionResponse{
		Type: types,
	}, nil
}
