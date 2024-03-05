package handler

import (
	"fmt"
	"net/http"

	"github.com/drizzleent/hezzl/internal/model"
	"github.com/gin-gonic/gin"
)

func (h *handler) Create(ctx *gin.Context) {
	var good model.Good
	ctx.Bind(&good)
	fmt.Printf("%v\n", good)
	res, err := h.s.Create(ctx)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err.Error())
	}
	ctx.JSON(http.StatusOK, res)
}
