package handler_test

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"gorm.io/gorm"

	"github.com/inoue0124/web-agent-dev-template/backend/internal/handler"
	"github.com/inoue0124/web-agent-dev-template/backend/internal/model"
)

type mockItemService struct {
	items []model.Item
	err   error
}

func (m *mockItemService) List(_ context.Context) ([]model.Item, error) {
	return m.items, m.err
}

func (m *mockItemService) GetByID(_ context.Context, id string) (*model.Item, error) {
	if m.err != nil {
		return nil, m.err
	}
	for i := range m.items {
		if m.items[i].ID.String() == id {
			return &m.items[i], nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}

func (m *mockItemService) Create(_ context.Context, req model.CreateItemRequest) (*model.Item, error) {
	if m.err != nil {
		return nil, m.err
	}
	item := &model.Item{
		ID:          uuid.New(),
		Name:        req.Name,
		Description: req.Description,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	return item, nil
}

func (m *mockItemService) Update(_ context.Context, id string, req model.UpdateItemRequest) (*model.Item, error) {
	if m.err != nil {
		return nil, m.err
	}
	for i := range m.items {
		if m.items[i].ID.String() == id {
			if req.Name != nil {
				m.items[i].Name = *req.Name
			}
			if req.Description != nil {
				m.items[i].Description = req.Description
			}
			return &m.items[i], nil
		}
	}
	return nil, gorm.ErrRecordNotFound
}

func (m *mockItemService) Delete(_ context.Context, id string) error {
	if m.err != nil {
		return m.err
	}
	for _, item := range m.items {
		if item.ID.String() == id {
			return nil
		}
	}
	return gorm.ErrRecordNotFound
}

func setupRouter(svc handler.ItemService) *gin.Engine {
	gin.SetMode(gin.TestMode)
	r := gin.New()
	h := handler.NewItemHandler(svc)
	g := r.Group("/api/v1/items")
	{
		g.GET("", h.List)
		g.GET("/:id", h.GetByID)
		g.POST("", h.Create)
		g.PUT("/:id", h.Update)
		g.DELETE("/:id", h.Delete)
	}
	return r
}

func TestListItems(t *testing.T) {
	id := uuid.New()
	svc := &mockItemService{
		items: []model.Item{
			{ID: id, Name: "Test Item", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		},
	}
	r := setupRouter(svc)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/items", nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}

	var items []model.Item
	if err := json.Unmarshal(w.Body.Bytes(), &items); err != nil {
		t.Fatalf("failed to decode response: %v", err)
	}
	if len(items) != 1 {
		t.Errorf("expected 1 item, got %d", len(items))
	}
}

func TestGetByID(t *testing.T) {
	id := uuid.New()
	svc := &mockItemService{
		items: []model.Item{
			{ID: id, Name: "Test Item", CreatedAt: time.Now(), UpdatedAt: time.Now()},
		},
	}
	r := setupRouter(svc)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/items/"+id.String(), nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusOK {
		t.Errorf("expected status %d, got %d", http.StatusOK, w.Code)
	}
}

func TestGetByIDNotFound(t *testing.T) {
	svc := &mockItemService{items: []model.Item{}}
	r := setupRouter(svc)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodGet, "/api/v1/items/"+uuid.New().String(), nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNotFound {
		t.Errorf("expected status %d, got %d", http.StatusNotFound, w.Code)
	}
}

func TestCreateItem(t *testing.T) {
	svc := &mockItemService{}
	r := setupRouter(svc)

	body, _ := json.Marshal(model.CreateItemRequest{Name: "New Item"})
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/items", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusCreated {
		t.Errorf("expected status %d, got %d", http.StatusCreated, w.Code)
	}
}

func TestCreateItemBadRequest(t *testing.T) {
	svc := &mockItemService{}
	r := setupRouter(svc)

	body := []byte(`{"invalid": true}`)
	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodPost, "/api/v1/items", bytes.NewReader(body))
	req.Header.Set("Content-Type", "application/json")
	r.ServeHTTP(w, req)

	if w.Code != http.StatusBadRequest {
		t.Errorf("expected status %d, got %d", http.StatusBadRequest, w.Code)
	}
}

func TestDeleteItem(t *testing.T) {
	id := uuid.New()
	svc := &mockItemService{
		items: []model.Item{
			{ID: id, Name: "Test Item"},
		},
	}
	r := setupRouter(svc)

	w := httptest.NewRecorder()
	req, _ := http.NewRequest(http.MethodDelete, "/api/v1/items/"+id.String(), nil)
	r.ServeHTTP(w, req)

	if w.Code != http.StatusNoContent {
		t.Errorf("expected status %d, got %d", http.StatusNoContent, w.Code)
	}
}
