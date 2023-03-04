package app

import (
	"github.com/Alfeenn/article/controller"
	"github.com/gin-gonic/gin"
)

func NewRouter(controller controller.Controller) *gin.Engine {
	engine := gin.New()

	engine.GET("/api/categories", controller.FindAll)
	engine.GET("/api/categories/:id", controller.Find)
	engine.PUT("/api/categories/:id", controller.Update)
	engine.POST("/api/categories", controller.Create)
	engine.POST("/api/categories/:id", controller.Delete)
	engine.Run(":8000")
	return engine
}
