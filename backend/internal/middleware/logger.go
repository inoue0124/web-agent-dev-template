package middleware

import (
	"log/slog"
	"time"

	"github.com/gin-gonic/gin"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		method := c.Request.Method

		c.Next()

		slog.Info("request",
			"method", method,
			"path", path,
			"status", c.Writer.Status(),
			"duration", time.Since(start).String(),
		)
	}
}
