package utils

import (
	"fmt"

	"github.com/merteldem1r/DevOps-Fundamentals/3-Docker-Containerization/src/internal/config"
)

func BuildPostgresDSN(cfg *config.Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PG.POSTGRES_USER,
		cfg.PG.POSTGRES_PASSWORD,
		cfg.PG.POSTGRES_HOST,
		cfg.PG.POSTGRES_PORT,
		cfg.PG.POSTGRES_DB,
	)
}
