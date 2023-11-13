package rating

import (
	"context"

	"github.com/bookpanda/mygraderlist-gateway/src/app/dto"
	proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/backend/rating"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type ClientMock struct {
	mock.Mock
}

func (c *ClientMock) FindAll(_ context.Context, in *proto.FindAllRatingRequest, _ ...grpc.CallOption) (res *proto.FindAllRatingResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.FindAllRatingResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) FindByUserId(_ context.Context, in *proto.FindByUserIdRatingRequest, _ ...grpc.CallOption) (res *proto.FindByUserIdRatingResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.FindByUserIdRatingResponse)
	}

	return res, args.Error(1)
}

func (s *ClientMock) Create(_ context.Context, in *proto.CreateRatingRequest, _ ...grpc.CallOption) (res *proto.CreateRatingResponse, err error) {
	args := s.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.CreateRatingResponse)
	}

	return res, args.Error(1)
}

func (s *ClientMock) Update(_ context.Context, in *proto.UpdateRatingRequest, _ ...grpc.CallOption) (res *proto.UpdateRatingResponse, err error) {
	args := s.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.UpdateRatingResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) Delete(_ context.Context, in *proto.DeleteRatingRequest, _ ...grpc.CallOption) (res *proto.DeleteRatingResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.DeleteRatingResponse)
	}

	return res, args.Error(1)
}

type ServiceMock struct {
	mock.Mock
}

func (s *ServiceMock) FindAll() (res []*proto.Rating, err *dto.ResponseErr) {
	args := s.Called()

	if args.Get(0) != nil {
		res = args.Get(0).([]*proto.Rating)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return res, err
}

func (s *ServiceMock) FindByUserId(userId string) (res []*proto.Rating, err *dto.ResponseErr) {
	args := s.Called(userId)

	if args.Get(0) != nil {
		res = args.Get(0).([]*proto.Rating)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return res, err
}

func (s *ServiceMock) Create(in *dto.UserDto) (res *proto.Rating, err *dto.ResponseErr) {
	args := s.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.Rating)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return res, err
}

func (s *ServiceMock) Update(id string, in *dto.UpdateRatingDto) (res *proto.Rating, err *dto.ResponseErr) {
	args := s.Called(id, in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.Rating)
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
