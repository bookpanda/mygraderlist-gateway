package course

import (
	"context"

	"github.com/bookpanda/mygraderlist-gateway/src/app/dto"
	proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/backend/course"
	"github.com/stretchr/testify/mock"
	"google.golang.org/grpc"
)

type ClientMock struct {
	mock.Mock
}

func (c *ClientMock) FindAll(_ context.Context, in *proto.FindAllCourseRequest, _ ...grpc.CallOption) (res *proto.FindAllCourseResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.FindAllCourseResponse)
	}

	return res, args.Error(1)
}

func (s *ClientMock) Create(_ context.Context, in *proto.CreateCourseRequest, _ ...grpc.CallOption) (res *proto.CreateCourseResponse, err error) {
	args := s.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.CreateCourseResponse)
	}

	return res, args.Error(1)
}

func (s *ClientMock) Update(_ context.Context, in *proto.UpdateCourseRequest, _ ...grpc.CallOption) (res *proto.UpdateCourseResponse, err error) {
	args := s.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.UpdateCourseResponse)
	}

	return res, args.Error(1)
}

func (c *ClientMock) Delete(_ context.Context, in *proto.DeleteCourseRequest, _ ...grpc.CallOption) (res *proto.DeleteCourseResponse, err error) {
	args := c.Called(in)

	if args.Get(0) != nil {
		res = args.Get(0).(*proto.DeleteCourseResponse)
	}

	return res, args.Error(1)
}

type ServiceMock struct {
	mock.Mock
}

func (s *ServiceMock) FindAll() (res []*proto.Course, err *dto.ResponseErr) {
	args := s.Called()

	if args.Get(0) != nil {
		res = args.Get(0).([]*proto.Course)
	}

	if args.Get(1) != nil {
		err = args.Get(1).(*dto.ResponseErr)
	}

	return res, err
}
