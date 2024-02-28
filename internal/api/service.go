package api

import "github.com/gin-gonic/gin"

type Handler interface {
	Create(*gin.Context)
	GetList(*gin.Context)
	Remove(*gin.Context)
	Update(*gin.Context)
	Reprioritiize(*gin.Context)
	InitRoutes() *gin.Engine
}
