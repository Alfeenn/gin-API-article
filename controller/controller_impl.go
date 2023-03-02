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

func NewController(c service.Service) *ControllerImpl {
	return &ControllerImpl{
		ServiceModel: c,
	}
}

func (c *ControllerImpl) Create(g *gin.Context) {
	req := web.CatRequest{}

	err := g.BindJSON(&req)
	helper.PanicIfErr(err)
	resp := c.ServiceModel.Create(g.Request.Context(), req)

	g.JSON(http.StatusOK, resp)
}

func (c *ControllerImpl) Update(g *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (c *ControllerImpl) Delete(g *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (c *ControllerImpl) Find(g *gin.Context) {
	panic("not implemented") // TODO: Implement
}

func (c *ControllerImpl) FindAll(g *gin.Context) {
	limit := 3
	offset := 0
	result := c.ServiceModel.FindAll(g.Request.Context(), limit, offset)
	g.JSON(http.StatusOK, result)
	panic("not implemented") // TODO: Implement
}
