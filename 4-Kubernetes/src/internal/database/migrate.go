package database

import (
	"fmt"
	"log/slog"
	"strings"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/pgx/v5"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/merteldem1r/DevOps-Fundamentals/3-Docker-Containerization/src/internal/config"
	"github.com/merteldem1r/DevOps-Fundamentals/3-Docker-Containerization/src/internal/utils"
	"github.com/merteldem1r/DevOps-Fundamentals/3-Docker-Containerization/src/migrations"
)

func RunMigrations(cfg *config.Config, logger *slog.Logger) error {
	source, err := iofs.New(migrations.FS, ".")
	if err != nil {
		return fmt.Errorf("migrations: create source: %w", err)
	}

	// golang-migrate pgx/v5 driver requires "pgx5://" scheme
	dsn := strings.Replace(utils.BuildPostgresDSN(cfg), "postgres://", "pgx5://", 1)

	m, err := migrate.NewWithSourceInstance("iofs", source, dsn)
	if err != nil {
		return fmt.Errorf("migrations: init: %w", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return fmt.Errorf("migrations: apply: %w", err)
	}

	version, dirty, _ := m.Version()
	logger.Info("Migrations applied", "version", version, "dirty", dirty)

	return nil
}
