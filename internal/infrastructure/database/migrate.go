package database

import (
	"fmt"
	"workly-backend/pkg/logger"

	"github.com/jmoiron/sqlx"
	"github.com/pressly/goose/v3"
)

func RunMigrations(db *sqlx.DB) error {
	stdDB := db.DB

	if err := goose.SetDialect("postgres"); err != nil {
		return fmt.Errorf("failed to set goose dialect: %w", err)
	}

	migrationsDir := "migrations"
	if err := goose.Up(stdDB, migrationsDir); err != nil {
		return fmt.Errorf("failed to run database migrations: %w", err)
	}

	logger.Info("✅ Database migrations applied")

	return nil
}
