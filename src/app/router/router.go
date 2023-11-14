package router

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type GinRouter struct {
	*gin.Engine
	user    *gin.RouterGroup
	auth    *gin.RouterGroup
	course  *gin.RouterGroup
	problem *gin.RouterGroup
	like    *gin.RouterGroup
	emoji   *gin.RouterGroup
	rating  *gin.RouterGroup
}

type IGuard interface {
	Use(*GinCtx)
}

func NewGinRouter(authGuard IGuard) *GinRouter {
	r := gin.Default()

	config := cors.DefaultConfig()
	config.AllowOrigins = []string{"http://localhost:3003"}
	r.Use(cors.New(config))

	user := GroupWithAuthMiddleware(r, "/user", authGuard.Use)
	auth := GroupWithAuthMiddleware(r, "/auth", authGuard.Use)
	course := GroupWithAuthMiddleware(r, "/course", authGuard.Use)
	problem := GroupWithAuthMiddleware(r, "/problem", authGuard.Use)
	like := GroupWithAuthMiddleware(r, "/like", authGuard.Use)
	emoji := GroupWithAuthMiddleware(r, "/emoji", authGuard.Use)
	rating := GroupWithAuthMiddleware(r, "/rating", authGuard.Use)

	return &GinRouter{r, user, auth, course, problem, like, emoji, rating}
}

func GroupWithAuthMiddleware(r *gin.Engine, path string, middleware func(ctx *GinCtx)) *gin.RouterGroup {
	return r.Group(path, func(c *gin.Context) {
		middleware(NewGinCtx(c))
	})
}
