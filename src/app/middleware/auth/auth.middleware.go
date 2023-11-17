package auth

import (
	"net/http"

	"github.com/bookpanda/mygraderlist-gateway/src/app/dto"
	"github.com/bookpanda/mygraderlist-gateway/src/app/handler/auth"
	"github.com/bookpanda/mygraderlist-gateway/src/app/router"
	"github.com/bookpanda/mygraderlist-gateway/src/app/utils"
	"github.com/bookpanda/mygraderlist-gateway/src/config"
)

type Guard struct {
	service    auth.IService
	excludes   map[string]struct{}
	conf       config.App
	isValidate bool
}

func NewAuthGuard(s auth.IService, e map[string]struct{}, conf config.App) Guard {
	return Guard{
		service:    s,
		excludes:   e,
		conf:       conf,
		isValidate: true,
	}
}

func (m *Guard) Use(ctx *router.FiberCtx) {
	m.isValidate = true

	m.Validate(ctx)

	if !m.isValidate {
		return
	}

	ctx.Next()

}

func (m *Guard) Validate(ctx *router.FiberCtx) {
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
