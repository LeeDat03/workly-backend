package container

import (
	"fmt"

	"workly-backend/config"
	"workly-backend/internal/domain"
	"workly-backend/internal/infrastructure/database"
	"workly-backend/internal/repository/postgres"
	"workly-backend/internal/usecase"
	"workly-backend/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type Container struct {
	config *config.Config
	db     *sqlx.DB

	UserRepo domain.UserRepository

	UserUseCase *usecase.UserUseCase
	AuthUseCase *usecase.AuthUseCase
}

func NewContainer(cfg *config.Config) (*Container, error) {
	container := &Container{
		config: cfg,
	}

	if err := container.initInfrastructure(); err != nil {
		return nil, err
	}

	container.initRepositories()
	container.initUseCases()

	return container, nil
}

func (c *Container) initInfrastructure() error {
	db, err := database.NewPostgresDB(c.config.Database)
	if err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}
	c.db = db
	logger.Info("✅ Database connection established")

	// TODO: Initialize other infrastructure
	// - Redis cache
	// - Email service
	// - Message queue
	// - S3/Object storage
	// etc.

	return nil
}

func (c *Container) initRepositories() {
	c.UserRepo = postgres.NewUserDbRepository(c.db)

	// TODO: Add more repositories as your application grows
	// c.ProjectRepo = postgres.NewProjectPostgresRepository(c.db)
	// c.TaskRepo = postgres.NewTaskPostgresRepository(c.db)
}

func (c *Container) initUseCases() {
	c.UserUseCase = usecase.NewUserUseCase(c.UserRepo)
	c.AuthUseCase = usecase.NewAuthUseCase(c.UserRepo)

	// TODO: Add more use cases with their dependencies
	// c.ProjectUseCase = usecase.NewProjectUseCase(c.ProjectRepo, c.UserRepo)
	// c.TaskUseCase = usecase.NewTaskUseCase(c.TaskRepo, c.ProjectRepo, c.UserRepo)
}

func (c *Container) Close() error {
	if c.db != nil {
		database.CloseDB(c.db)
		logger.Info("✅ Database connection closed")
	}

	return nil
}
