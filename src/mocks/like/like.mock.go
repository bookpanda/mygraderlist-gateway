package like

import (
	"context"

	"github.com/bookpanda/mygraderlist-gateway/src/app/dto"
	proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/backend/like"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type ClientMock struct {
	mock.Mock
}

func (c *ClientMock) FindByUserId(_ context.Context, in *proto.FindByUserIdLikeRequest, _ ...grpc.CallOption) (res *proto.FindByUserIdLikeResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.FindByUserIdLikeResponse)
	}

	return res, args.Error(1)
}

func (s *ClientMock) Create(_ context.Context, in *proto.CreateLikeRequest, _ ...grpc.CallOption) (res *proto.CreateLikeResponse, err error) {
	args := s.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.CreateLikeResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) Delete(_ context.Context, in *proto.DeleteLikeRequest, _ ...grpc.CallOption) (res *proto.DeleteLikeResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.DeleteLikeResponse)
	}

	return res, args.Error(1)
}

type ServiceMock struct {
	mock.Mock
}

func (s *ServiceMock) FindByUserId(userId string) (res []*proto.Like, err *dto.ResponseErr) {
	args := s.Called(userId)

	if args.Get(0) != nil {
		res = args.Get(0).([]*proto.Like)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return res, err
}

func (s *ServiceMock) Create(in *dto.UserDto) (res *proto.Like, err *dto.ResponseErr) {
	args := s.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.Like)
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
