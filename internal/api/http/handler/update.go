package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (h *handler) Update(ctx *gin.Context) {
	h.s.Update(ctx)
	ctx.JSON(http.StatusOK, "Update")
}
