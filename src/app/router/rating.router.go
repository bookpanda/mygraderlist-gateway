package router

import (
	"github.com/gin-gonic/gin"
)

func (r *GinRouter) GetRating(path string, h func(ctx *GinCtx)) {
	r.rating.GET(path, func(c *gin.Context) { h(NewGinCtx(c)) })
}

func (r *GinRouter) PostRating(path string, h func(ctx *GinCtx)) {
	r.rating.POST(path, func(c *gin.Context) { h(NewGinCtx(c)) })
}

func (r *GinRouter) PutRating(path string, h func(ctx *GinCtx)) {
	r.rating.PUT(path, func(c *gin.Context) { h(NewGinCtx(c)) })
}

func (r *GinRouter) DeleteRating(path string, h func(ctx *GinCtx)) {
	r.rating.DELETE(path, func(c *gin.Context) { h(NewGinCtx(c)) })
}
