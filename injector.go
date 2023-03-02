package main

import (
	"net/http"

	"github.com/Alfeenn/article/app"
	"github.com/Alfeenn/article/controller"
	"github.com/Alfeenn/article/middleware"
	"github.com/Alfeenn/article/repository"
	"github.com/Alfeenn/article/service"
	"github.com/gin-gonic/gin"
	"github.com/google/wire"
)

var CategorySet = wire.NewSet(
	repository.NewRepository,
	wire.Bind(new(repository.Repository), new(*repository.RepoImpl)),
	service.NewService,
	wire.Bind(new(service.Service), new(*service.ServiceImpl)),
	controller.NewController,
	wire.Bind(new(controller.Controller), new(*controller.ControllerImpl)),
	app.NewRouter,
	wire.Bind(new(http.Handler), new(*gin.Engine)),
)

func InitializeServer() *gin.Engine {
	wire.Build(
		app.NewDB,
		CategorySet,
		app.NewRouter,
		middleware.NewMiddleware,
		NewServer,
	)
	return nil
}
