package http

import (
	"time"

	global "github.com/Rawipass/chat-service/global_variable"
	"github.com/Rawipass/chat-service/logger"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"go.uber.org/zap"
)

func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		path := c.Request.URL.Path
		query := c.Request.URL.RawQuery

		c.Next()

		// Log Request
		latency := time.Since(t)
		requestId := c.GetString(global.KEY_REQUEST_ID)
		logger.Logger.Infow(path,
			zap.String("request-id", requestId),
			zap.Int("status", c.Writer.Status()),
			zap.String("method", c.Request.Method),
			zap.String("path", path),
			zap.String("query", query),
			zap.String("ip", c.ClientIP()),
			zap.String("user-agent", c.Request.UserAgent()),
			zap.String("errors", c.Errors.ByType(gin.ErrorTypePrivate).String()),
			zap.Duration("latency", latency),
		)
	}
}

func SetRequestId() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestId := c.GetHeader(global.HTTP_HEADER_REQUEST_ID)
		if requestId == "" {
			requestId = uuid.New().String()
		}

		c.Set(global.KEY_LOGGER, logger.Logger.With(
			global.KEY_REQUEST_ID, requestId,
			global.KEY_PART, "interface",
		))

		c.Set(global.KEY_REQUEST_ID, requestId)
		c.Next()
	}
}
