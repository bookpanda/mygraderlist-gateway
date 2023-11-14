package router

import (
	"github.com/gin-gonic/gin"
)

func (r *GinRouter) GetProblem(path string, h func(ctx *GinCtx)) {
	r.problem.GET(path, func(c *gin.Context) { h(NewGinCtx(c)) })
}
