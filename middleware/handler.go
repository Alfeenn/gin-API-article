package middleware

import (
	"github.com/gin-gonic/gin"
)

type Middleware struct {
	Handler gin.Engine
}

func NewMiddleware(handler gin.Engine) *Middleware {
	return &Middleware{
		Handler: handler,
	}
}
func (m *Middleware) ServeHTTP(c *gin.Context) {
	m.Handler.ServeHTTP(c.Writer, c.Request)
}
