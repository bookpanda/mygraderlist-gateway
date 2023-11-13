package emoji

import (
	"context"

	"github.com/bookpanda/mygraderlist-gateway/src/app/dto"
	proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/backend/emoji"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type ClientMock struct {
	mock.Mock
}

func (c *ClientMock) FindAll(_ context.Context, in *proto.FindAllEmojiRequest, _ ...grpc.CallOption) (res *proto.FindAllEmojiResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.FindAllEmojiResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) FindByUserId(_ context.Context, in *proto.FindByUserIdEmojiRequest, _ ...grpc.CallOption) (res *proto.FindByUserIdEmojiResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.FindByUserIdEmojiResponse)
	}

	return res, args.Error(1)
}

func (s *ClientMock) Create(_ context.Context, in *proto.CreateEmojiRequest, _ ...grpc.CallOption) (res *proto.CreateEmojiResponse, err error) {
	args := s.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.CreateEmojiResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) Delete(_ context.Context, in *proto.DeleteEmojiRequest, _ ...grpc.CallOption) (res *proto.DeleteEmojiResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.DeleteEmojiResponse)
	}

	return res, args.Error(1)
}

type ServiceMock struct {
	mock.Mock
}

func (s *ServiceMock) FindAll() (res []*proto.Emoji, err *dto.ResponseErr) {
	args := s.Called()

	if args.Get(0) != nil {
		res = args.Get(0).([]*proto.Emoji)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return res, err
}

func (s *ServiceMock) FindByUserId(userId string) (res []*proto.Emoji, err *dto.ResponseErr) {
	args := s.Called(userId)

	if args.Get(0) != nil {
		res = args.Get(0).([]*proto.Emoji)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return res, err
}

func (s *ServiceMock) Create(in *dto.UserDto) (res *proto.Emoji, err *dto.ResponseErr) {
	args := s.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.Emoji)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return res, err
}

func (s *ServiceMock) Delete(id string) (res bool, err *dto.ResponseErr) {
	args := s.Called(id)

	res = args.Bool(0)

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return res, err
}
