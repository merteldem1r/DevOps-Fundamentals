package database

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"

	"github.com/merteldem1r/DevOps-Fundamentals/4-Kubernetes/src/internal/config"
	"github.com/merteldem1r/DevOps-Fundamentals/4-Kubernetes/src/internal/utils"
)

func NewPostgres(ctx context.Context, cfg *config.Config) (*pgxpool.Pool, error) {
	dsn := utils.BuildPostgresDSN(cfg)

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("postgres: connect: %w", err)
	}

	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return nil, fmt.Errorf("postgres: ping: %w", err)
	}

	return pool, nil
}

func ConnectPostgresWithRetry(
	ctx context.Context,
	cfg *config.Config,
	logger *slog.Logger,
	maxAttempts int,
	retryDelay time.Duration,
) (*pgxpool.Pool, error) {
	var lastErr error

	for attempt := 1; attempt <= maxAttempts; attempt++ {
		pool, err := NewPostgres(ctx, cfg)
		if err == nil {
			logger.Info("connected to PostgreSQL", "attempt", attempt)
			return pool, nil
		}

		lastErr = err

		logger.Warn("failed to connect to PostgreSQL, retrying",
			"attempt", attempt,
			"max_attempts", maxAttempts,
			"retry_delay", retryDelay,
			"error", err,
		)

		if attempt == maxAttempts {
			break
		}

		select {
		case <-ctx.Done():
			return nil, fmt.Errorf("postgres: retry canceled: %w", ctx.Err())
		case <-time.After(retryDelay):
		}
	}

	return nil, fmt.Errorf("postgres: all retry attempts failed: %w", lastErr)
}
