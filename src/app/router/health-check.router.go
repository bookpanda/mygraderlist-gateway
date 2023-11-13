package router

import (
	"github.com/gin-gonic/gin"
)

func (r *GinRouter) GetHealthCheck(path string, h func(ctx *GinCtx)) {
	r.GET(path, func(c *gin.Context) {
		h(NewGinCtx(c))
	})
}
