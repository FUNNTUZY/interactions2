package domain

import (
	"context"
	"interactions/internal/domain/entity"
)

// InteractionRepository - интерфейс для взаимодействий
type InteractionRepository interface {
	CreateInteraction(ctx context.Context, interaction entity.Interaction) error
	GetInteractions(ctx context.Context, userID string, adID string, sellerID string) ([]entity.Interaction, error)
}
