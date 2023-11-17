package like

import (
	"net/http"

	"github.com/bookpanda/mygraderlist-gateway/src/app/dto"
	"github.com/bookpanda/mygraderlist-gateway/src/app/router"
	"github.com/bookpanda/mygraderlist-gateway/src/app/validator"
	proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/backend/like"
)

type Handler struct {
	service  IService
	validate *validator.DtoValidator
}

type IService interface {
	FindByUserId(string) ([]*proto.Like, *dto.ResponseErr)
	Create(*dto.LikeDto) (*proto.Like, *dto.ResponseErr)
	Delete(string) (bool, *dto.ResponseErr)
}

func NewHandler(service IService, validate *validator.DtoValidator) *Handler {
	return &Handler{service, validate}
}

func (h *Handler) FindByUserId(c *router.FiberCtx) {

	userId := c.UserID()

	result, errRes := h.service.FindByUserId(userId)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *Handler) Create(c *router.FiberCtx) {
	likeDto := dto.LikeDto{}

	err := c.Bind(&likeDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if errors := h.validate.Validate(likeDto); errors != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid body request",
			Data:       errors,
		})
		return
	}

	like, errRes := h.service.Create(&likeDto)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusCreated, like)
	return
}

func (h *Handler) Delete(c *router.FiberCtx) {
	id, err := c.ID()
	if err != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    err.Error(),
			Data:       nil,
		})
		return
	}

	success, errRes := h.service.Delete(id)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, success)
	return
}
