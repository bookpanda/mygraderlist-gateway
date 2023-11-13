package router

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type GinCtx struct {
	*gin.Context
}

func NewGinCtx(c *gin.Context) *GinCtx {
	return &GinCtx{c}
}

func (c *GinCtx) UserID() string {
	return c.GetString("UserId")
}

func (c *GinCtx) Bind(v interface{}) error {
	return c.ShouldBind(v)
}

func (c *GinCtx) ID() (id string, err error) {
	id = c.Param("id")

	_, err = uuid.Parse(id)
	if err != nil {
		return "", err
	}

	return id, nil
}

func (c *GinCtx) Token() string {
	return c.GetHeader("Authorization")
}

func (c *GinCtx) Method() string {
	return c.Method()
}

func (c *GinCtx) Path() string {
	return c.Path()
}

func (c *GinCtx) Params(key string) (value string, err error) {
	value = c.Param(key)

	if key == "id" {
		_, err = uuid.Parse(value)
		if err != nil {
			return "", err
		}
	}

	return value, nil
}

func (c *GinCtx) StoreValue(k string, v string) {
	c.Set(k, v)
}
