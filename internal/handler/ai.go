package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	v1 "robot-system-server/api/v1"
	"robot-system-server/internal/service"
	"unicode/utf8"
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

func (h *AiHandler) Poem(ctx *gin.Context) {
	var req v1.PoemRequest
	if err := ctx.ShouldBindQuery(&req); err != nil {
		v1.HandleError(ctx, http.StatusBadRequest, err, nil)
		return
	}
	keywordLen := utf8.RuneCountInString(req.Keyword)
	if keywordLen < 2 || keywordLen > 8 {
		v1.HandleError(ctx, http.StatusBadRequest, v1.NewParamError("藏字内容，2-8个字"), nil)
		return
	}
	if req.Num != 5 && req.Num != 7 {
		req.Num = 5
	}
	poemType := req.Type
	if poemType != 1 && poemType != 2 && poemType != 3 && poemType != 4 && poemType != 5 {
		poemType = 1 // 默认为1
	}
	rhyme := req.Rhyme
	if rhyme != 1 && rhyme != 2 && rhyme != 3 {
		rhyme = 1 // 默认为1
	}
	poem, err := h.aiService.Poem(ctx, &req)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, poem)
}
