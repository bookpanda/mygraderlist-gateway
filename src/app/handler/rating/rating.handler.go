package rating

import (
	"net/http"

	"github.com/bookpanda/mygraderlist-gateway/src/app/dto"
	"github.com/bookpanda/mygraderlist-gateway/src/app/router"
	"github.com/bookpanda/mygraderlist-gateway/src/app/validator"
	proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/backend/rating"
)

type Handler struct {
	service  IService
	validate *validator.DtoValidator
}

type IService interface {
	FindAll() ([]*proto.Rating, *dto.ResponseErr)
	FindByUserId(string) ([]*proto.Rating, *dto.ResponseErr)
	Create(*dto.RatingDto) (*proto.Rating, *dto.ResponseErr)
	Update(string, *dto.UpdateRatingDto) (*proto.Rating, *dto.ResponseErr)
	Delete(string) (bool, *dto.ResponseErr)
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
	ratingDto := dto.RatingDto{}

	err := c.Bind(&ratingDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	if errors := h.validate.Validate(ratingDto); errors != nil {
		c.JSON(http.StatusBadRequest, &dto.ResponseErr{
			StatusCode: http.StatusBadRequest,
			Message:    "Invalid body request",
			Data:       errors,
		})
		return
	}

	Rating, errRes := h.service.Create(&ratingDto)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusCreated, Rating)
	return
}

func (h *Handler) Update(c *router.FiberCtx) {
	id, err := c.ID()
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseErr{
			StatusCode: http.StatusInternalServerError,
			Message:    "Invalid ID",
			Data:       nil,
		})
		return
	}

	ratingDto := dto.UpdateRatingDto{}
	err = c.Bind(&ratingDto)
	if err != nil {
		c.JSON(http.StatusBadRequest, err)
		return
	}

	rating, errRes := h.service.Update(id, &ratingDto)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, rating)
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
