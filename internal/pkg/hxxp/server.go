package hxxp

import (
	"context"
	"errors"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/halimdotnet/grango-tesorow/internal/pkg/logger"
)

type Config struct {
	Host            string        `mapstructure:"host"`
	Port            int           `mapstructure:"port"`
	ReadTimeout     time.Duration `mapstructure:"read_timeout"`
	WriteTimeout    time.Duration `mapstructure:"write_timeout"`
	IdleTimeout     time.Duration `mapstructure:"idle_timeout"`
	ShutdownTimeout time.Duration `mapstructure:"shutdown_timeout"`
}

type Route interface {
}

type Server struct {
	config     *Config
	httpServer *http.Server
	logger     logger.Logger
}

func New(cfg *Config, log logger.Logger, env string) *Server {
	return &Server{
		config: cfg,
		httpServer: &http.Server{
			Addr:         fmt.Sprintf("%s:%d", cfg.Host, cfg.Port),
			Handler:      nil,
			ReadTimeout:  cfg.ReadTimeout * time.Second,
			WriteTimeout: cfg.WriteTimeout * time.Second,
			IdleTimeout:  cfg.IdleTimeout * time.Second,
		},
		logger: log,
	}
}

func (s *Server) Run() error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	serverErr := make(chan error, 1)

	fmt.Println("Attempting to start server on:", s.httpServer.Addr)

	go func() {
		if err := s.Start(); err != nil {
			serverErr <- err
		}
	}()

	select {
	case err := <-serverErr:
		return err
	case sig := <-quit:

		fmt.Println("Received shutdown signal: ", sig)
		fmt.Println("Shutting down server gracefully...")

		ctx, cancel := context.WithTimeout(context.Background(), s.config.ShutdownTimeout)
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {
			return fmt.Errorf("server forced to shutdown: %w", err)
		}

		fmt.Println("Server shutdown completed")

		return nil
	}
}

func (s *Server) Start() error {

	fmt.Println("Server successfully started and listening on:", s.httpServer.Addr)

	err := s.httpServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (s *Server) Shutdown(ctx context.Context) error {
	if s.httpServer == nil {
		return nil
	}

	return s.httpServer.Shutdown(ctx)
}
