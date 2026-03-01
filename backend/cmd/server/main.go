package main

import (
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/inoue0124/web-agent-dev-template/backend/internal/config"
	"github.com/inoue0124/web-agent-dev-template/backend/internal/handler"
	"github.com/inoue0124/web-agent-dev-template/backend/internal/middleware"
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

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(middleware.Logger())
	r.Use(middleware.CORS())

	healthHandler := handler.NewHealthHandler()
	r.GET("/health", healthHandler.Health)

	slog.Info("server starting", "port", cfg.Port)
	if err := r.Run(":" + cfg.Port); err != nil {
		slog.Error("server failed", "error", err)
		os.Exit(1)
	}
}
