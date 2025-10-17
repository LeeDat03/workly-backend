package main

import (
	"workly-backend/config"
	"workly-backend/internal/container"
	"workly-backend/internal/server"
	"workly-backend/pkg/logger"
)

func main() {
	logger.Init("info", "🚀 Workly ")

	cfg, err := config.Load()
	if err != nil {
		logger.Fatal("Failed to load configuration", "error", err)
	}

	di, err := container.NewContainer(cfg)
	if err != nil {
		logger.Fatal("Failed to initialize container", "error", err)
	}
	defer di.Close()

	srv := server.NewServer(cfg, di)
	if err := srv.Run(); err != nil {
		logger.Error("Server error", "error", err)
	}
}
