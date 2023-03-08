package test

import (
	"database/sql"

	"github.com/Alfeenn/article/app"
	"github.com/Alfeenn/article/controller"
	"github.com/Alfeenn/article/repository"
	"github.com/Alfeenn/article/service"
	"github.com/gin-gonic/gin"
)

func SetupRouter(db *sql.DB) *gin.Engine {
	repository := repository.NewRepository()
	service := service.NewService(repository, db)
	controller := controller.NewController(service)
	router := app.NewRouter(controller)
	return router
}
