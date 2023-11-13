package router

import (
	"github.com/gin-gonic/gin"
)

func (r *GinRouter) GetUser(path string, h func(ctx *GinCtx)) {
	r.user.GET(path, func(c *gin.Context) { h(NewGinCtx(c)) })
}

func (r *GinRouter) PostUser(path string, h func(ctx *GinCtx)) {
	r.user.POST(path, func(c *gin.Context) { h(NewGinCtx(c)) })
}

func (r *GinRouter) PutUser(path string, h func(ctx *GinCtx)) {
	r.user.PUT(path, func(c *gin.Context) { h(NewGinCtx(c)) })
}

func (r *GinRouter) DeleteUser(path string, h func(ctx *GinCtx)) {
	r.user.DELETE(path, func(c *gin.Context) { h(NewGinCtx(c)) })
}
