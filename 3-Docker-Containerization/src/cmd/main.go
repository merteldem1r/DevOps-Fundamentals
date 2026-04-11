package main

import (
	"log/slog"
	"net/http"
	"os"

	"github.com/lmittmann/tint"
	"github.com/merteldem1r/DevOps-Fundamentals/3-Docker-Containerization/src/internal/config"
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

	r := routes.NewRouter(cfg)

	addr := ":" + cfg.PORT

	slog.Info("HTTP server is starting",
		"address", addr,
	)

	if err := http.ListenAndServe(addr, r); err != nil {
		slog.Error("Failed to start HTTP server", "error", err)
		os.Exit(1)
	}
}
