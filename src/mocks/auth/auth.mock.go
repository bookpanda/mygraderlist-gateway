package auth

import (
	"context"

	"github.com/bookpanda/mygraderlist-gateway/src/app/dto"
	auth_proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/auth"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type ContextMock struct {
	mock.Mock
	V               interface{}
	Header          map[string]string
	RefreshTokenDto *dto.RedeemNewToken
}

func (c *ContextMock) Bind(v interface{}) error {
	args := c.Called(v)

	switch v.(type) {
	case *dto.RedeemNewToken:
		*v.(*dto.RedeemNewToken) = *c.RefreshTokenDto
	}

	return args.Error(1)
}

func (c *ContextMock) JSON(_ int, v interface{}) {
	c.V = v
}

func (c *ContextMock) UserID() string {
	args := c.Called()

	return args.String(0)
}

func (c *ContextMock) Token() string {
	args := c.Called()

	return args.String(0)
}

func (c *ContextMock) StoreValue(key string, val string) {
	_ = c.Called(key, val)

	c.Header[key] = val
}

func (c *ContextMock) Method() string {
	args := c.Called()

	return args.String(0)
}

func (c *ContextMock) Path() string {
	args := c.Called()

	return args.String(0)
}

func (c *ContextMock) Next() {
	_ = c.Called()

	return
}

type ClientMock struct {
	mock.Mock
}

func (c *ClientMock) Validate(_ context.Context, in *auth_proto.ValidateRequest, _ ...grpc.CallOption) (res *auth_proto.ValidateResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*auth_proto.ValidateResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) RefreshToken(_ context.Context, in *auth_proto.RefreshTokenRequest, _ ...grpc.CallOption) (res *auth_proto.RefreshTokenResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*auth_proto.RefreshTokenResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) GetGoogleLoginUrl(_ context.Context, in *auth_proto.GetGoogleLoginUrlRequest, _ ...grpc.CallOption) (res *auth_proto.GetGoogleLoginUrlResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*auth_proto.GetGoogleLoginUrlResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) VerifyGoogleLogin(_ context.Context, in *auth_proto.VerifyGoogleLoginRequest, _ ...grpc.CallOption) (res *auth_proto.VerifyGoogleLoginResponse, err error) {
	args := c.Called()

	if args.Get(0) != nil {
		res = args.Get(0).(*auth_proto.VerifyGoogleLoginResponse)
	}

	return res, args.Error(1)
}

type ServiceMock struct {
	mock.Mock
}

func (s *ServiceMock) Validate(token string) (payload *dto.TokenPayloadAuth, err *dto.ResponseErr) {
	args := s.Called(token)

	if args.Get(0) != nil {
		payload = args.Get(0).(*dto.TokenPayloadAuth)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return payload, err
}

func (s *ServiceMock) RefreshToken(token string) (credential *auth_proto.Credential, err *dto.ResponseErr) {
	args := s.Called(token)

	if args.Get(0) != nil {
		credential = args.Get(0).(*auth_proto.Credential)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return credential, err
}

func (s *ServiceMock) GetGoogleLoginUrl() (url string, err *dto.ResponseErr) {
	args := s.Called()

	if args.Get(0) != nil {
		url = args.Get(0).(string)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return url, err
}

func (s *ServiceMock) VerifyGoogleLogin(code string) (credential *auth_proto.Credential, err *dto.ResponseErr) {
	args := s.Called()

	if args.Get(0) != nil {
		credential = args.Get(0).(*auth_proto.Credential)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return credential, err
}
