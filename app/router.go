package app

import (
	"github.com/Alfeenn/article/controller"
	"github.com/Alfeenn/article/middleware"
	"github.com/gin-gonic/gin"
)

func NewRouter(controller controller.Controller) *gin.Engine {
	engine := gin.Default()
	engine.Use(middleware.NewMiddleware())
	engine.GET("/api/categories", controller.FindAll)
	engine.GET("/api/categories/:id", controller.Find)
	engine.PUT("/api/categories/:id", controller.Update)
	engine.POST("/api/categories", controller.Create)
	engine.POST("/api/categories/:id", controller.Delete)
	return engine
}
