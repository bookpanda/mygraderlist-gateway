package rating

import (
	"net/http"
	"testing"

	"github.com/bookpanda/mygraderlist-gateway/src/app/dto"
	"github.com/bookpanda/mygraderlist-gateway/src/mocks/rating"
	proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/backend/rating"
	user_proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/backend/user"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type RatingServiceTest struct {
	suite.Suite
	User            *user_proto.User
	Rating          *proto.Rating
	Ratings         []*proto.Rating
	RatingDto       *dto.RatingDto
	UpdateRatingDto *dto.UpdateRatingDto
	NotFoundErr     *dto.ResponseErr
	ServiceDownErr  *dto.ResponseErr
}

func TestRatingService(t *testing.T) {
	suite.Run(t, new(RatingServiceTest))
}

func (t *RatingServiceTest) SetupTest() {
	t.User = &user_proto.User{
		Id:       faker.UUIDDigit(),
		Username: faker.Username(),
		Email:    faker.Email(),
		Password: faker.Password(),
	}

	t.RatingDto = &dto.RatingDto{
		Score:      2,
		Difficulty: 3,
		ProblemID:  uuid.New(),
		UserID:     uuid.New(),
	}

	t.Rating = &proto.Rating{
		Id:         faker.UUIDDigit(),
		Score:      int32(t.RatingDto.Score),
		Difficulty: int32(t.RatingDto.Difficulty),
		ProblemId:  t.RatingDto.ProblemID.String(),
		UserId:     t.RatingDto.UserID.String(),
	}

	Rating2 := &proto.Rating{
		Id:         faker.UUIDDigit(),
		Score:      1,
		Difficulty: 7,
		ProblemId:  uuid.New().String(),
		UserId:     uuid.New().String(),
	}

	Rating3 := &proto.Rating{
		Id:         faker.UUIDDigit(),
		Score:      3,
		Difficulty: 1,
		ProblemId:  uuid.New().String(),
		UserId:     uuid.New().String(),
	}
	t.Ratings = make([]*proto.Rating, 0)
	t.Ratings = append(t.Ratings, t.Rating, Rating2, Rating3)

	t.UpdateRatingDto = &dto.UpdateRatingDto{
		Score:      t.RatingDto.Score,
		Difficulty: t.RatingDto.Difficulty,
	}

	t.ServiceDownErr = &dto.ResponseErr{
		StatusCode: http.StatusServiceUnavailable,
		Message:    "Service is down",
		Data:       nil,
	}

	t.NotFoundErr = &dto.ResponseErr{
		StatusCode: http.StatusNotFound,
		Message:    "Rating not found",
		Data:       nil,
	}
}

func (t *RatingServiceTest) TestFindAllSuccess() {
	want := t.Ratings

	c := &rating.ClientMock{}
	c.On("FindAll", &proto.FindAllRatingRequest{}).Return(&proto.FindAllRatingResponse{Ratings: t.Ratings}, nil)

	srv := NewService(c)
	actual, err := srv.FindAll()

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *RatingServiceTest) TestFindAllGrpcErr() {
	want := t.ServiceDownErr

	c := &rating.ClientMock{}
	c.On("FindAll", &proto.FindAllRatingRequest{}).Return(nil, errors.New("Service is down"))

	srv := NewService(c)
	actual, err := srv.FindAll()

	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), want, err)
}

func (t *RatingServiceTest) TestFindByUserIdSuccess() {
	want := t.Ratings

	c := &rating.ClientMock{}
	c.On("FindByUserId", &proto.FindByUserIdRatingRequest{UserId: t.User.Id}).Return(&proto.FindByUserIdRatingResponse{Ratings: t.Ratings}, nil)

	srv := NewService(c)
	actual, err := srv.FindByUserId(t.User.Id)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *RatingServiceTest) TestFindByUserIdNotFound() {
	want := t.NotFoundErr

	c := &rating.ClientMock{}
	c.On("FindByUserId", &proto.FindByUserIdRatingRequest{UserId: t.User.Id}).Return(nil, status.Error(codes.NotFound, "Rating not found"))

	srv := NewService(c)
	actual, err := srv.FindByUserId(t.User.Id)

	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), want, err)
}

