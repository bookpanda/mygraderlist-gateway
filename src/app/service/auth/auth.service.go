package auth

import (
	"context"
	"net/http"
	"time"

	"github.com/bookpanda/mygraderlist-gateway/src/app/dto"
	auth_proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/auth"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type Service struct {
	client auth_proto.AuthServiceClient
}

func NewService(client auth_proto.AuthServiceClient) *Service {
	return &Service{
		client: client,
	}
}

func (s *Service) Validate(token string) (*dto.TokenPayloadAuth, *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.Validate(ctx, &auth_proto.ValidateRequest{Token: token})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			switch st.Code() {
			case codes.Unauthenticated:
				return nil, &dto.ResponseErr{
					StatusCode: http.StatusUnauthorized,
					Message:    st.Message(),
					Data:       nil,
				}
			default:

				log.Error().
					Err(err).
					Str("service", "auth").
					Str("module", "validate").
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
			Str("service", "auth").
			Str("module", "validate").
			Msg("Error while connecting to service")

		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	return &dto.TokenPayloadAuth{
		UserId: res.UserId,
	}, nil
}

func (s *Service) RefreshToken(token string) (*auth_proto.Credential, *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.RefreshToken(ctx, &auth_proto.RefreshTokenRequest{RefreshToken: token})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			switch st.Code() {
			case codes.Unauthenticated:
				return nil, &dto.ResponseErr{
					StatusCode: http.StatusUnauthorized,
					Message:    st.Message(),
					Data:       nil,
				}
			default:

				log.Error().
					Err(err).
					Str("service", "auth").
					Str("module", "refresh token").
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
			Str("service", "auth").
			Str("module", "refresh token").
			Msg("Error while connecting to service")

		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	return res.Credential, nil
}

func (s *Service) GetGoogleLoginUrl() (string, *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.GetGoogleLoginUrl(ctx, &auth_proto.GetGoogleLoginUrlRequest{})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			switch st.Code() {
			case codes.Unauthenticated:
				return "", &dto.ResponseErr{
					StatusCode: http.StatusUnauthorized,
					Message:    st.Message(),
					Data:       nil,
				}
			default:

				log.Error().
					Err(err).
					Str("service", "auth").
					Str("module", "get google login url").
					Msg("Error while connecting to service")

				return "", &dto.ResponseErr{
					StatusCode: http.StatusServiceUnavailable,
					Message:    "Service is down",
					Data:       nil,
				}
			}
		}

		log.Error().
			Err(err).
			Str("service", "auth").
			Str("module", "get google login url").
			Msg("Error while connecting to service")

		return "", &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	return res.Url, nil
}

func (s *Service) VerifyGoogleLogin(code string) (*auth_proto.Credential, *dto.ResponseErr) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := s.client.VerifyGoogleLogin(ctx, &auth_proto.VerifyGoogleLoginRequest{Code: code})
	if err != nil {
		st, ok := status.FromError(err)
		if ok {
			switch st.Code() {
			case codes.Unauthenticated:
				return nil, &dto.ResponseErr{
					StatusCode: http.StatusUnauthorized,
					Message:    st.Message(),
					Data:       nil,
				}
			default:

				log.Error().
					Err(err).
					Str("service", "auth").
					Str("module", "verify google login").
					Msg("Error while connecting to service, status")

				return nil, &dto.ResponseErr{
					StatusCode: http.StatusServiceUnavailable,
					Message:    "Service is down",
					Data:       nil,
				}
			}
		}

		log.Error().
			Err(err).
			Str("service", "auth").
			Str("module", "verify google login").
			Msg("Error while connecting to service, no status")

		return nil, &dto.ResponseErr{
			StatusCode: http.StatusServiceUnavailable,
			Message:    "Service is down",
			Data:       nil,
		}
	}

	return res.Credential, nil
}
