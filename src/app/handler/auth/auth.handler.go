package auth

import (
	"net/http"

	"github.com/bookpanda/mygraderlist-gateway/src/app/dto"
	"github.com/bookpanda/mygraderlist-gateway/src/app/router"
	"github.com/bookpanda/mygraderlist-gateway/src/app/validator"
	auth_proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/auth"
	user_proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/backend/user"
)

type Handler struct {
	service    IService
	usrService IUserService
	validate   *validator.DtoValidator
}

type IService interface {
	Validate(string) (*dto.TokenPayloadAuth, *dto.ResponseErr)
	RefreshToken(string) (*auth_proto.Credential, *dto.ResponseErr)
	GetGoogleLoginUrl() (string, *dto.ResponseErr)
	VerifyGoogleLogin(string) (*auth_proto.Credential, *dto.ResponseErr)
}

type IUserService interface {
	FindOne(string) (*user_proto.User, *dto.ResponseErr)
}

func NewHandler(service IService, usrService IUserService, validate *validator.DtoValidator) *Handler {
	return &Handler{service, usrService, validate}
}

func (h *Handler) Validate(c *router.FiberCtx) {
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

func (h *Handler) RefreshToken(c *router.FiberCtx) {
	refreshToken := dto.RedeemNewToken{}

	err := c.Bind(&refreshToken)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseErr{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while binding the token: " + err.Error(),
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

func (h *Handler) GetGoogleLoginUrl(c *router.FiberCtx) {
	url, errRes := h.service.GetGoogleLoginUrl()
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, url)
}

func (h *Handler) VerifyGoogleLogin(c *router.FiberCtx) {
	code := dto.VerifyGoogle{}
	err := c.Bind(&code)
	if err != nil {
		c.JSON(http.StatusInternalServerError, dto.ResponseErr{
			StatusCode: http.StatusInternalServerError,
			Message:    "Error while binding the code: " + err.Error(),
			Data:       nil,
		})
		return
	}

	credential, errRes := h.service.VerifyGoogleLogin(code.Code)
	if errRes != nil {
		c.JSON(errRes.StatusCode, errRes)
		return
	}

	c.JSON(http.StatusOK, credential)
}
