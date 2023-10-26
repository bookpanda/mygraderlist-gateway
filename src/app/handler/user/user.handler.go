package auth

import (
	"net/http"

	"github.com/bookpanda/mygraderlist-gateway/src/app/dto"
	"github.com/bookpanda/mygraderlist-gateway/src/app/router"
	"github.com/bookpanda/mygraderlist-gateway/src/app/validator"
	"github.com/bookpanda/mygraderlist-gateway/src/proto"
)

type Handler struct {
	service  IService
	validate *validator.DtoValidator
}

type IService interface {
	FindOne(string) (*proto.User, *dto.ResponseErr)
	Create(*dto.UserDto) (*proto.User, *dto.ResponseErr)
	Update(string, *dto.UpdateUserDto) (*proto.User, *dto.ResponseErr)
	Delete(string) (bool, *dto.ResponseErr)
	Verify(string, string) (bool, *dto.ResponseErr)
}

func NewHandler(service IService, validate *validator.DtoValidator) *Handler {
	return &Handler{service, validate}
}

func (h *Handler) FindOne(c *router.GinCtx) {
	id, err := c.ID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseErr{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid ID",
			Data:       nil,
		})
		return
	}

	user, errRes := h.service.FindOne(id)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusCreated, user)
	return
}

func (h *Handler) Create(c *router.GinCtx) {
	usrDto := dto.UserDto{}

	err := c.Bind(&usrDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if errors := h.validate.Validate(usrDto); errors != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid body request",
			Data:       errors,
		})
		return
	}

	user, errRes := h.service.Create(&usrDto)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusCreated, user)
	return
}

func (h *Handler) Update(c router.GinCtx) {
	usrId := c.UserID()

	usrDto := dto.UpdateUserDto{}

	err := c.Bind(&usrDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	user, errRes := h.service.Update(usrId, &usrDto)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, user)
	return
}

func (h *Handler) Delete(c router.GinCtx) {
	id, err := c.ID()
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
		return
	}

	user, errRes := h.service.Delete(id)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, user)
	return
}
