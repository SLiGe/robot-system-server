package handler

import (
	"errors"
	"github.com/duke-git/lancet/v2/strutil"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"net/http"
	v1 "robot-system-server/api/v1"
	"robot-system-server/internal/model"
	"robot-system-server/internal/query"
	"robot-system-server/internal/service"
)

type FortuneHandler struct {
	*Handler
	fortuneService service.FortuneService
}

func NewFortuneHandler(handler *Handler, fortuneService service.FortuneService) *FortuneHandler {
	return &FortuneHandler{
		Handler:        handler,
		fortuneService: fortuneService,
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
	if req.QQ == nil {
		v1.HandleError(ctx, http.StatusBadRequest, nil, nil)
		return
	}
	todayFortune, err := h.fortuneService.GetFortuneOfToday(ctx, *req.QQ, req.IsOne, req.IsGroup, req.IsIntegral, *req.GroupNum, false)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	v1.HandleSuccess(ctx, todayFortune)
}

func (h *FortuneHandler) GetFortuneOfTodayOld(ctx *gin.Context) {
	qq := ctx.Param("qq")
	if strutil.IsBlank(qq) {
		v1.HandleError(ctx, http.StatusBadRequest, nil, nil)
		return
	}
	todayFortune, err := h.fortuneService.GetFortuneOfToday(ctx, qq, 0, 0, 0, "", true)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	v1.HandleSuccess(ctx, todayFortune)
}
