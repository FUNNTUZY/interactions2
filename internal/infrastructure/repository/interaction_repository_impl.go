package repository

import (
	"context"
	"interactions/internal/domain"
	"interactions/internal/domain/entity"

	_ "github.com/lib/pq"
	"github.com/uptrace/bun"
)

// InteractionRepositoryImpl - реализация интерфейса InteractionRepository
type InteractionRepositoryImpl struct {
	db *bun.DB
}

// NewInteractionRepositoryImpl создает новый экземпляр InteractionRepository
func NewInteractionRepositoryImpl(db *bun.DB) domain.InteractionRepository {
	return &InteractionRepositoryImpl{db: db}
}

// CreateInteraction создает новое взаимодействие в базе данных
func (repo *InteractionRepositoryImpl) CreateInteraction(ctx context.Context, interaction entity.Interaction) error {
	_, err := repo.db.NewInsert().Model(&interaction).Exec(ctx)
	return err
}

func (repo *InteractionRepositoryImpl) GetInteractions(ctx context.Context, userID string, adID string, sellerID string) ([]entity.Interaction, error) {
	var interactions []entity.Interaction
	err := repo.db.NewSelect().Model(&interactions).
		Where("user_id = ?", userID).
		Where("ad_id = ?", adID).
		Where("seller_id = ?", sellerID).
		Scan(ctx)
	if err != nil {
		return nil, err
	}
	return interactions, nil
}
