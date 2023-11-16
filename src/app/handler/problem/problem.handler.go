package problem

import (
	"net/http"

	"github.com/bookpanda/mygraderlist-gateway/src/app/dto"
	"github.com/bookpanda/mygraderlist-gateway/src/app/router"
	"github.com/bookpanda/mygraderlist-gateway/src/app/validator"
	proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/backend/problem"
)

type Handler struct {
	service  IService
	validate *validator.DtoValidator
}

type IService interface {
	FindAll() ([]*proto.Problem, *dto.ResponseErr)
}

func NewHandler(service IService, validate *validator.DtoValidator) *Handler {
	return &Handler{service, validate}
}

func (h *Handler) FindAll(c *router.FiberCtx) {
	result, err := h.service.FindAll()
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, result)
}
