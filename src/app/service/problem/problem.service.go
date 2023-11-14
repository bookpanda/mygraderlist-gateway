package problem

import (
	"context"
	"net/http"
	"time"

	"github.com/bookpanda/mygraderlist-gateway/src/app/dto"
	proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/backend/problem"
	"github.com/rs/zerolog/log"
)

type Service struct {
	client proto.ProblemServiceClient
}

func NewService(client proto.ProblemServiceClient) *Service {
	return &Service{
		client: client,
	}
}

func (s *Service) FindAll() ([]*proto.Problem, *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.FindAll(ctx, &proto.FindAllProblemRequest{})
	if err != nil {
		log.Error().Err(err).
			Str("service", "problem").
			Str("module", "find all").
			Msg("Error while find all problem")

		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
		}
	}

	return res.Problems, nil
}
