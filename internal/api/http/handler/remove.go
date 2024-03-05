package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) Remove(ctx *gin.Context) {
	h.s.Remove(ctx)
	ctx.JSON(http.StatusOK, "Remove")
}
