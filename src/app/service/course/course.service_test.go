package course

import (
	"net/http"
	"testing"

	"github.com/bookpanda/mygraderlist-gateway/src/app/dto"
	"github.com/bookpanda/mygraderlist-gateway/src/mocks/course"
	proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/backend/course"
	"github.com/bxcodec/faker/v3"

	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

type CourseServiceTest struct {
	suite.Suite
	Courses        []*proto.Course
	ServiceDownErr *dto.ResponseErr
}

func TestGroupService(t *testing.T) {
	suite.Run(t, new(CourseServiceTest))
}

func (t *CourseServiceTest) SetupTest() {
	Course1 := &proto.Course{
		Id:         faker.UUIDDigit(),
		CourseCode: faker.Name(),
		Name:       faker.Name(),
		Color:      faker.Name(),
	}

	Course2 := &proto.Course{
		Id:         faker.UUIDDigit(),
		CourseCode: faker.Name(),
		Name:       faker.Name(),
		Color:      faker.Name(),
	}

	Course3 := &proto.Course{
		Id:         faker.UUIDDigit(),
		CourseCode: faker.Name(),
		Name:       faker.Name(),
		Color:      faker.Name(),
	}
	t.Courses = make([]*proto.Course, 0)
	t.Courses = append(t.Courses, Course1, Course2, Course3)

	t.ServiceDownErr = &dto.ResponseErr{
		StatusCode: http.StatusServiceUnavailable,
		Message:    "Service is down",
		Data:       nil,
	}
}

func (t *CourseServiceTest) TestFindAllSuccess() {
	want := t.Courses

	c := &course.ClientMock{}
	c.On("FindAll", &proto.FindAllCourseRequest{}).Return(&proto.FindAllCourseResponse{Courses: t.Courses}, nil)

	srv := NewService(c)
	actual, err := srv.FindAll()

	assert.Nil(t.T(), err)
	assert.Equal(t.T(), want, actual)
}

func (t *CourseServiceTest) TestFindAllGrpcErr() {
	want := t.ServiceDownErr

	c := &course.ClientMock{}
	c.On("FindAll", &proto.FindAllCourseRequest{}).Return(nil, errors.New("Service is down"))

	srv := NewService(c)
	actual, err := srv.FindAll()

	assert.Nil(t.T(), actual)
	assert.Equal(t.T(), want, err)
}
