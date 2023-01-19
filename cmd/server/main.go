package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"
	"github.com/zibbp/eos/internal/channel"
	"github.com/zibbp/eos/internal/comment"
	"github.com/zibbp/eos/internal/config"
	"github.com/zibbp/eos/internal/database"
	"github.com/zibbp/eos/internal/scanner"
	transportHttp "github.com/zibbp/eos/internal/transport/http"
	"github.com/zibbp/eos/internal/video"
)

func Run() error {

	// Config
	err := config.NewConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("failed to load config")
	}

	// Setup logging
	configDebug := config.GetConfig().Debug
	zerolog.ErrorStackMarshaler = pkgerrors.MarshalStack
	if configDebug {
		log.Info().Msg("debug mode enabled")
		zerolog.SetGlobalLevel(zerolog.DebugLevel)
	} else {
		zerolog.SetGlobalLevel(zerolog.InfoLevel)
	}

	// Database
	err = database.InitializeDatabase()
	if err != nil {
		log.Panic().Err(err).Msg("failed to initialize database")
	}

	// Services
	channelService := channel.NewService()
	videoService := video.NewService()
	scannerService := scanner.NewService(channelService, videoService)
	commentService := comment.NewService()

	httpHandler := transportHttp.NewHandler(videoService, channelService, scannerService, commentService)

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
