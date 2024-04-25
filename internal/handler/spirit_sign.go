package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	v1 "robot-system-server/api/v1"
	"robot-system-server/internal/service"
)

type SpiritSignHandler struct {
	*Handler
	spiritSignService service.SpiritSignService
}

func NewSpiritSignHandler(handler *Handler, spiritSignService service.SpiritSignService) *SpiritSignHandler {
	return &SpiritSignHandler{
		Handler:           handler,
		spiritSignService: spiritSignService,
	}
}

func (h *SpiritSignHandler) OneSignPerDay(ctx *gin.Context) {
	var req v1.SignPerDayReq
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	res, err := h.spiritSignService.OneSignPerDay(req)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, res)
		return
	}
	v1.HandleSuccess(ctx, res)

}

func (h *SpiritSignHandler) ViewCs(ctx *gin.Context) {
	h.spiritSignService.ViewCsLq(ctx)
}

func (h *SpiritSignHandler) ViewYl(ctx *gin.Context) {
	h.spiritSignService.ViewYlLq(ctx)
}

func (h *SpiritSignHandler) ViewGy(ctx *gin.Context) {
	h.spiritSignService.ViewGyLq(ctx)
}
