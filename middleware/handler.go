package middleware

import (
	"net/http"

	"github.com/Alfeenn/article/model/web"
	"github.com/gin-gonic/gin"
)

func NewMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		if ctx.GetHeader("X-API-KEY") == "RAHASIA" {

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
