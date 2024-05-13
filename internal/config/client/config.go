package client

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config represents the configuration of the application
type Config struct {
	LogLevel   int    `env:"LOG_LEVEL" default:"0"`
	ServerAddr string `env:"SERVER_ADDR" default:"localhost:8080"`
}

// FromEnv loads the configuration from the environment
func FromEnv() (Config, error) {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return Config{}, fmt.Errorf("cannot parse envs: %w", err)
	}

	return cfg, nil
}
