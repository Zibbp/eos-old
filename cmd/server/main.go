package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/zibbp/eos/internal/channel"
	"github.com/zibbp/eos/internal/comment"
	"github.com/zibbp/eos/internal/config"
	"github.com/zibbp/eos/internal/database"
	"github.com/zibbp/eos/internal/redis"
	transportHttp "github.com/zibbp/eos/internal/transport/http"
	"github.com/zibbp/eos/internal/video"
)

func Run() error {

	// Config
	c, err := config.InitializeConfig()
	if err != nil {
		log.Panic().Err(err).Msg("failed to initialize config")
	}

	// Setup logging
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	if c.DEBUG {
		log.Info().Msg("logging debug enabled")
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	// Database
	err = database.InitializeDatabase(c.DB_HOST, c.DB_PORT, c.DB_USER, c.DB_PASS, c.DB_NAME)
	if err != nil {
		log.Panic().Err(err).Msg("failed to initialize database")
	}

	// redis
	redis.InitializeAsyncq(c.REDIS_HOST, c.REDIS_PORT, c.REDIS_PASS, c.REDIS_DB)

	// Services
	channelService := channel.NewService()
	videoService := video.NewService()
	commentService := comment.NewService()

	httpHandler := transportHttp.NewHandler(videoService, channelService, commentService)

	if err := httpHandler.Serve(); err != nil {
		return err
	}

	return nil
}

func main() {
	if err := Run(); err != nil {
		log.Fatal().Err(err).Msg("failed to run")
	}
}
