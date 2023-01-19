package config

import (
	"fmt"
	"os"

	"github.com/rs/zerolog/log"
	"github.com/spf13/viper"
)

type Config struct {
	Debug bool
}

func NewConfig() error {

	// New Viper config
	configLocation := "/data"
	configName := "config"
	configType := "yaml"
	configPath := fmt.Sprintf("%s/%s.%s", configLocation, configName, configType)

	viper.SetConfigFile(configPath)
	viper.SetConfigType(configType)
	viper.SetConfigName(configName)
	viper.AddConfigPath(configLocation)

	// Defaults
	viper.SetDefault("debug", false)

	// Read config
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		log.Info().Msgf("config file not found at %s, creating a new one", configPath)
		err := viper.SafeWriteConfigAs(configPath)
		if err != nil {
			return fmt.Errorf("error writing config file: %s", err)
		}
	} else {
		log.Info().Msgf("config file found at %s, loading...", configPath)
		err := viper.ReadInConfig()
		log.Debug().Msgf("config file loaded: %s", viper.ConfigFileUsed())
		if err != nil {
			return fmt.Errorf("error reading config file: %s", err)
		}
	}

	return nil

}

func GetConfig() *Config {
	return &Config{
		Debug: viper.GetBool("debug"),
	}
}
