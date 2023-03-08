package helper

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func PanicIfErr(err error) {
	if err != nil {
		panic(err)
	}
}

func ErrValidation(err error, g *gin.Context) {

	if err != nil {
		g.AbortWithStatusJSON(http.StatusBadRequest, gin.H{
			"code":   http.StatusBadRequest,
			"status": "BAD REQUEST",
		})
		return
	}

}
