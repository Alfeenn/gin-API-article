package controller

import (
	"net/http"

	"github.com/Alfeenn/article/helper"
	"github.com/Alfeenn/article/model/web"
	"github.com/Alfeenn/article/service"
	"github.com/gin-gonic/gin"
)

type ControllerImpl struct {
	ServiceModel service.Service
}

func NewController(c service.Service) Controller {
	return &ControllerImpl{
		ServiceModel: c,
	}
}

func (c *ControllerImpl) Create(g *gin.Context) {
	req := web.CatRequest{}

	err := g.BindJSON(&req)
	helper.PanicIfErr(err)
	resp := c.ServiceModel.Create(g.Request.Context(), req)
	response := web.WebResponse{
		Code:   http.StatusCreated,
		Status: "CREATED",
		Data:   resp,
	}
	g.JSON(http.StatusOK, response)
}

func (c *ControllerImpl) Update(g *gin.Context) {
	req := web.UpdateRequest{}
	err := g.BindJSON(&req)
	helper.PanicIfErr(err)
	result := c.ServiceModel.Update(g.Request.Context(), req)
	g.JSON(http.StatusOK, result)

}

func (c *ControllerImpl) Delete(g *gin.Context) {

	result := c.ServiceModel.FindAll(g.Request.Context())
	g.JSON(http.StatusOK, result)
	panic("not implemented") // TODO: Implement
}

func (c *ControllerImpl) Find(g *gin.Context) {
	id := g.Params.ByName("id")
	result := c.ServiceModel.Find(g.Request.Context(), id)
	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}

	g.JSON(http.StatusOK, response)

}

func (c *ControllerImpl) FindAll(g *gin.Context) {

	result := c.ServiceModel.FindAll(g.Request.Context())
	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}
	g.JSON(http.StatusOK, response)

}

func (c *ControllerImpl) Ping(g *gin.Context) {
	g.JSON(200, gin.H{
		"message": "pong",
	})
}
