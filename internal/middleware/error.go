package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
	v1 "robot-system-server/api/v1"
	"robot-system-server/pkg/log"
)

func ErrorHandler(logger *log.Logger) gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				logger.WithContext(c).Error("Server Error: ", zap.Any("err", err))
				//c.AbortWithStatus(http.StatusInternalServerError)
				v1.HandleError(c, http.StatusInternalServerError, v1.ErrInternalServerError, nil)
			}
		}()

		c.Next()
	}
}
