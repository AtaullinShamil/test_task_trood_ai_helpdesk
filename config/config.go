package config

import (
	"github.com/pkg/errors"
	"github.com/spf13/viper"
)

type Config struct {
	RabbitMQ RabbitMQConfig `mapstructure:"rabbitmq"`
}

type RabbitMQConfig struct {
	URL string `mapstructure:"url"`
}

func Load() (*Config, error) {
	viper.SetConfigFile("./config/config.yml")
	if err := viper.ReadInConfig(); err != nil {
		return nil, errors.Wrap(err, "ReadInConfig")
	}

	cfg := &Config{}

	if err := viper.Unmarshal(cfg); err != nil {
		return nil, errors.Wrap(err, "Unmarshal")
	}

	return cfg, nil
}
