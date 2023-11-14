package problem

import (
	"context"

	"github.com/bookpanda/mygraderlist-gateway/src/app/dto"
	proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/backend/problem"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type ClientMock struct {
	mock.Mock
}

func (c *ClientMock) FindAll(_ context.Context, in *proto.FindAllProblemRequest, _ ...grpc.CallOption) (res *proto.FindAllProblemResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.FindAllProblemResponse)
	}

	return res, args.Error(1)
}

func (s *ClientMock) Create(_ context.Context, in *proto.CreateProblemRequest, _ ...grpc.CallOption) (res *proto.CreateProblemResponse, err error) {
	args := s.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.CreateProblemResponse)
	}

	return res, args.Error(1)
}

func (s *ClientMock) Update(_ context.Context, in *proto.UpdateProblemRequest, _ ...grpc.CallOption) (res *proto.UpdateProblemResponse, err error) {
	args := s.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.UpdateProblemResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) Delete(_ context.Context, in *proto.DeleteProblemRequest, _ ...grpc.CallOption) (res *proto.DeleteProblemResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.DeleteProblemResponse)
	}

	return res, args.Error(1)
}

type ServiceMock struct {
	mock.Mock
}

func (s *ServiceMock) FindAll() (res []*proto.Problem, err *dto.ResponseErr) {
	args := s.Called()

	if args.Get(0) != nil {
		res = args.Get(0).([]*proto.Problem)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return res, err
}
