package course

import (
	"context"
	"net/http"
	"time"

	"github.com/bookpanda/mygraderlist-gateway/src/app/dto"
	proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/backend/course"
	"github.com/rs/zerolog/log"
)

type Service struct {
	client proto.CourseServiceClient
}

func NewService(client proto.CourseServiceClient) *Service {
	return &Service{
		client: client,
	}
}

func (s *Service) FindAll() ([]*proto.Course, *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.FindAll(ctx, &proto.FindAllCourseRequest{})
	if err != nil {
		log.Error().Err(err).
			Str("service", "course").
			Str("module", "find all").
			Msg("Error while find all course")

		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
		}
	}

	return res.Courses, nil
}
