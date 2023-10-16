package redis

import (
	"github.com/hibiken/asynq"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
)

var client *AsynqClient

type AsynqClient struct {
	Client *asynq.Client
}

func InitializeAsyncq(redisHost string, redisPort string, redisPass string, redisDB int) {
	asynq_client := asynq.NewClient(asynq.RedisClientOpt{
		Addr:     redisHost + ":" + redisPort,
		Password: redisPass,
		DB:       redisDB,
	})
	// defer asynq_client.Close()
	log.Debug().Msg("asynq client initialized")
	client = &AsynqClient{Client: asynq_client}
}

func GetAsynqClient() *AsynqClient {
	return client
}

type AsynqLogger struct {
	*zerolog.Logger
}

func NewAsynqLogger() *AsynqLogger {
	return &AsynqLogger{&log.Logger}
}

func (l *AsynqLogger) Debug(args ...interface{}) {
	l.Logger.Debug().Msgf("%v", args...)
}

func (l *AsynqLogger) Info(args ...interface{}) {
	l.Logger.Info().Msgf("%v", args...)
}

func (l *AsynqLogger) Warn(args ...interface{}) {
	l.Logger.Warn().Msgf("%v", args...)
}

func (l *AsynqLogger) Error(args ...interface{}) {
	l.Logger.Error().Msgf("%v", args...)
}

func (l *AsynqLogger) Fatal(args ...interface{}) {
	l.Logger.Fatal().Msgf("%v", args...)
}
