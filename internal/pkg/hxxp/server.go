package hxxp

import (
	"context"
	"errors"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/go-chi/chi/v5"
	"github.com/halimdotnet/grango-tesorow/internal/pkg/constants"
	"github.com/halimdotnet/grango-tesorow/internal/pkg/logger"
)

type server struct {
	httpServer *http.Server
	config     *Config
	logger     logger.Logger
	router     *Router
}

func NewServer(cfg *Config, log logger.Logger) Server {
	return &server{
		config: cfg,
		logger: log,
	}
}

func (s *server) RunServer() error {
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)

	serverErr := make(chan error, 1)

	go func() {
		if err := s.start(); err != nil {
			serverErr <- err
		}
	}()

	select {
	case err := <-serverErr:
		return err
	case sig := <-quit:

		fmt.Println("Received shutdown signal: ", sig)

		ctx, cancel := context.WithTimeout(context.Background(), constants.ShutdownTimeout)
		defer cancel()

		if err := s.Shutdown(ctx); err != nil {
			return fmt.Errorf("server forced to shutdown: %w", err)
		}
		log.Println("Server gracefully stopped")

		return nil
	}
}

func (s *server) start() error {
	s.httpServer = &http.Server{
		Addr:           fmt.Sprintf("%s:%d", s.config.Host, s.config.Port),
		Handler:        s.router,
		ReadTimeout:    constants.ReadTimeout,
		WriteTimeout:   constants.WriteTimeout,
		IdleTimeout:    constants.IdleTimeout,
		MaxHeaderBytes: constants.MaxHeaderBytes,
	}

	log.Printf("Starting server at %s", s.httpServer.Addr)

	err := s.httpServer.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (s *server) Shutdown(ctx context.Context) error {
	log.Println("Shutting down server...")
	return s.httpServer.Shutdown(ctx)
}

func (s *server) BuildRouter() *Router {
	s.router = &Router{
		chi: chi.NewRouter(),
	}

	return s.router
}
