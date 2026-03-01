package main

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"

	"github.com/inoue0124/web-agent-dev-template/backend/internal/config"
	"github.com/inoue0124/web-agent-dev-template/backend/internal/database"
	"github.com/inoue0124/web-agent-dev-template/backend/internal/handler"
	"github.com/inoue0124/web-agent-dev-template/backend/internal/middleware"
	"github.com/inoue0124/web-agent-dev-template/backend/internal/repository"
	"github.com/inoue0124/web-agent-dev-template/backend/internal/service"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		slog.Error("failed to load config", "error", err)
		os.Exit(1)
	}

	if cfg.GinMode == "release" {
		gin.SetMode(gin.ReleaseMode)
	}

	db, err := database.NewDB(cfg.DatabaseURL)
	if err != nil {
		slog.Error("failed to connect to database", "error", err)
		os.Exit(1)
	}

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())
	r.Use(middleware.CORS())

	healthHandler := handler.NewHealthHandler()
	r.GET("/health", healthHandler.Health)

	// Item CRUD
	itemRepo := repository.NewItemRepository(db)
	itemService := service.NewItemService(itemRepo)
	itemHandler := handler.NewItemHandler(itemService)

	v1 := r.Group("/api/v1")
	{
		items := v1.Group("/items")
		{
			items.GET("", itemHandler.List)
			items.GET("/:id", itemHandler.GetByID)
			items.POST("", itemHandler.Create)
			items.PUT("/:id", itemHandler.Update)
			items.DELETE("/:id", itemHandler.Delete)
		}
	}

	slog.Info("server starting", "port", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
}
