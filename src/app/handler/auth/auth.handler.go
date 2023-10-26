package auth

import (
	"net/http"

	"github.com/bookpanda/mygraderlist-gateway/src/app/dto"
	"github.com/bookpanda/mygraderlist-gateway/src/app/router"
	"github.com/bookpanda/mygraderlist-gateway/src/app/validator"
	"github.com/bookpanda/mygraderlist-gateway/src/proto"
)

type Handler struct {
	service    IService
	usrService IUserService
	validate   *validator.DtoValidator
}

type IService interface {
	// Signup(req *dto.Signup) (*token.TokenPair, *dto.ResponseErr)
	// Signin(req *dto.Signin) (*token.TokenPair, *dto.ResponseErr)
	Validate(string) (*dto.TokenPayloadAuth, *dto.ResponseErr)
	RefreshToken(string) (*proto.Credential, *dto.ResponseErr)
}

type IUserService interface {
	FindOne(string) (*proto.User, *dto.ResponseErr)
}

func NewHandler(service IService, usrService IUserService, validate *validator.DtoValidator) *Handler {
	return &Handler{service, usrService, validate}
}

func (h *Handler) Validate(c *router.GinCtx) {
	userId := c.UserID()

	usr, err := h.usrService.FindOne(userId)
	if err != nil {
		switch err.StatusCode {
		case http.StatusNotFound:
			c.JSON(http.StatusUnauthorized, &dto.ResponseErr{
				StatusCode: http.StatusUnauthorized,
				Message:    "Invalid user",
			})
		default:
			c.JSON(err.StatusCode, err)
		}
		return
	}

	c.JSON(http.StatusOK, usr)
}

func (h *Handler) RefreshToken(c *router.GinCtx) {
	refreshToken := dto.RedeemNewToken{}

	err := c.Bind(&refreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseErr{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while binding the ticket: " + err.Error(),
			Data:       nil,
		})
		return
	}

	credential, errRes := h.service.RefreshToken(refreshToken.RefreshToken)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, credential)
}

// func (h *Handler) Signup(ctx *router.GinCtx) {
// 	signupDto := &dto.Signup{}

// 	err := ctx.Bind(&signupDto)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, err)
// 		return
// 	}

// 	if errors := h.validate.Validate(signupDto); errors != nil {
// 		ctx.JSON(http.StatusBadRequest, &dto.ResponseErr{
// 			StatusCode: http.StatusBadRequest,
// 			Message:    "Invalid request body",
// 			Data:       errors,
// 		})
// 		return
// 	}

// 	tokens, errRes := h.service.Signup(signupDto)
// 	if errRes != nil {
// 		ctx.JSON(errRes.StatusCode, errRes)
// 		return
// 	}

// 	ctx.JSON(http.StatusCreated, gin.H{
// 		"tokens": tokens,
// 	})
// }

// func (h *Handler) Signin(ctx *router.GinCtx) {
// 	signinDto := &dto.Signin{}

// 	err := ctx.Bind(&signinDto)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadRequest, err)
// 		return
// 	}

// 	if errors := h.validate.Validate(signinDto); errors != nil {
// 		ctx.JSON(http.StatusBadRequest, &dto.ResponseErr{
// 			StatusCode: http.StatusBadRequest,
// 			Message:    "Invalid request body",
// 			Data:       errors,
// 		})
// 		return
// 	}

// 	tokens, errRes := h.service.Signin(signinDto)
// 	if errRes != nil {
// 		ctx.JSON(errRes.StatusCode, errRes)
// 		return
// 	}

// 	ctx.JSON(http.StatusCreated, gin.H{
// 		"tokens": tokens,
// 	})
// }
