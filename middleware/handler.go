package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

var handler http.Handler

func NewMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		handler.ServeHTTP(ctx.Writer, ctx.Request)

	}
}
