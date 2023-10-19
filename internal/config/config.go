package config

import (
	"github.com/rs/zerolog/log"

	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DB_HOST                string `required:"true"`
	DB_PORT                string `required:"true"`
	DB_USER                string `required:"true"`
	DB_PASS                string `required:"true"`
	DB_NAME                string `required:"true"`
	DEBUG                  bool   `default:"false"`
	REDIS_HOST             string `required:"true"`
	REDIS_PORT             string `required:"true"`
	REDIS_PASS             string
	REDIS_DB               int `default:"0"`
	WORKER_CONCURRENCY     int `default:"10"`
	WORKER_QUEUE_SCANNER   int `default:"10"`
	WORKER_QUEUE_THUMBNAIL int `default:"1"`
}

func InitializeConfig() (*Config, error) {
	var c Config
	err := envconfig.Process("", &c)
	if err != nil {
		return nil, err
	}

	// ensure queue worker count does not exceed concurrency
	if c.WORKER_QUEUE_SCANNER > c.WORKER_CONCURRENCY {
		log.Fatal().Msg("queue worker count cannot exceed concurrency")
	}

	return &c, nil
}
