package main

import (
	"context"
	"log/slog"
	"net/http"
	"os"
	"time"

	"github.com/lmittmann/tint"

	"github.com/merteldem1r/DevOps-Fundamentals/3-Docker-Containerization/src/internal/config"
	"github.com/merteldem1r/DevOps-Fundamentals/3-Docker-Containerization/src/internal/database"
	"github.com/merteldem1r/DevOps-Fundamentals/3-Docker-Containerization/src/internal/routes"
)

func main() {
	handler := tint.NewHandler(os.Stdout, &tint.Options{
		Level: slog.LevelDebug,
	})

	logger := slog.New(handler)
	slog.SetDefault(logger)

	slog.Info("Application starting...")

	cfg, err := config.LoadConfig()
	if err != nil {
		slog.Error("Failed to load config", "error", err)
		os.Exit(1)
	}

	// Context for initialization
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()

	// connect PG
	pg, err := database.NewPostgres(ctx, cfg.POSTGRES_DSN, logger)

	if err != nil {
		slog.Error("Error while connecting database", "error", err)
		os.Exit(1)
	}

	// Run database migrations
	if err := database.RunMigrations(cfg.POSTGRES_DSN, logger); err != nil {
		slog.Error("Error while running migrations", "error", err)
		os.Exit(1)
	}

	r := routes.NewRouter(cfg, pg, logger)

	addr := ":" + cfg.PORT

	slog.Info("HTTP server is starting",
		"address", addr,
	)

	if err := http.ListenAndServe(addr, r); err != nil {
		slog.Error("Failed to start HTTP server", "error", err)
		os.Exit(1)
	}
}
