package container

import (
	"fmt"

	"workly-backend/config"
	"workly-backend/internal/domain"
	v1 "workly-backend/internal/handler/http/v1"
	"workly-backend/internal/infrastructure/database"
	"workly-backend/internal/repository/postgres"
	"workly-backend/internal/usecase"
	"workly-backend/pkg/logger"

	"github.com/jmoiron/sqlx"
)

type Container struct {
	Config *config.Config
	db     *sqlx.DB
	//Layer Repo
	UserRepo domain.UserRepository
	//Layer UseCase
	UserUseCase *usecase.UserUseCase
	AuthUseCase *usecase.AuthUseCase
	//Layer Handler
	AuthHandler *v1.AuthHandler
	UserHandler *v1.UserHandler
}

func NewContainer(cfg *config.Config) (*Container, error) {
	container := &Container{
		Config: cfg,
	}

	if err := container.initInfrastructure(); err != nil {
		return nil, err
	}

	container.initRepositories()
	container.initUseCases()
	container.initHandler()

	return container, nil
}

func (c *Container) initInfrastructure() error {
	db, err := database.NewPostgresDB(c.Config.Database)
	if err != nil {
		return fmt.Errorf("failed to initialize database: %w", err)
	}
	c.db = db
	logger.Info("✅ Database connection established")

	ok := false
	if ok {
		if err := database.RunMigrations(c.db); err != nil {
			return fmt.Errorf("database migrations failed: %w", err)
		}
		logger.Info("✅ Database migrations applied")
	}

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

func (c *Container) initHandler() {
	c.AuthHandler = v1.NewAuthHandler(c.AuthUseCase)
	c.UserHandler = v1.NewUserHandler(c.UserUseCase)
}

func (c *Container) Close() error {
	if c.db != nil {
		database.CloseDB(c.db)
		logger.Info("✅ Database connection closed")
	}

	return nil
}
