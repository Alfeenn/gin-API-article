package main

import (
	"github.com/Alfeenn/article/app"
	"github.com/Alfeenn/article/controller"
	"github.com/Alfeenn/article/middleware"
	"github.com/Alfeenn/article/repository"
	"github.com/Alfeenn/article/service"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	db := app.NewDB()
	repo := repository.NewRepository()
	service := service.NewService(repo, db)
	controller := controller.NewController(service)
	middleware := middleware.NewMiddleware()

	engine.GET("/api/categories", controller.FindAll)
	engine.GET("/api/categories/:id", controller.Find)
	engine.PUT("/api/categories/:id", controller.Update)
	engine.POST("/api/categories", controller.Create)
	engine.POST("/api/categories/:id", controller.Delete)
	engine.Use(middleware)
	engine.Run("localhost:8000")
}
