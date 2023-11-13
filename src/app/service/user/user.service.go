package user

import (
	"context"
	"net/http"
	"time"

	"github.com/bookpanda/mygraderlist-gateway/src/app/dto"
	user_proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/backend/user"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	client user_proto.UserServiceClient
}

func NewService(client user_proto.UserServiceClient) *Service {
	return &Service{
		client: client,
	}
}

func (s *Service) FindOne(id string) (result *user_proto.User, err *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, errRes := s.client.FindOne(ctx, &user_proto.FindOneUserRequest{Id: id})
	if errRes != nil {
		st, ok := status.FromError(errRes)
		if ok {
			switch st.Code() {
			case codes.NotFound:
				return nil, &dto.ResponseErr{
					StatusCode: http.StatusNotFound,
					Message:    "User not found",
					Data:       nil,
				}
			default:

				log.Error().
					Err(errRes).
					Str("service", "user").
					Str("module", "findOne").
					Msg("Error while connecting to service")

				return nil, &dto.ResponseErr{
					StatusCode: http.StatusServiceUnavailable,
					Message:    "Service is down",
					Data:       nil,
				}
			}
		}

		log.Error().
			Err(errRes).
			Str("service", "user").
			Str("module", "findOne").
			Msg("Error while connecting to service")

		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	return res.User, nil
}

func (s *Service) Create(in *dto.UserDto) (result *user_proto.User, err *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	usrDto := &user_proto.User{
		Username: in.Username,
		Email:    in.Email,
		Password: in.Password,
	}

	res, errRes := s.client.Create(ctx, &user_proto.CreateUserRequest{User: usrDto})
	if errRes != nil {

		log.Error().
			Err(errRes).
			Str("service", "user").
			Str("module", "create").
			Msg("Error while connecting to service")

		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	return res.User, nil
}

func (s *Service) Update(id string, in *dto.UpdateUserDto) (result *user_proto.User, err *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	usrReq := &user_proto.UpdateUserRequest{
		Username: in.Username,
		Password: in.Password,
	}

	res, errRes := s.client.Update(ctx, usrReq)
	if errRes != nil {
		st, ok := status.FromError(errRes)
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
					Err(errRes).
					Str("service", "user").
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
			Err(errRes).
			Str("service", "user").
			Str("module", "update").
			Msg("Error while connecting to service")

		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	return res.User, nil
}

func (s *Service) Delete(id string) (result bool, err *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, errRes := s.client.Delete(ctx, &user_proto.DeleteUserRequest{Id: id})
	if errRes != nil {
		st, ok := status.FromError(errRes)
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
					Err(errRes).
					Str("service", "user").
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
			Err(errRes).
			Str("service", "user").
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
