package database

import (
	"fmt"
	"workly-backend/config"
	"workly-backend/pkg/logger"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
)

func NewPostgresDB(cfg config.DatabaseConfig) (*sqlx.DB, error) {
	dsn := fmt.Sprintf(
		"host=%s port=%d user=%s password=%s dbname=%s sslmode=%s",
		cfg.Host, cfg.Port, cfg.User, cfg.Password, cfg.DBName, cfg.SSLMode,
	)

	db, err := sqlx.Connect("postgres", dsn)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to database: %w", err)
	}

	db.SetMaxOpenConns(cfg.MaxOpenConns)
	db.SetMaxIdleConns(cfg.MaxIdleConns)
	db.SetConnMaxLifetime(cfg.ConnMaxLifetime)

	if err := db.Ping(); err != nil {
		logger.Error("Failed to ping database", "error", err)
		return nil, err
	}

	logger.Info("Successfully connected to PostgreSQL database with sqlx")
	return db, nil
}

func CloseDB(db *sqlx.DB) error {
	if db != nil {
		return db.Close()
	}
	return nil
}
