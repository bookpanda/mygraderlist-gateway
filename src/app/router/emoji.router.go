package router

import (
	"github.com/gin-gonic/gin"
)

func (r *GinRouter) GetEmoji(path string, h func(ctx *GinCtx)) {
	r.emoji.GET(path, func(c *gin.Context) { h(NewGinCtx(c)) })
}

func (r *GinRouter) PostEmoji(path string, h func(ctx *GinCtx)) {
	r.emoji.POST(path, func(c *gin.Context) { h(NewGinCtx(c)) })
}

func (r *GinRouter) DeleteEmoji(path string, h func(ctx *GinCtx)) {
	r.emoji.DELETE(path, func(c *gin.Context) { h(NewGinCtx(c)) })
}
