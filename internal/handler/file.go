package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
	v1 "robot-system-server/api/v1"
	"robot-system-server/internal/service"
)

type FileHandler struct {
	*Handler
	fileService service.FileService
}

func NewFileHandler(handler *Handler, fileService service.FileService) *FileHandler {
	return &FileHandler{
		Handler:     handler,
		fileService: fileService,
	}
}

func (h *FileHandler) UploadFile(ctx *gin.Context) {
	uploadInfo, err := h.fileService.UploadFile(ctx)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
	v1.HandleSuccess(ctx, uploadInfo)
}

func (h *FileHandler) Img(ctx *gin.Context) {
	err := h.fileService.Img(ctx)
	if err != nil {
		v1.HandleError(ctx, http.StatusInternalServerError, err, nil)
		return
	}
}
