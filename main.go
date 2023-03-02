package main

import (
	"github.com/Alfeenn/article/middleware"
	"github.com/gin-gonic/gin"
)

func NewServer(authMiddleware middleware.Middleware) gin.RouteInfo {

	return gin.RouteInfo{
		Path: ":8000",
	}
}

func main() {

}
