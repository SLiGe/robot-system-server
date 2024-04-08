package handler

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	v1 "robot-system-server/api/v1"
	"robot-system-server/internal/model"
	"robot-system-server/internal/query"
	"time"
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
	var fortune *model.QrFortuneDatum
	err := query.QrFortuneDatum.UnderlyingDB().
		Select("id, qq, json_data, group_num, fortune_date, update_date, create_date").
		Where("DATE(FORTUNE_DATE) = ? and QQ = ?", time.Now().Format("2006-01-02"), req.QQ).First(&fortune).Error
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		v1.HandleSuccess(ctx, &fortune)
		return
	}
	h.RandFortune(ctx)

}
