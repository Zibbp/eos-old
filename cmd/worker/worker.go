package main

import (
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"github.com/rs/zerolog/pkgerrors"

	"github.com/hibiken/asynq"
	"github.com/zibbp/eos/internal/config"
	"github.com/zibbp/eos/internal/database"
	"github.com/zibbp/eos/internal/redis"
	"github.com/zibbp/eos/internal/tasks"
	"github.com/zibbp/eos/internal/utils"
)

func main() {

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

	srv := asynq.NewServer(
		asynq.RedisClientOpt{
			Addr:     c.REDIS_HOST + ":" + c.REDIS_PORT,
			Password: c.REDIS_PASS,
			DB:       c.REDIS_DB,
		},
		asynq.Config{
			Concurrency: c.WORKER_CONCURRENCY,
			Queues: map[string]int{
				string(utils.ScannerQueue):            c.WORKER_QUEUE_SCANNER,
				string(utils.ThumbnailGeneratorQueue): c.WORKER_QUEUE_THUMBNAIL,
			},
			// Custom logger for structured logging
			Logger: redis.NewAsynqLogger(),
		},
	)

	log.Info().Msgf("starting worker: concurrency=%d %s=%d %s=%d", c.WORKER_CONCURRENCY, utils.ScannerQueue, c.WORKER_QUEUE_SCANNER, utils.ThumbnailGeneratorQueue, c.WORKER_QUEUE_THUMBNAIL)

	err = database.InitializeDatabase(c.DB_HOST, c.DB_PORT, c.DB_USER, c.DB_PASS, c.DB_NAME)
	if err != nil {
		log.Panic().Err(err).Msg("failed to initialize database")
	}
	redis.InitializeAsyncq(c.REDIS_HOST, c.REDIS_PORT, c.REDIS_PASS, c.REDIS_DB)

	mux := asynq.NewServeMux()
	mux.HandleFunc("video:start_scanner", tasks.HandleVideoStartScannerTask)
	mux.HandleFunc("video:scan_channel", tasks.HandleVideoScanChannelTask)
	mux.HandleFunc("video:process", tasks.HandleVideoProcessTask)
	mux.HandleFunc("video:generate_thumbnails", tasks.HandleVideoGenerateThumbnailsTask)

	if err := srv.Run(mux); err != nil {
		log.Fatal().Err(err).Msg("error running server")
	}
}
