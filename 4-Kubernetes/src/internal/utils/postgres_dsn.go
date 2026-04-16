package utils

import (
	"fmt"

	"github.com/merteldem1r/DevOps-Fundamentals/3-Docker-Containerization/src/internal/config"
)

func BuildPostgresDSN(cfg *config.Config) string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=disable",
		cfg.PG.User,
		cfg.PG.Password,
		cfg.PG.Host,
		cfg.PG.Port,
		cfg.PG.DB,
	)
}
