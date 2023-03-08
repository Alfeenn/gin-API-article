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
	engine := gin.Default()
	db := app.NewDB()
	router := middleware.NewMiddleware()
	repo := repository.NewRepository()
	service := service.NewService(repo, db)

	controller := controller.NewController(service)
	// engine.NoRoute(func(c *gin.Context) {
	// 	c.JSON(http.StatusNotFound, gin.H{"code": "404", "message": "Page not found"})
	// })
	engine.Use(router)
	engine.GET("/api/categories", controller.FindAll)
	engine.GET("/api/categories/:id", controller.Find)
	engine.PUT("/api/categories/:id", controller.Update)
	engine.POST("/api/categories", controller.Create)
	engine.POST("/api/categories/:id", controller.Delete)
	engine.Run("localhost:8000")
}
