package middleware

import (
	"net/http"

	"github.com/Alfeenn/article/exception"
	"github.com/Alfeenn/article/model/web"
	"github.com/gin-gonic/gin"
)

var err interface{}

type ErrMiddleware struct {
	Exception exception.NotFound
}

func NewMiddleware() gin.HandlerFunc {

	return func(ctx *gin.Context) {
		if ctx.GetHeader("X-API-KEY") == "RAHASIA" {
			ctx.Next()

			exception.ErrHandler(ctx, err)

			return

		}

		response := web.WebResponse{
			Code:   http.StatusUnauthorized,
			Status: "UNAUTHORIZED",
		}

		ctx.AbortWithStatusJSON(http.StatusUnauthorized, response)
		ctx.Next()

	}
}
