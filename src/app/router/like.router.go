package router

import (
	"github.com/gin-gonic/gin"
)

func (r *GinRouter) GetLike(path string, h func(ctx *GinCtx)) {
	r.like.GET(path, func(c *gin.Context) { h(NewGinCtx(c)) })
}

func (r *GinRouter) PostLike(path string, h func(ctx *GinCtx)) {
	r.like.POST(path, func(c *gin.Context) { h(NewGinCtx(c)) })
}

func (r *GinRouter) DeleteLike(path string, h func(ctx *GinCtx)) {
	r.like.DELETE(path, func(c *gin.Context) { h(NewGinCtx(c)) })
}
