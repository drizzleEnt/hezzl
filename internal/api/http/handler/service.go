package handler

import (
	"github.com/drizzleent/hezzl/internal/service"
	"github.com/gin-gonic/gin"
)

type handler struct {
	s service.Service
}

func NewHandler(srv service.Service) *handler {
	return &handler{
		s: srv,
	}
}

func (h *handler) InitRoutes() *gin.Engine {
	router := gin.New()

	goods := router.Group("/goods")
	{
		goods.GET("/list/:limit/:offset", h.GetList)
	}

	good := router.Group("/good")
	{
		good.POST("/create/:projectId", h.Create)
		good.PATCH("/update/:id/:projectId", h.Update)
		good.PATCH("/reprioritiize/:id/:projectId", h.GetList)
		good.DELETE("/remove/:id/:projectId", h.GetList)
	}

	return router
}
