package handler

import "github.com/gin-gonic/gin"

type handler struct {
}

func NewHandler() *handler {
	return &handler{}
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
