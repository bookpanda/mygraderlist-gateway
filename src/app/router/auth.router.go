package router

import (
	"github.com/gin-gonic/gin"
)

func (r *GinRouter) GetAuth(path string, h func(ctx *GinCtx)) {
	r.auth.GET(path, func(c *gin.Context) { h(NewGinCtx(c)) })
}

func (r *GinRouter) PostAuth(path string, h func(ctx *GinCtx)) {
	r.auth.POST(path, func(c *gin.Context) { h(NewGinCtx(c)) })
}
