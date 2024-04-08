package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"net/http"
	v1 "robot-system-server/api/v1"
	"robot-system-server/internal/service"
)

type BeasenHandler struct {
	*Handler
	beasenService service.BeasenService
}

func NewBeasenHandler(handler *Handler, beasenService service.BeasenService) *BeasenHandler {
	return &BeasenHandler{
		Handler:       handler,
		beasenService: beasenService,
	}
}

func (h *BeasenHandler) GetBeasen(ctx *gin.Context) {
	id := ctx.Query("id")
	toInt64 := cast.ToInt64(id)
	beasen, err := h.beasenService.GetBeasen(ctx, toInt64)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	v1.HandleSuccess(ctx, beasen)
}

func (h *BeasenHandler) RandResult(ctx *gin.Context) {
	result, err := h.beasenService.RandResult(ctx)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	err = ctx.ShouldBindJSON(result)
	if err != nil {
		v1.HandleSuccess(ctx, result)
	}

}
