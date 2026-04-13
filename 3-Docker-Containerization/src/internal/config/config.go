package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	MESSAGE      string `env:"EXAMPLE_MSG" env-required:"true"`
	PORT         string `env:"PORT" env-default:"8080"`
	POSTGRES_DSN string `env:"POSTGRES_DSN" env-required:"true"`
}

func LoadConfig() (*Config, error) {
	var cfg Config

	if err := cleanenv.ReadConfig(".env", &cfg); err != nil {
		// Fallback: read from OS environment variables only
		if err := cleanenv.ReadEnv(&cfg); err != nil {
			return nil, fmt.Errorf("config: %w", err)
		}
	}

	return &cfg, nil
}
