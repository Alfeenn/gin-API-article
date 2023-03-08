package controller

import (
	"fmt"
	"net/http"

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
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "BAD REQUEST",
			"error":   "VALIDATEERR-1",
			"message": "Invalid inputs. Please check your inputs"})
		return
	}

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
	req.Id = g.Params.ByName("id")
	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":    http.StatusBadRequest,
			"status":  "BAD REQUEST",
			"error":   "VALIDATEERR-1",
			"message": "Invalid inputs. Please check your inputs"})
		return
	}
	result := c.ServiceModel.Update(g.Request.Context(), req)
	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
		Data:   result,
	}
	g.JSON(http.StatusOK, response)

}

func (c *ControllerImpl) Delete(g *gin.Context) {

	id := g.Params.ByName("id")

	c.ServiceModel.Delete(g.Request.Context(), id)
	response := web.WebResponse{
		Code:   http.StatusOK,
		Status: "OK",
	}
	fmt.Println(response)
	g.JSON(http.StatusOK, response)
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
