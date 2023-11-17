package user

import (
	"net/http"
	"testing"

	"github.com/bookpanda/mygraderlist-gateway/src/app/dto"
	"github.com/bookpanda/mygraderlist-gateway/src/mocks/user"
	proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/backend/user"
	"github.com/bxcodec/faker/v3"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type UserServiceTest struct {
	suite.Suite
	User           *proto.User
	UserReq        *proto.User
	UserDto        *dto.UserDto
	UpdateUserDto  *dto.UpdateUserDto
	UpdateUserReq  *proto.UpdateUserRequest
	NotFoundErr    *dto.ResponseErr
	ServiceDownErr *dto.ResponseErr
	InternalErr    *dto.ResponseErr
}

func TestUserService(t *testing.T) {
	suite.Run(t, new(UserServiceTest))
}

func (t *UserServiceTest) SetupTest() {
	t.User = &proto.User{
		Id:       faker.UUIDDigit(),
		Username: faker.Name(),
		Password: faker.Name(),
		Email:    faker.Email(),
	}

	t.UserReq = &proto.User{
		Username: t.User.Username,
		Password: t.User.Password,
		Email:    t.User.Email,
	}

	t.UserDto = &dto.UserDto{
		Email:    t.User.Email,
		Username: t.User.Username,
		Password: t.User.Password,
	}

	t.UpdateUserDto = &dto.UpdateUserDto{
		Username: t.User.Username,
		Password: t.User.Password,
	}

	t.UpdateUserReq = &proto.UpdateUserRequest{
		Username: t.User.Username,
		Password: t.User.Password,
	}

	t.ServiceDownErr = &dto.ResponseErr{
		StatusCode: http.StatusServiceUnavailable,
		Message:    "Service is down",
		Data:       nil,
	}

	t.NotFoundErr = &dto.ResponseErr{
		StatusCode: http.StatusNotFound,
		Message:    "User not found",
		Data:       nil,
	}

	t.InternalErr = &dto.ResponseErr{
		StatusCode: http.StatusInternalServerError,
		Message:    "Internal Server Error",
		Data:       nil,
	}
}

func (t *UserServiceTest) TestFindOneSuccess() {
	want := t.User

	c := &user.ClientMock{}
	c.On("FindOne", &proto.FindOneUserRequest{Id: t.User.Id}).Return(&proto.FindOneUserResponse{User: want}, nil)
	srv := NewService(c)

	actual, err := srv.FindOne(t.User.Id)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *UserServiceTest) TestFindOneNotFound() {
	want := t.NotFoundErr

	c := &user.ClientMock{}
	c.On("FindOne", &proto.FindOneUserRequest{Id: t.User.Id}).Return(nil, status.Error(codes.NotFound, "User not found"))

	srv := NewService(c)

	actual, err := srv.FindOne(t.User.Id)

	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), want, err)
}

func (t *UserServiceTest) TestFindOneGrpcErr() {
	want := t.ServiceDownErr

	c := &user.ClientMock{}
	c.On("FindOne", &proto.FindOneUserRequest{Id: t.User.Id}).Return(nil, errors.New("Service is down"))
	srv := NewService(c)

	actual, err := srv.FindOne(t.User.Id)

	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), want, err)
}

func (t *UserServiceTest) TestCreateSuccess() {
	want := t.User

	c := &user.ClientMock{}
	c.On("Create", t.UserReq).Return(&proto.CreateUserResponse{User: want}, nil)

	srv := NewService(c)

	actual, err := srv.Create(t.UserDto)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *UserServiceTest) TestCreateGrpcErr() {
	want := t.ServiceDownErr

	c := &user.ClientMock{}
	c.On("Create", t.UserReq).Return(nil, errors.New("Service is down"))

	srv := NewService(c)

	actual, err := srv.Create(t.UserDto)

	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), want, err)
}

func (t *UserServiceTest) TestUpdateSuccess() {
	want := t.User

	c := &user.ClientMock{}
	c.On("Update", t.UpdateUserReq).Return(&proto.UpdateUserResponse{User: t.User}, nil)

	srv := NewService(c)

	actual, err := srv.Update(t.User.Id, t.UpdateUserDto)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *UserServiceTest) TestUpdateNotFound() {
	want := t.NotFoundErr

	c := &user.ClientMock{}
	c.On("Update", t.UpdateUserReq).Return(nil, status.Error(codes.NotFound, "User not found"))

	srv := NewService(c)

	actual, err := srv.Update(t.User.Id, t.UpdateUserDto)

	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), want, err)
}

func (t *UserServiceTest) TestUpdateGrpcErr() {
	want := t.ServiceDownErr

	c := &user.ClientMock{}
	c.On("Update", t.UpdateUserReq).Return(nil, errors.New("Service is down"))

	srv := NewService(c)

	actual, err := srv.Update(t.User.Id, t.UpdateUserDto)

	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), want, err)
}

func (t *UserServiceTest) TestDeleteSuccess() {
	c := &user.ClientMock{}
	c.On("Delete", &proto.DeleteUserRequest{Id: t.User.Id}).Return(&proto.DeleteUserResponse{Success: true}, nil)

	srv := NewService(c)

	actual, err := srv.Delete(t.User.Id)

	assert.Nil(t.T(), err)
	assert.True(t.T(), actual)
}

func (t *UserServiceTest) TestDeleteNotFound() {
	want := t.NotFoundErr

	c := &user.ClientMock{}
	c.On("Delete", &proto.DeleteUserRequest{Id: t.User.Id}).Return(nil, status.Error(codes.NotFound, "User not found"))

	srv := NewService(c)

	actual, err := srv.Delete(t.User.Id)

	assert.False(t.T(), actual)
	assert.Equal(t.T(), want, err)
}

func (t *UserServiceTest) TestDeleteGrpcErr() {
	want := t.ServiceDownErr

	c := &user.ClientMock{}
	c.On("Delete", &proto.DeleteUserRequest{Id: t.User.Id}).Return(nil, errors.New("Service is down"))

	srv := NewService(c)

	actual, err := srv.Delete(t.User.Id)

	assert.False(t.T(), actual)
	assert.Equal(t.T(), want, err)
}
