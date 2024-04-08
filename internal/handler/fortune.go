package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	v1 "robot-system-server/api/v1"
	"robot-system-server/internal/model"
	"robot-system-server/internal/query"
)

type FortuneHandler struct {
	*Handler
}

func NewFortuneHandler(handler *Handler) *FortuneHandler {
	return &FortuneHandler{
		Handler: handler,
	}
}

func (h *FortuneHandler) RandFortune(ctx *gin.Context) {
	var fortune model.QrFortune
	if err := query.QrFortune.UnderlyingDB().Order("rand()").First(&fortune).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			v1.HandleError(ctx, http.StatusNotFound, err, nil)
			return
		}
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
	}
	v1.HandleSuccess(ctx, &fortune)
}

func (h *FortuneHandler) GetFortuneOfToday(ctx *gin.Context) {
	req := new(v1.GetFortuneRequest)
	if err := ctx.ShouldBind(req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

}
