package router

import (
	"github.com/gin-gonic/gin"
)

func (r *GinRouter) GetCourse(path string, h func(ctx *GinCtx)) {
	r.course.GET(path, func(c *gin.Context) { h(NewGinCtx(c)) })
}
