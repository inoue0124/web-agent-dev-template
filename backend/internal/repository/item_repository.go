package repository

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/inoue0124/web-agent-dev-template/backend/internal/model"
)

type ItemRepository struct {
	db *gorm.DB
}

func NewItemRepository(db *gorm.DB) *ItemRepository {
	return &ItemRepository{db: db}
}

func (r *ItemRepository) List(ctx context.Context) ([]model.Item, error) {
	var items []model.Item
	if err := r.db.WithContext(ctx).Order("created_at DESC").Find(&items).Error; err != nil {
		return nil, fmt.Errorf("failed to list items: %w", err)
	}
	return items, nil
}

func (r *ItemRepository) GetByID(ctx context.Context, id uuid.UUID) (*model.Item, error) {
	var item model.Item
	if err := r.db.WithContext(ctx).First(&item, "id = ?", id).Error; err != nil {
		return nil, fmt.Errorf("failed to get item: %w", err)
	}
	return &item, nil
}

func (r *ItemRepository) Create(ctx context.Context, item *model.Item) error {
	if err := r.db.WithContext(ctx).Create(item).Error; err != nil {
		return fmt.Errorf("failed to create item: %w", err)
	}
	return nil
}

func (r *ItemRepository) Update(ctx context.Context, item *model.Item) error {
	if err := r.db.WithContext(ctx).Save(item).Error; err != nil {
		return fmt.Errorf("failed to update item: %w", err)
	}
	return nil
}

func (r *ItemRepository) Delete(ctx context.Context, id uuid.UUID) error {
	if err := r.db.WithContext(ctx).Delete(&model.Item{}, "id = ?", id).Error; err != nil {
		return fmt.Errorf("failed to delete item: %w", err)
	}
	return nil
}
