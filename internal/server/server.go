package server

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"workly-backend/config"
	"workly-backend/internal/container"
	httpHandler "workly-backend/internal/handler/http"
	"workly-backend/pkg/logger"

	"github.com/gin-gonic/gin"
)

type Server struct {
	config     *config.Config
	httpServer *http.Server
	router     *gin.Engine
}

func NewServer(cfg *config.Config, c *container.Container) *Server {
	server := &Server{
		config: cfg,
	}

	server.router = httpHandler.SetupRoutes(c)

	serverAddr := fmt.Sprintf("%s:%s", cfg.Server.Host, cfg.Server.Port)
	server.httpServer = &http.Server{
		Addr:    serverAddr,
		Handler: server.router,
	}

	return server
}

func (s *Server) Start() error {
	serverAddr := s.httpServer.Addr
	logger.Infof("🚀 Server is running on http://%s", serverAddr)
	logger.Infof("💚 Health check available at http://%s/health", serverAddr)

	if err := s.httpServer.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		return fmt.Errorf("server failed to start: %w", err)
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	logger.Info("🛑 Shutting down server...")

	if err := s.httpServer.Shutdown(ctx); err != nil {
		return fmt.Errorf("server shutdown failed: %w", err)
	}

	logger.Info("✅ Server stopped gracefully")
	return nil
}

func (s *Server) Run() error {
	serverErrors := make(chan error, 1)
	go func() {
		serverErrors <- s.Start()
	}()

	shutdown := make(chan os.Signal, 1)
	signal.Notify(shutdown, syscall.SIGINT, syscall.SIGTERM)

	select {
	case err := <-serverErrors:
		return fmt.Errorf("server error: %w", err)

	case sig := <-shutdown:
		logger.Infof("Received signal: %v", sig)

		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		defer cancel()

		return s.Shutdown(ctx)
	}
}
