package server

import (
	"fmt"
	"time"

	"github.com/ilyakaznacheev/cleanenv"
)

// Config represents the configuration of the application
type Config struct {
	ListenPort      int           `env:"LISTEN_PORT" default:"8080"`
	LogLevel        int           `env:"LOG_LEVEL" default:"0"`
	Difficulty      int           `env:"DIFFICULTY" default:"3"`
	ChallengeExpiry time.Duration `env:"CHALLENGE_EXPIRY" default:"10s"`
}

// FromEnv loads the configuration from the environment
func FromEnv() (Config, error) {
	var cfg Config

	if err := cleanenv.ReadEnv(&cfg); err != nil {
		return Config{}, fmt.Errorf("cannot parse envs: %w", err)
	}

	return cfg, nil
}
