package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) Create(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, "create")
}
