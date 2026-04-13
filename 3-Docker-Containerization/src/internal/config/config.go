package config

import (
	"fmt"

	"github.com/ilyakaznacheev/cleanenv"
)

type Config struct {
	MESSAGE string `env:"EXAMPLE_MSG" env-required:"true"`
	PORT    string `env:"PORT" env-default:"8080"`
	PG      PostgresConfig
}

type PostgresConfig struct {
	POSTGRES_USER     string `env:"POSTGRES_USER" env-required:"true"`
	POSTGRES_PASSWORD string `env:"POSTGRES_PASSWORD" env-required:"true"`
	POSTGRES_DB       string `env:"POSTGRES_DB" env-required:"true"`
	POSTGRES_HOST     string `env:"POSTGRES_HOST" env-required:"true"`
	POSTGRES_PORT     string `env:"POSTGRES_PORT" env-default:"5432"`
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
