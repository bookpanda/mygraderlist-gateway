package auth

import (
	"net/http"
	"os/user"

	"github.com/bookpanda/mygraderlist-gateway/src/app/dto"
	"github.com/bookpanda/mygraderlist-gateway/src/app/router"
	"github.com/bookpanda/mygraderlist-gateway/src/app/validator"
)

type Handler struct {
	service  IService
	validate *validator.DtoValidator
}

type IService interface {
	FindOne(string) (*user.User, *dto.ResponseErr)
	Update(string, *dto.User) (*user.User, *dto.ResponseErr)
	Delete(string) (bool, *dto.ResponseErr)
}

func NewHandler(service IService, validate *validator.DtoValidator) *Handler {
	return &Handler{service, validate}
}

func (h *Handler) FindOne(ctx *router.GinCtx) {
	id, err := ctx.ID()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseErr{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid ID",
			Data:       nil,
		})
		return
	}

	user, errRes := h.service.FindOne(id)
	if errRes != nil {
		ctx.JSON(errRes.StatusCode, errRes)
		return
	}

	ctx.JSON(http.StatusCreated, user)
	return
}

func (h *Handler) Update(ctx *router.GinCtx) {
	req := &dto.User{}
	err := ctx.Bind(&req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseErr{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while binding the request: " + err.Error(),
			Data:       nil,
		})
		return
	}

	id, err := ctx.ID()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, dto.ResponseErr{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid ID",
			Data:       nil,
		})
		return
	}

	user, errRes := h.service.Update(id, req)
	if errRes != nil {
		ctx.JSON(errRes.StatusCode, errRes)
		return
	}

	ctx.JSON(http.StatusCreated, user)
	return
}
