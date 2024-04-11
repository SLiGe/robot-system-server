package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	v1 "robot-system-server/api/v1"
	"robot-system-server/internal/service"
)

type SignInHandler struct {
	*Handler
	signInService service.SignInService
}

func NewSignInHandler(handler *Handler, signInService service.SignInService) *SignInHandler {
	return &SignInHandler{
		Handler:       handler,
		signInService: signInService,
	}
}

func (h *SignInHandler) SignIn(ctx *gin.Context) {
	var req v1.SignInForQqRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	signInData, err := h.signInService.DoSignIn(req)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}

	v1.HandleSuccess(ctx, signInData)
}
func (h *SignInHandler) QuerySignInData(ctx *gin.Context) {
	var req v1.QuerySignInDataRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	signInData, err := h.signInService.QuerySignInData(req)
	if err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	v1.HandleSuccess(ctx, signInData)
}
