package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"slovenia_petconnect/utils"
	"time"
)

func ZapLogger() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		c.Next()

		latency := time.Since(start)
		status := c.Writer.Status()

		if raw != "" {
			path = path + "?" + raw
		}

		fields := []zap.Field{
			zap.Int("status", status),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("client_ip", c.ClientIP()),
			zap.Duration("latency", latency),
		}

		switch {
		case status >= 500:
			utils.Logger.Error("SERVER ERROR", fields...)
		case status >= 400:
			utils.Logger.Warn("CLIENT ERROR", fields...)
		default:
			utils.Logger.Info("REQUEST", fields...)
		}
	}
}
