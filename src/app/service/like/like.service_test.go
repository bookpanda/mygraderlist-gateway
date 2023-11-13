package like

import (
	"net/http"
	"testing"

	"github.com/bookpanda/mygraderlist-gateway/src/app/dto"
	"github.com/bookpanda/mygraderlist-gateway/src/mocks/like"
	proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/backend/like"
	user_proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/backend/user"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type LikeServiceTest struct {
	suite.Suite
	User           *user_proto.User
	Like           *proto.Like
	Likes          []*proto.Like
	LikeDto        *dto.LikeDto
	NotFoundErr    *dto.ResponseErr
	ServiceDownErr *dto.ResponseErr
}

func TestGroupService(t *testing.T) {
	suite.Run(t, new(LikeServiceTest))
}

func (t *LikeServiceTest) SetupTest() {
	t.User = &user_proto.User{
		Id:       faker.UUIDDigit(),
		Username: faker.Username(),
		Email:    faker.Email(),
		Password: faker.Password(),
	}

	t.LikeDto = &dto.LikeDto{
		ProblemID: uuid.New(),
		UserID:    uuid.New(),
	}

	t.Like = &proto.Like{
		Id:        faker.UUIDDigit(),
		ProblemId: t.LikeDto.ProblemID.String(),
		UserId:    t.LikeDto.UserID.String(),
	}

	Like2 := &proto.Like{
		Id:        faker.UUIDDigit(),
		ProblemId: uuid.New().String(),
		UserId:    uuid.New().String(),
	}

	Like3 := &proto.Like{
		Id:        faker.UUIDDigit(),
		ProblemId: uuid.New().String(),
		UserId:    uuid.New().String(),
	}
	t.Likes = make([]*proto.Like, 0)
	t.Likes = append(t.Likes, t.Like, Like2, Like3)

	t.ServiceDownErr = &dto.ResponseErr{
		StatusCode: http.StatusServiceUnavailable,
		Message:    "Service is down",
		Data:       nil,
	}

	t.NotFoundErr = &dto.ResponseErr{
		StatusCode: http.StatusNotFound,
		Message:    "Like not found",
		Data:       nil,
	}
}

func (t *LikeServiceTest) TestFindByUserIdSuccess() {
	want := t.Likes

	c := &like.ClientMock{}
	c.On("FindByUserId", &proto.FindByUserIdLikeRequest{UserId: t.User.Id}).Return(&proto.FindByUserIdLikeResponse{Likes: t.Likes}, nil)

	srv := NewService(c)
	actual, err := srv.FindByUserId(t.User.Id)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *LikeServiceTest) TestFindByUserIdNotFound() {
	want := t.NotFoundErr

	c := &like.ClientMock{}
	c.On("FindByUserId", &proto.FindByUserIdLikeRequest{UserId: t.User.Id}).Return(nil, status.Error(codes.NotFound, "Like not found"))

	srv := NewService(c)
	actual, err := srv.FindByUserId(t.User.Id)

	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), want, err)
}

func (t *LikeServiceTest) TestFindByUserIdGrpcErr() {
	want := t.ServiceDownErr

	c := &like.ClientMock{}
	c.On("FindByUserId", &proto.FindByUserIdLikeRequest{UserId: t.User.Id}).Return(nil, errors.New("Service is down"))

	srv := NewService(c)
	actual, err := srv.FindByUserId(t.User.Id)

	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), want, err)
}

func (t *LikeServiceTest) TestCreateSuccess() {
	want := t.Like

	in := &proto.Like{
		ProblemId: t.LikeDto.ProblemID.String(),
		UserId:    t.LikeDto.UserID.String(),
	}

	c := &like.ClientMock{}
	c.On("Create", &proto.CreateLikeRequest{Like: in}).Return(&proto.CreateLikeResponse{Like: t.Like}, nil)

	srv := NewService(c)
	actual, err := srv.Create(t.LikeDto)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *LikeServiceTest) TestCreateGrpcErr() {
	want := t.ServiceDownErr

	in := &proto.Like{
		ProblemId: t.LikeDto.ProblemID.String(),
		UserId:    t.LikeDto.UserID.String(),
	}

	c := &like.ClientMock{}
	c.On("Create", &proto.CreateLikeRequest{Like: in}).Return(nil, errors.New("Service is down"))

	srv := NewService(c)
	actual, err := srv.Create(t.LikeDto)

	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), want, err)
}

func (t *LikeServiceTest) TestDeleteSuccess() {
	want := true

	c := &like.ClientMock{}
	c.On("Delete", &proto.DeleteLikeRequest{Id: t.Like.Id}).Return(&proto.DeleteLikeResponse{Success: true}, nil)

	srv := NewService(c)
	actual, err := srv.Delete(t.Like.Id)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *LikeServiceTest) TestDeleteNotFound() {
	want := t.NotFoundErr

	c := &like.ClientMock{}
	c.On("Delete", &proto.DeleteLikeRequest{Id: t.Like.Id}).Return(nil, status.Error(codes.NotFound, "Like not found"))

	srv := NewService(c)
	actual, err := srv.Delete(t.Like.Id)

	assert.False(t.T(), actual)
	assert.Equal(t.T(), want, err)
}

func (t *LikeServiceTest) TestDeleteGrpcErr() {
	want := t.ServiceDownErr

	c := &like.ClientMock{}
	c.On("Delete", &proto.DeleteLikeRequest{Id: t.Like.Id}).Return(nil, errors.New("Service is down"))

	srv := NewService(c)
	actual, err := srv.Delete(t.Like.Id)

	assert.False(t.T(), actual)
	assert.Equal(t.T(), want, err)
}
