package database

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/merteldem1r/DevOps-Fundamentals/3-Docker-Containerization/src/internal/config"
	"github.com/merteldem1r/DevOps-Fundamentals/3-Docker-Containerization/src/internal/utils"
)

func NewPostgres(ctx context.Context, cfg *config.Config, logger *slog.Logger) (*pgxpool.Pool, error) {
	dsn := utils.BuildPostgresDSN(cfg)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("postgres: connect: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("postgres: ping: %w", err)
	}

	logger.Info("connected to PostgreSQL")
	return pool, nil
}
