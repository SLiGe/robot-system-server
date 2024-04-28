package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	"robot-system-server/pkg/log"
)

func ErrorHandler(logger *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.WithContext(c).Error("Server Error: ", zap.Any("err", err))
				c.AbortWithStatus(http.StatusInternalServerError)
			}
		}()

		c.Next()
	}
}
