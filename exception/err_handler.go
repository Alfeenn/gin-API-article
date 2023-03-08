package exception

import (
	"net/http"

	"github.com/Alfeenn/article/model/web"
	"github.com/gin-gonic/gin"
)

func ErrHandler(c *gin.Context, err interface{}) {
	if ErrNotFound(c, err) {
		c.Next()
		return
	}
	if NoRoute(c, err) {
		c.Next()
		return
	}

	InternalServer(c, err)
	c.Next()
}

func ErrNotFound(c *gin.Context, err interface{}) bool {
	excption, ok := err.(NotFound)
	if ok {
		response := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: "NOT FOUND",
			Data:   excption.Err,
		}
		c.JSON(http.StatusNotFound, response)

		return true
	} else {
		return false
	}
}

func NoRoute(c *gin.Context, err interface{}) bool {
	excption, ok := err.(NotFound)
	if ok {
		response := web.WebResponse{
			Code:   http.StatusNotFound,
			Status: excption.Err,
		}

		c.JSON(http.StatusInternalServerError, response)
		return true
	} else {
		return false
	}

}

func InternalServer(c *gin.Context, err interface{}) {

	response := web.WebResponse{
		Code:   http.StatusInternalServerError,
		Status: "INTERNAL STATUS ERROR",
	}

	c.JSON(http.StatusInternalServerError, response)
}
