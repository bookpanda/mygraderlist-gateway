package emoji

import (
	"net/http"
	"testing"

	"github.com/bookpanda/mygraderlist-gateway/src/app/dto"
	"github.com/bookpanda/mygraderlist-gateway/src/mocks/emoji"
	proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/backend/emoji"
	user_proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/backend/user"
	"github.com/bxcodec/faker/v3"
	"github.com/google/uuid"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type EmojiServiceTest struct {
	suite.Suite
	User           *user_proto.User
	Emoji          *proto.Emoji
	Emojis         []*proto.Emoji
	EmojiDto       *dto.EmojiDto
	NotFoundErr    *dto.ResponseErr
	ServiceDownErr *dto.ResponseErr
}

func TestGroupService(t *testing.T) {
	suite.Run(t, new(EmojiServiceTest))
}

func (t *EmojiServiceTest) SetupTest() {
	t.User = &user_proto.User{
		Id:       faker.UUIDDigit(),
		Username: faker.Username(),
		Email:    faker.Email(),
		Password: faker.Password(),
	}

	t.EmojiDto = &dto.EmojiDto{
		Emoji:     "😀",
		ProblemID: uuid.New(),
		UserID:    uuid.New(),
	}

	t.Emoji = &proto.Emoji{
		Id:        faker.UUIDDigit(),
		Emoji:     t.EmojiDto.Emoji,
		ProblemId: t.EmojiDto.ProblemID.String(),
		UserId:    t.EmojiDto.UserID.String(),
	}

	Emoji2 := &proto.Emoji{
		Id: faker.UUIDDigit(),

		Emoji:     "🛌",
		ProblemId: uuid.New().String(),
		UserId:    uuid.New().String(),
	}

	Emoji3 := &proto.Emoji{
		Id:        faker.UUIDDigit(),
		Emoji:     "🌱",
		ProblemId: uuid.New().String(),
		UserId:    uuid.New().String(),
	}
	t.Emojis = make([]*proto.Emoji, 0)
	t.Emojis = append(t.Emojis, t.Emoji, Emoji2, Emoji3)

	t.ServiceDownErr = &dto.ResponseErr{
		StatusCode: http.StatusServiceUnavailable,
		Message:    "Service is down",
		Data:       nil,
	}

	t.NotFoundErr = &dto.ResponseErr{
		StatusCode: http.StatusNotFound,
		Message:    "Emoji not found",
		Data:       nil,
	}
}

func (t *EmojiServiceTest) TestFindAllSuccess() {
	want := t.Emojis

	c := &emoji.ClientMock{}
	c.On("FindAll", &proto.FindAllEmojiRequest{}).Return(&proto.FindAllEmojiResponse{Emojis: t.Emojis}, nil)

	srv := NewService(c)
	actual, err := srv.FindAll()

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *EmojiServiceTest) TestFindAllGrpcErr() {
	want := t.ServiceDownErr

	c := &emoji.ClientMock{}
	c.On("FindAll", &proto.FindAllEmojiRequest{}).Return(nil, errors.New("Service is down"))

	srv := NewService(c)
	actual, err := srv.FindAll()

	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), want, err)
}

func (t *EmojiServiceTest) TestFindByUserIdSuccess() {
	want := t.Emojis

	c := &emoji.ClientMock{}
	c.On("FindByUserId", &proto.FindByUserIdEmojiRequest{UserId: t.User.Id}).Return(&proto.FindByUserIdEmojiResponse{Emojis: t.Emojis}, nil)

	srv := NewService(c)
	actual, err := srv.FindByUserId(t.User.Id)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *EmojiServiceTest) TestFindByUserIdNotFound() {
	want := t.NotFoundErr

	c := &emoji.ClientMock{}
	c.On("FindByUserId", &proto.FindByUserIdEmojiRequest{UserId: t.User.Id}).Return(nil, status.Error(codes.NotFound, "Emoji not found"))

	srv := NewService(c)
	actual, err := srv.FindByUserId(t.User.Id)

	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), want, err)
}

func (t *EmojiServiceTest) TestFindByUserIdGrpcErr() {
	want := t.ServiceDownErr

	c := &emoji.ClientMock{}
	c.On("FindByUserId", &proto.FindByUserIdEmojiRequest{UserId: t.User.Id}).Return(nil, errors.New("Service is down"))

	srv := NewService(c)
	actual, err := srv.FindByUserId(t.User.Id)

	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), want, err)
}

func (t *EmojiServiceTest) TestCreateSuccess() {
	want := t.Emoji

	in := &proto.Emoji{
		ProblemId: t.EmojiDto.ProblemID.String(),
		UserId:    t.EmojiDto.UserID.String(),
	}

	c := &emoji.ClientMock{}
	c.On("Create", &proto.CreateEmojiRequest{Emoji: in}).Return(&proto.CreateEmojiResponse{Emoji: t.Emoji}, nil)

	srv := NewService(c)
	actual, err := srv.Create(t.EmojiDto)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *EmojiServiceTest) TestCreateGrpcErr() {
	want := t.ServiceDownErr

	in := &proto.Emoji{
		ProblemId: t.EmojiDto.ProblemID.String(),
		UserId:    t.EmojiDto.UserID.String(),
	}

	c := &emoji.ClientMock{}
	c.On("Create", &proto.CreateEmojiRequest{Emoji: in}).Return(nil, errors.New("Service is down"))

	srv := NewService(c)
	actual, err := srv.Create(t.EmojiDto)

	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), want, err)
}

func (t *EmojiServiceTest) TestDeleteSuccess() {
	want := true

	c := &emoji.ClientMock{}
	c.On("Delete", &proto.DeleteEmojiRequest{Id: t.Emoji.Id}).Return(&proto.DeleteEmojiResponse{Success: true}, nil)

	srv := NewService(c)
	actual, err := srv.Delete(t.Emoji.Id)

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *EmojiServiceTest) TestDeleteNotFound() {
	want := t.NotFoundErr

	c := &emoji.ClientMock{}
	c.On("Delete", &proto.DeleteEmojiRequest{Id: t.Emoji.Id}).Return(nil, status.Error(codes.NotFound, "Emoji not found"))

	srv := NewService(c)
	actual, err := srv.Delete(t.Emoji.Id)

	assert.False(t.T(), actual)
	assert.Equal(t.T(), want, err)
}

func (t *EmojiServiceTest) TestDeleteGrpcErr() {
	want := t.ServiceDownErr

	c := &emoji.ClientMock{}
	c.On("Delete", &proto.DeleteEmojiRequest{Id: t.Emoji.Id}).Return(nil, errors.New("Service is down"))

	srv := NewService(c)
	actual, err := srv.Delete(t.Emoji.Id)

	assert.False(t.T(), actual)
	assert.Equal(t.T(), want, err)
}
