package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) GetList(ctx *gin.Context) {
	h.s.GetList(ctx)
	ctx.JSON(http.StatusOK, "list")
}
