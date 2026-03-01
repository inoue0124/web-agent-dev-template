package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"

	"github.com/inoue0124/web-agent-dev-template/backend/internal/model"
)

type ItemRepository interface {
	List(ctx context.Context) ([]model.Item, error)
	GetByID(ctx context.Context, id uuid.UUID) (*model.Item, error)
	Create(ctx context.Context, item *model.Item) error
	Update(ctx context.Context, item *model.Item) error
	Delete(ctx context.Context, id uuid.UUID) error
}

type ItemService struct {
	repo ItemRepository
}

func NewItemService(repo ItemRepository) *ItemService {
	return &ItemService{repo: repo}
}

func (s *ItemService) List(ctx context.Context) ([]model.Item, error) {
	return s.repo.List(ctx)
}

func (s *ItemService) GetByID(ctx context.Context, id string) (*model.Item, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid item ID: %w", err)
	}
	return s.repo.GetByID(ctx, uid)
}

func (s *ItemService) Create(ctx context.Context, req model.CreateItemRequest) (*model.Item, error) {
	item := &model.Item{
		Name:        req.Name,
		Description: req.Description,
	}
	if err := s.repo.Create(ctx, item); err != nil {
		return nil, err
	}
	return item, nil
}

func (s *ItemService) Update(ctx context.Context, id string, req model.UpdateItemRequest) (*model.Item, error) {
	uid, err := uuid.Parse(id)
	if err != nil {
		return nil, fmt.Errorf("invalid item ID: %w", err)
	}

	item, err := s.repo.GetByID(ctx, uid)
	if err != nil {
		return nil, err
	}

	if req.Name != nil {
		item.Name = *req.Name
	}
	if req.Description != nil {
		item.Description = req.Description
	}

	if err := s.repo.Update(ctx, item); err != nil {
		return nil, err
	}
	return item, nil
}

func (s *ItemService) Delete(ctx context.Context, id string) error {
	uid, err := uuid.Parse(id)
	if err != nil {
		return fmt.Errorf("invalid item ID: %w", err)
	}
	return s.repo.Delete(ctx, uid)
}
