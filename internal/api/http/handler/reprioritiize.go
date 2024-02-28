package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) Reprioritiize(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "reprioritiize")
}
