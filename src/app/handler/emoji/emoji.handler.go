package emoji

import (
	"net/http"

	"github.com/bookpanda/mygraderlist-gateway/src/app/dto"
	"github.com/bookpanda/mygraderlist-gateway/src/app/router"
	"github.com/bookpanda/mygraderlist-gateway/src/app/validator"
	proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/backend/emoji"
)

type Handler struct {
	service  IService
	validate *validator.DtoValidator
}

type IService interface {
	FindAll() ([]*proto.Emoji, *dto.ResponseErr)
	FindByUserId(string) ([]*proto.Emoji, *dto.ResponseErr)
	Create(*dto.EmojiDto) (*proto.Emoji, *dto.ResponseErr)
	Delete(string) (bool, *dto.ResponseErr)
}

func NewHandler(service IService, validate *validator.DtoValidator) *Handler {
	return &Handler{service, validate}
}

func (h *Handler) FindAll(c *router.GinCtx) {
	result, err := h.service.FindAll()
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *Handler) FindByUserId(c *router.GinCtx) {
	userId := c.UserID()

	result, err := h.service.FindByUserId(userId)
	if err != nil {
		c.JSON(err.StatusCode, err)
		return
	}

	c.JSON(http.StatusOK, result)
}

func (h *Handler) Create(c *router.GinCtx) {
	emojiDto := dto.EmojiDto{}

	err := c.Bind(&emojiDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if errors := h.validate.Validate(emojiDto); errors != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid body request",
			Data:       errors,
		})
		return
	}

	emoji, errRes := h.service.Create(&emojiDto)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusCreated, emoji)
	return
}

func (h *Handler) Delete(c *router.GinCtx) {
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
