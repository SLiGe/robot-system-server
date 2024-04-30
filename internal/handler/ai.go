package handler

import (
	"github.com/gin-gonic/gin"
	"robot-system-server/internal/service"
)

type AiHandler struct {
	*Handler
	aiService service.AiService
}

func NewAiHandler(handler *Handler, aiService service.AiService) *AiHandler {
	return &AiHandler{
		Handler:   handler,
		aiService: aiService,
	}
}

func (h *AiHandler) GetAi(ctx *gin.Context) {

}
