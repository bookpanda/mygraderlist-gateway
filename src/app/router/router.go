package router

import (
	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	*gin.Engine
	user *gin.RouterGroup
	auth *gin.RouterGroup
}

func NewGinRouter() *GinRouter {
	r := gin.Default()

	user := r.Group("/user")
	// user.Use(middleware.AuthMiddleware())
	auth := r.Group("/auth")

	return &GinRouter{r, user, auth}
}