func (t *RatingServiceTest) TestFindByUserIdGrpcErr() {
	want := t.ServiceDownErr

	c := &rating.ClientMock{}
	c.On("FindByUserId", &proto.FindByUserIdRatingRequest{UserId: t.User.Id}).Return(nil, errors.New("Service is down"))

	srv := NewService(c)
	actual, err := srv.FindByUserId(t.User.Id)

	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), want, err)
}

func (t *RatingServiceTest) TestUpdateSuccess() {
	want := t.Rating

	c := &rating.ClientMock{}
	c.On("Update", &proto.UpdateRatingRequest{Id: t.Rating.Id, Score: t.Rating.Score, Difficulty: t.Rating.Difficulty}).Return(&proto.UpdateRatingResponse{Rating: t.Rating}, nil)

	srv := NewService(c)

	actual, err := srv.Update(t.Rating.Id, t.UpdateRatingDto)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *RatingServiceTest) TestUpdateNotFound() {
	want := t.NotFoundErr

	c := &rating.ClientMock{}
	c.On("Update", &proto.UpdateRatingRequest{Id: t.Rating.Id, Score: t.Rating.Score, Difficulty: t.Rating.Difficulty}).Return(nil, status.Error(codes.NotFound, "Rating not found"))

	srv := NewService(c)

	actual, err := srv.Update(t.Rating.Id, t.UpdateRatingDto)

	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), want, err)
}

func (t *RatingServiceTest) TestUpdateGrpcErr() {
	want := t.ServiceDownErr

	c := &rating.ClientMock{}
	c.On("Update", &proto.UpdateRatingRequest{Id: t.Rating.Id, Score: t.Rating.Score, Difficulty: t.Rating.Difficulty}).Return(nil, errors.New("Service is down"))

	srv := NewService(c)

	actual, err := srv.Update(t.Rating.Id, t.UpdateRatingDto)

	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), want, err)
}

func (t *RatingServiceTest) TestCreateSuccess() {
	want := t.Rating

	in := &proto.Rating{
		Score:      t.Rating.Score,
		Difficulty: t.Rating.Difficulty,
		ProblemId:  t.RatingDto.ProblemID.String(),
		UserId:     t.RatingDto.UserID.String(),
	}

	c := &rating.ClientMock{}
	c.On("Create", &proto.CreateRatingRequest{Rating: in}).Return(&proto.CreateRatingResponse{Rating: t.Rating}, nil)

	srv := NewService(c)
	actual, err := srv.Create(t.RatingDto)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *RatingServiceTest) TestCreateGrpcErr() {
	want := t.ServiceDownErr

	in := &proto.Rating{
		Score:      t.Rating.Score,
		Difficulty: t.Rating.Difficulty,
		ProblemId:  t.RatingDto.ProblemID.String(),
		UserId:     t.RatingDto.UserID.String(),
	}

	c := &rating.ClientMock{}
	c.On("Create", &proto.CreateRatingRequest{Rating: in}).Return(nil, errors.New("Service is down"))

	srv := NewService(c)
	actual, err := srv.Create(t.RatingDto)

	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), want, err)
}

func (t *RatingServiceTest) TestDeleteSuccess() {
	want := true

	c := &rating.ClientMock{}
	c.On("Delete", &proto.DeleteRatingRequest{Id: t.Rating.Id}).Return(&proto.DeleteRatingResponse{Success: true}, nil)

	srv := NewService(c)
	actual, err := srv.Delete(t.Rating.Id)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *RatingServiceTest) TestDeleteNotFound() {
	want := t.NotFoundErr

	c := &rating.ClientMock{}
	c.On("Delete", &proto.DeleteRatingRequest{Id: t.Rating.Id}).Return(nil, status.Error(codes.NotFound, "Rating not found"))

	srv := NewService(c)
	actual, err := srv.Delete(t.Rating.Id)

	assert.False(t.T(), actual)
	assert.Equal(t.T(), want, err)
}

func (t *RatingServiceTest) TestDeleteGrpcErr() {
	want := t.ServiceDownErr

	c := &rating.ClientMock{}
	c.On("Delete", &proto.DeleteRatingRequest{Id: t.Rating.Id}).Return(nil, errors.New("Service is down"))

	srv := NewService(c)
	actual, err := srv.Delete(t.Rating.Id)

	assert.False(t.T(), actual)
	assert.Equal(t.T(), want, err)
}
