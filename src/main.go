package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	_log "log"

	athHdr "github.com/bookpanda/mygraderlist-gateway/src/app/handler/auth"
	emjHdr "github.com/bookpanda/mygraderlist-gateway/src/app/handler/emoji"
	health_check "github.com/bookpanda/mygraderlist-gateway/src/app/handler/health-check"
	lkHdr "github.com/bookpanda/mygraderlist-gateway/src/app/handler/like"
	rtngHdr "github.com/bookpanda/mygraderlist-gateway/src/app/handler/rating"
	usrHdr "github.com/bookpanda/mygraderlist-gateway/src/app/handler/user"
	guard "github.com/bookpanda/mygraderlist-gateway/src/app/middleware/auth"
	"github.com/bookpanda/mygraderlist-gateway/src/app/router"
	athSrv "github.com/bookpanda/mygraderlist-gateway/src/app/service/auth"
	emjSrv "github.com/bookpanda/mygraderlist-gateway/src/app/service/emoji"
	lkSrv "github.com/bookpanda/mygraderlist-gateway/src/app/service/like"
	rtngSrv "github.com/bookpanda/mygraderlist-gateway/src/app/service/rating"
	usrSrv "github.com/bookpanda/mygraderlist-gateway/src/app/service/user"
	"github.com/bookpanda/mygraderlist-gateway/src/app/validator"
	"github.com/bookpanda/mygraderlist-gateway/src/config"
	"github.com/bookpanda/mygraderlist-gateway/src/constant/auth"
	auth_proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/auth"
	emoji_proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/backend/emoji"
	like_proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/backend/like"
	rating_proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/backend/rating"
	user_proto "github.com/bookpanda/mygraderlist-proto/MyGraderList/backend/user"
	"github.com/rs/zerolog/log"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conf, err := config.LoadConfig()
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "config").
			Msg("Failed to start service")
	}

	v, err := validator.NewValidator()
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "validator").
			Msg("Failed to start service")
	}

	backendConn, err := grpc.Dial(conf.Service.Backend, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "mgl-backend").
			Msg("Cannot connect to service")
	}

	authConn, err := grpc.Dial(conf.Service.Auth, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal().
			Err(err).
			Str("service", "mgl-auth").
			Msg("Cannot connect to service")
	}

	hc := health_check.NewHandler()

	userClient := user_proto.NewUserServiceClient(backendConn)
	userSrv := usrSrv.NewService(userClient)
	userHdr := usrHdr.NewHandler(userSrv, v)

	authClient := auth_proto.NewAuthServiceClient(authConn)
	authSrv := athSrv.NewService(authClient)
	authHdr := athHdr.NewHandler(authSrv, userSrv, v)

	likeClient := like_proto.NewLikeServiceClient(backendConn)
	likeSrv := lkSrv.NewService(likeClient)
	likeHdr := lkHdr.NewHandler(likeSrv, v)

	emojiClient := emoji_proto.NewEmojiServiceClient(backendConn)
	emojiSrv := emjSrv.NewService(emojiClient)
	emojiHdr := emjHdr.NewHandler(emojiSrv, v)

	ratingClient := rating_proto.NewRatingServiceClient(backendConn)
	ratingSrv := rtngSrv.NewService(ratingClient)
	ratingHdr := rtngHdr.NewHandler(ratingSrv, v)

	authGuard := guard.NewAuthGuard(authSrv, auth.ExcludePath, conf.App)

	r := router.NewGinRouter(&authGuard, conf.App)

	r.GetHealthCheck("/", hc.HealthCheck)

	// r.PostUser("/login", userHdr.FindOne)

	if conf.App.Debug {
		r.GetUser("/:id", userHdr.FindOne)
		r.PostUser("/", userHdr.Create)
		r.DeleteUser("/:id", userHdr.Delete)
	}

	r.GetAuth("/me", authHdr.Validate)
	r.PostAuth("/refreshToken", authHdr.RefreshToken)

	r.GetLike("/mylikes", likeHdr.FindByUserId)
	r.PostLike("/", likeHdr.Create)
	r.DeleteLike("/:id", likeHdr.Create)

	r.GetEmoji("/", emojiHdr.FindAll)
	r.GetEmoji("/myemojis", emojiHdr.FindByUserId)
	r.PostEmoji("/", emojiHdr.Create)
	r.DeleteEmoji("/:id", emojiHdr.Create)

	r.GetRating("/", ratingHdr.FindAll)
	r.GetRating("/myratings", ratingHdr.FindByUserId)
	r.PostRating("/", ratingHdr.Create)
	r.DeleteRating("/:id", ratingHdr.Create)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%v", conf.App.Port),
		Handler: r,
	}

	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			_log.Fatalf("listen: %s\n", err)
		}
	}()

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatal().
				Err(err).
				Str("service", "mgl-gateway").
				Msg("Server not close properly")
		}
	}()

	wait := gracefulShutdown(context.Background(), 2*time.Second, map[string]operation{
		"server": func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	<-wait
}

type operation func(ctx context.Context) error

func gracefulShutdown(ctx context.Context, timeout time.Duration, ops map[string]operation) <-chan struct{} {
	wait := make(chan struct{})
	go func() {
		s := make(chan os.Signal, 1)

		signal.Notify(s, syscall.SIGINT, syscall.SIGTERM, os.Interrupt)
		sig := <-s

		log.Info().
			Str("service", "graceful shutdown").
			Msgf("got signal \"%v\" shutting down service", sig)

		timeoutFunc := time.AfterFunc(timeout, func() {
			log.Error().
				Str("service", "graceful shutdown").
				Msgf("timeout %v ms has been elapsed, force exit", timeout.Milliseconds())
			os.Exit(0)
		})

		defer timeoutFunc.Stop()

		var wg sync.WaitGroup

		for key, op := range ops {
			wg.Add(1)
			innerOp := op
			innerKey := key
			go func() {
				defer wg.Done()

				log.Info().
					Str("service", "graceful shutdown").
					Msgf("cleaning up: %v", innerKey)
				if err := innerOp(ctx); err != nil {
					log.Error().
						Str("service", "graceful shutdown").
						Err(err).
						Msgf("%v: clean up failed: %v", innerKey, err.Error())
					return
				}

				log.Info().
					Str("service", "graceful shutdown").
					Msgf("%v was shutdown gracefully", innerKey)
			}()
		}

		wg.Wait()
		close(wait)
	}()

	return wait
}
