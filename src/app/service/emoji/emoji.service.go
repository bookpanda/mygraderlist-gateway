package emoji

import (
	"context"
	"net/http"
	"time"

	"github.com/bookpanda/mygraderlist-gateway/src/app/dto"
	proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/backend/emoji"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	client proto.EmojiServiceClient
}

func NewService(client proto.EmojiServiceClient) *Service {
	return &Service{
		client: client,
	}
}

func (s *Service) FindAll() ([]*proto.Emoji, *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.FindAll(ctx, &proto.FindAllEmojiRequest{})
	if err != nil {
		log.Error().Err(err).
			Str("service", "emoji").
			Str("module", "find all").
			Msg("Error while find all emoji")

		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
		}
	}

	return res.Emojis, nil
}

func (s *Service) FindByUserId(userId string) ([]*proto.Emoji, *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.FindByUserId(ctx, &proto.FindByUserIdEmojiRequest{UserId: userId})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			switch st.Code() {
			case codes.NotFound:
				return nil, &dto.ResponseErr{
					StatusCode: http.StatusNotFound,
					Message:    "Emoji not found",
					Data:       nil,
				}
			default:
				log.Error().
					Err(err).
					Str("service", "emoji").
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
			Str("service", "emoji").
			Str("module", "find by user id").
			Msg("Error while connecting to service")

		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	return res.Emojis, nil
}

func (s *Service) Create(in *dto.EmojiDto) (*proto.Emoji, *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	EmojiDto := &proto.Emoji{
		Emoji:     in.Emoji,
		ProblemId: in.ProblemID.String(),
		UserId:    in.UserID.String(),
	}

	res, err := s.client.Create(ctx, &proto.CreateEmojiRequest{Emoji: EmojiDto})
	if err != nil {
		log.Error().
			Err(err).
			Str("service", "emoji").
			Str("module", "create").
			Msg("Error while connecting to service")

		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	return res.Emoji, nil
}

func (s *Service) Delete(id string) (bool, *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.Delete(ctx, &proto.DeleteEmojiRequest{Id: id})
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
					Str("service", "emoji").
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
			Str("service", "emoji").
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
