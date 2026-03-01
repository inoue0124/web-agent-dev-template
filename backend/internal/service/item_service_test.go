package service_test

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"

	"github.com/inoue0124/web-agent-dev-template/backend/internal/model"
	"github.com/inoue0124/web-agent-dev-template/backend/internal/service"
)

type mockItemRepo struct {
	items []model.Item
	err   error
}

func (m *mockItemRepo) List(_ context.Context) ([]model.Item, error) {
	if m.err != nil {
		return nil, m.err
	}
	return m.items, nil
}

func (m *mockItemRepo) GetByID(_ context.Context, id uuid.UUID) (*model.Item, error) {
	if m.err != nil {
		return nil, m.err
	}
	for i := range m.items {
		if m.items[i].ID == id {
			return &m.items[i], nil
		}
	}
	return nil, errors.New("not found")
}

func (m *mockItemRepo) Create(_ context.Context, item *model.Item) error {
	if m.err != nil {
		return m.err
	}
	item.ID = uuid.New()
	item.CreatedAt = time.Now()
	item.UpdatedAt = time.Now()
	m.items = append(m.items, *item)
	return nil
}

func (m *mockItemRepo) Update(_ context.Context, item *model.Item) error {
	if m.err != nil {
		return m.err
	}
	for i := range m.items {
		if m.items[i].ID == item.ID {
			m.items[i] = *item
			return nil
		}
	}
	return errors.New("not found")
}

func (m *mockItemRepo) Delete(_ context.Context, id uuid.UUID) error {
	if m.err != nil {
		return m.err
	}
	for i := range m.items {
		if m.items[i].ID == id {
			m.items = append(m.items[:i], m.items[i+1:]...)
			return nil
		}
	}
	return errors.New("not found")
}

func TestServiceList(t *testing.T) {
	repo := &mockItemRepo{
		items: []model.Item{
			{ID: uuid.New(), Name: "Item 1"},
			{ID: uuid.New(), Name: "Item 2"},
		},
	}
	svc := service.NewItemService(repo)

	items, err := svc.List(context.Background())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(items) != 2 {
		t.Errorf("expected 2 items, got %d", len(items))
	}
}

func TestServiceListError(t *testing.T) {
	repo := &mockItemRepo{err: errors.New("db error")}
	svc := service.NewItemService(repo)

	_, err := svc.List(context.Background())
	if err == nil {
		t.Error("expected error, got nil")
	}
}

func TestServiceGetByID(t *testing.T) {
	id := uuid.New()
	repo := &mockItemRepo{
		items: []model.Item{
			{ID: id, Name: "Test Item"},
		},
	}
	svc := service.NewItemService(repo)

	item, err := svc.GetByID(context.Background(), id.String())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if item.Name != "Test Item" {
		t.Errorf("expected name 'Test Item', got '%s'", item.Name)
	}
}

func TestServiceGetByIDInvalidUUID(t *testing.T) {
	repo := &mockItemRepo{}
	svc := service.NewItemService(repo)

	_, err := svc.GetByID(context.Background(), "not-a-uuid")
	if err == nil {
		t.Error("expected error for invalid UUID, got nil")
	}
}

func TestServiceCreate(t *testing.T) {
	repo := &mockItemRepo{}
	svc := service.NewItemService(repo)

	desc := "A description"
	req := model.CreateItemRequest{
		Name:        "New Item",
		Description: &desc,
	}

	item, err := svc.Create(context.Background(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if item.Name != "New Item" {
		t.Errorf("expected name 'New Item', got '%s'", item.Name)
	}
	if item.Description == nil || *item.Description != desc {
		t.Errorf("expected description '%s', got '%v'", desc, item.Description)
	}
}

func TestServiceUpdate(t *testing.T) {
	id := uuid.New()
	repo := &mockItemRepo{
		items: []model.Item{
			{ID: id, Name: "Old Name"},
		},
	}
	svc := service.NewItemService(repo)

	newName := "New Name"
	req := model.UpdateItemRequest{Name: &newName}

	item, err := svc.Update(context.Background(), id.String(), req)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if item.Name != "New Name" {
		t.Errorf("expected name 'New Name', got '%s'", item.Name)
	}
}

func TestServiceDelete(t *testing.T) {
	id := uuid.New()
	repo := &mockItemRepo{
		items: []model.Item{
			{ID: id, Name: "To Delete"},
		},
	}
	svc := service.NewItemService(repo)

	err := svc.Delete(context.Background(), id.String())
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(repo.items) != 0 {
		t.Errorf("expected 0 items after delete, got %d", len(repo.items))
	}
}

func TestServiceDeleteInvalidUUID(t *testing.T) {
	repo := &mockItemRepo{}
	svc := service.NewItemService(repo)

	err := svc.Delete(context.Background(), "not-a-uuid")
	if err == nil {
		t.Error("expected error for invalid UUID, got nil")
	}
}
