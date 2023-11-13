package router

import (
	"github.com/bookpanda/mygraderlist-gateway/src/config"
	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	*gin.Engine
	user   *gin.RouterGroup
	auth   *gin.RouterGroup
	like   *gin.RouterGroup
	emoji  *gin.RouterGroup
	rating *gin.RouterGroup
}

type IGuard interface {
	Use(*GinCtx)
}

func NewGinRouter(authGuard IGuard, conf config.App) *GinRouter {
	r := gin.Default()

	user := r.Group("/user")
	// user.Use(middleware.AuthMiddleware())
	auth := r.Group("/auth")

	like := r.Group("/like")
	emoji := r.Group("/emoji")
	rating := r.Group("/rating")

	return &GinRouter{r, user, auth, like, emoji, rating}
}
