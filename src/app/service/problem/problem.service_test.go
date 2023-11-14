package problem

import (
	"net/http"
	"testing"

	"github.com/bookpanda/mygraderlist-gateway/src/app/dto"
	"github.com/bookpanda/mygraderlist-gateway/src/mocks/problem"
	proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/backend/problem"
	"github.com/bxcodec/faker/v3"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type ProblemServiceTest struct {
	suite.Suite
	Problems       []*proto.Problem
	ServiceDownErr *dto.ResponseErr
}

func TestGroupService(t *testing.T) {
	suite.Run(t, new(ProblemServiceTest))
}

func (t *ProblemServiceTest) SetupTest() {
	Problem1 := &proto.Problem{
		Id:         faker.UUIDDigit(),
		CourseCode: faker.Name(),
		Group:      faker.Name(),
		Code:       faker.Name(),
		Name:       faker.Name(),
	}

	Problem2 := &proto.Problem{
		Id:         faker.UUIDDigit(),
		CourseCode: faker.Name(),
		Group:      faker.Name(),
		Code:       faker.Name(),
		Name:       faker.Name(),
	}

	Problem3 := &proto.Problem{
		Id:         faker.UUIDDigit(),
		CourseCode: faker.Name(),
		Group:      faker.Name(),
		Code:       faker.Name(),
		Name:       faker.Name(),
	}
	t.Problems = make([]*proto.Problem, 0)
	t.Problems = append(t.Problems, Problem1, Problem2, Problem3)

	t.ServiceDownErr = &dto.ResponseErr{
		StatusCode: http.StatusServiceUnavailable,
		Message:    "Service is down",
		Data:       nil,
	}
}

func (t *ProblemServiceTest) TestFindAllSuccess() {
	want := t.Problems

	c := &problem.ClientMock{}
	c.On("FindAll", &proto.FindAllProblemRequest{}).Return(&proto.FindAllProblemResponse{Problems: t.Problems}, nil)

	srv := NewService(c)
	actual, err := srv.FindAll()

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *ProblemServiceTest) TestFindAllGrpcErr() {
	want := t.ServiceDownErr

	c := &problem.ClientMock{}
	c.On("FindAll", &proto.FindAllProblemRequest{}).Return(nil, errors.New("Service is down"))

	srv := NewService(c)
	actual, err := srv.FindAll()

	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), want, err)
}
