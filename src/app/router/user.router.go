package router

import (
	"github.com/gin-gonic/gin"
)

func (r *GinRouter) GetUser(path string, h func(ctx *gin.Context)) {
	r.user.GET(path, h)
}

func (r *GinRouter) PostUser(path string, h func(ctx *gin.Context)) {
	r.user.POST(path, h)
}

func (r *GinRouter) PutUser(path string, h func(ctx *gin.Context)) {
	r.user.PUT(path, h)
}

func (r *GinRouter) DeleteUser(path string, h func(ctx *gin.Context)) {
	r.user.DELETE(path, h)
}
