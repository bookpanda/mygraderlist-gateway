package rating

import (
	"context"
	"net/http"
	"time"

	"github.com/bookpanda/mygraderlist-gateway/src/app/dto"
	proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/backend/rating"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	client proto.RatingServiceClient
}

func NewService(client proto.RatingServiceClient) *Service {
	return &Service{
		client: client,
	}
}

func (s *Service) FindAll() ([]*proto.Rating, *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.FindAll(ctx, &proto.FindAllRatingRequest{})
	if err != nil {
		log.Error().Err(err).
			Str("service", "rating").
			Str("module", "find all").
			Msg("Error while find all rating")

		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
		}
	}

	return res.Ratings, nil
}

func (s *Service) FindByUserId(userId string) ([]*proto.Rating, *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.FindByUserId(ctx, &proto.FindByUserIdRatingRequest{UserId: userId})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			switch st.Code() {
			case codes.NotFound:
				return nil, &dto.ResponseErr{
					StatusCode: http.StatusNotFound,
					Message:    "Rating not found",
					Data:       nil,
				}
			default:
				log.Error().
					Err(err).
					Str("service", "rating").
					Str("module", "find by user id").
					Msg("Error while connecting to service")

				return nil, &dto.ResponseErr{
					StatusCode: http.StatusServiceUnavailable,
					Message:    "Service is down",
					Data:       nil,
				}
			}
		}

		log.Error().
			Err(err).
			Str("service", "rating").
			Str("module", "find by user id").
			Msg("Error while connecting to service")

		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	return res.Ratings, nil
}

func (s *Service) Create(in *dto.RatingDto) (*proto.Rating, *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	ratingDto := &proto.Rating{
		ProblemId: in.ProblemID.String(),
		UserId:    in.UserID.String(),
	}

	res, err := s.client.Create(ctx, &proto.CreateRatingRequest{Rating: ratingDto})
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "rating").
			Str("module", "create").
			Msg("Error while connecting to service")

		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	return res.Rating, nil
}

func (s *Service) Update(id string, in *dto.UpdateRatingDto) (*proto.Rating, *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	usrReq := &proto.UpdateRatingRequest{
		Id:         id,
		Score:      int32(in.Score),
		Difficulty: int32(in.Difficulty),
	}

	res, err := s.client.Update(ctx, usrReq)
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			switch st.Code() {
			case codes.NotFound:
				return nil, &dto.ResponseErr{
					StatusCode: http.StatusNotFound,
					Message:    st.Message(),
					Data:       nil,
				}
			default:

				log.Error().
					Err(err).
					Str("service", "rating").
					Str("module", "update").
					Msg("Error while connecting to service")

				return nil, &dto.ResponseErr{
					StatusCode: http.StatusServiceUnavailable,
					Message:    "Service is down",
					Data:       nil,
				}
			}
		}

		log.Error().
			Err(err).
			Str("service", "rating").
			Str("module", "update").
			Msg("Error while connecting to service")

		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	return res.Rating, nil
}

func (s *Service) Delete(id string) (bool, *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.Delete(ctx, &proto.DeleteRatingRequest{Id: id})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			switch st.Code() {
			case codes.NotFound:
				return false, &dto.ResponseErr{
					StatusCode: http.StatusNotFound,
					Message:    st.Message(),
					Data:       nil,
				}
			default:

				log.Error().
					Err(err).
					Str("service", "rating").
					Str("module", "delete").
					Msg("Error while connecting to service")

				return false, &dto.ResponseErr{
					StatusCode: http.StatusServiceUnavailable,
					Message:    "Service is down",
					Data:       nil,
				}
			}
		}

		log.Error().
			Err(err).
			Str("service", "rating").
			Str("module", "delete").
			Msg("Error while connecting to service")

		return false, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	return res.Success, nil
}
