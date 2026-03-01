package handler

import (
	"context"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/inoue0124/web-agent-dev-template/backend/internal/model"
)

type ItemService interface {
	List(ctx context.Context) ([]model.Item, error)
	GetByID(ctx context.Context, id string) (*model.Item, error)
	Create(ctx context.Context, req model.CreateItemRequest) (*model.Item, error)
	Update(ctx context.Context, id string, req model.UpdateItemRequest) (*model.Item, error)
	Delete(ctx context.Context, id string) error
}

type ItemHandler struct {
	service ItemService
}

func NewItemHandler(service ItemService) *ItemHandler {
	return &ItemHandler{service: service}
}

func (h *ItemHandler) List(c *gin.Context) {
	items, err := h.service.List(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch items"})
		return
	}
	c.JSON(http.StatusOK, items)
}

func (h *ItemHandler) GetByID(c *gin.Context) {
	item, err := h.service.GetByID(c.Request.Context(), c.Param("id"))
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to fetch item"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *ItemHandler) Create(c *gin.Context) {
	var req model.CreateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	item, err := h.service.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to create item"})
		return
	}
	c.JSON(http.StatusCreated, item)
}

func (h *ItemHandler) Update(c *gin.Context) {
	var req model.UpdateItemRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request body"})
		return
	}
	item, err := h.service.Update(c.Request.Context(), c.Param("id"), req)
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to update item"})
		return
	}
	c.JSON(http.StatusOK, item)
}

func (h *ItemHandler) Delete(c *gin.Context) {
	if err := h.service.Delete(c.Request.Context(), c.Param("id")); err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "item not found"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": "failed to delete item"})
		return
	}
	c.JSON(http.StatusNoContent, nil)
}
