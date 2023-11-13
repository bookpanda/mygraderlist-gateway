package user

import (
	"context"

	"github.com/bookpanda/mygraderlist-gateway/src/app/dto"
	user_proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/backend/user"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type ServiceMock struct {
	mock.Mock
}

func (s *ServiceMock) FindOne(id string) (result *user_proto.User, err *dto.ResponseErr) {
	args := s.Called(id)

	if args.Get(0) != nil {
		result = args.Get(0).(*user_proto.User)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *ServiceMock) Create(in *dto.UserDto) (result *user_proto.User, err *dto.ResponseErr) {
	args := s.Called(in)

	if args.Get(0) != nil {
		result = args.Get(0).(*user_proto.User)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *ServiceMock) Verify(id string, verifyType string) (result bool, err *dto.ResponseErr) {
	args := s.Called(id, verifyType)

	if args.Get(0) != nil {
		result = args.Bool(0)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *ServiceMock) Update(id string, in *dto.UpdateUserDto) (result *user_proto.User, err *dto.ResponseErr) {
	args := s.Called(id, in)

	if args.Get(0) != nil {
		result = args.Get(0).(*user_proto.User)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *ServiceMock) CreateOrUpdate(in *dto.UserDto) (result *user_proto.User, err *dto.ResponseErr) {
	args := s.Called(in)

	if args.Get(0) != nil {
		result = args.Get(0).(*user_proto.User)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

func (s *ServiceMock) Delete(id string) (result bool, err *dto.ResponseErr) {
	args := s.Called(id)

	result = args.Bool(0)

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return
}

type ClientMock struct {
	mock.Mock
}

func (c *ClientMock) Verify(_ context.Context, in *user_proto.VerifyUserRequest, _ ...grpc.CallOption) (res *user_proto.VerifyUserResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*user_proto.VerifyUserResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) FindOne(_ context.Context, in *user_proto.FindOneUserRequest, _ ...grpc.CallOption) (res *user_proto.FindOneUserResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*user_proto.FindOneUserResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) FindByEmail(_ context.Context, in *user_proto.FindByEmailUserRequest, _ ...grpc.CallOption) (res *user_proto.FindByEmailUserResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*user_proto.FindByEmailUserResponse)
	}

	return res, args.Error(1)

}

func (c *ClientMock) Create(_ context.Context, in *user_proto.CreateUserRequest, _ ...grpc.CallOption) (res *user_proto.CreateUserResponse, err error) {
	args := c.Called(in.User)

	if args.Get(0) != nil {
		res = args.Get(0).(*user_proto.CreateUserResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) Update(_ context.Context, in *user_proto.UpdateUserRequest, _ ...grpc.CallOption) (res *user_proto.UpdateUserResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*user_proto.UpdateUserResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) Delete(_ context.Context, in *user_proto.DeleteUserRequest, _ ...grpc.CallOption) (res *user_proto.DeleteUserResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*user_proto.DeleteUserResponse)
	}

	return res, args.Error(1)
}

type ContextMock struct {
	mock.Mock
	V      interface{}
	Status int
}

func (c *ContextMock) JSON(status int, v interface{}) {
	c.V = v
	c.Status = status
}

func (c *ContextMock) Bind(v interface{}) error {
	args := c.Called(v)

	if args.Get(0) != nil {
		switch v.(type) {
		case *dto.UserDto:
			*v.(*dto.UserDto) = *args.Get(0).(*dto.UserDto)
		case *dto.UpdateUserDto:
			*v.(*dto.UpdateUserDto) = *args.Get(0).(*dto.UpdateUserDto)
		}
	}

	return args.Error(1)
}

func (c *ContextMock) ID() (string, error) {
	args := c.Called()

	return args.String(0), args.Error(1)
}

func (c *ContextMock) Host() string {
	args := c.Called()

	return args.String(0)
}

func (c *ContextMock) UserID() string {
	args := c.Called()
	return args.String(0)
}

func (c *ContextMock) Query(key string) string {
	args := c.Called(key)
	return args.String(0)
}
