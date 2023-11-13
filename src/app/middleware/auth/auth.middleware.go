package auth

import (
	"net/http"

	"github.com/bookpanda/mygraderlist-gateway/src/app/dto"
	"github.com/bookpanda/mygraderlist-gateway/src/app/handler/auth"
	"github.com/bookpanda/mygraderlist-gateway/src/app/utils"
	"github.com/bookpanda/mygraderlist-gateway/src/config"
)

type Guard struct {
	service    auth.IService
	excludes   map[string]struct{}
	conf       config.App
	isValidate bool
}

type IContext interface {
	Token() string
	Method() string
	Path() string
	StoreValue(string, string)
	JSON(int, interface{})
	Next()
}

func NewAuthGuard(s auth.IService, e map[string]struct{}, conf config.App) Guard {
	return Guard{
		service:    s,
		excludes:   e,
		conf:       conf,
		isValidate: true,
	}
}

func (m *Guard) Use(ctx IContext) {
	m.isValidate = true

	m.Validate(ctx)

	if !m.isValidate {
		return
	}

	ctx.Next()

}

func (m *Guard) Validate(ctx IContext) {
	method := ctx.Method()
	path := ctx.Path()

	ids := utils.FindIDFromPath(path)

	path = utils.FormatPath(method, path, ids)
	if utils.IsExisted(m.excludes, path) {
		ctx.Next()
		return
	}

	token := ctx.Token()
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, &dto.ResponseErr{
			StatusCode: http.StatusUnauthorized,
			Message:    "Invalid token",
		})
		m.isValidate = false
		return
	}

	payload, errRes := m.service.Validate(token)
	if errRes != nil {
		ctx.JSON(errRes.StatusCode, errRes)
		m.isValidate = false
		return
	}

	ctx.StoreValue("UserId", payload.UserId)
	ctx.Next()
}
