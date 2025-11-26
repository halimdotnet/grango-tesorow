package app

import (
	"log"

	"github.com/halimdotnet/grango-tesorow/internal/pkg/config"
	"github.com/halimdotnet/grango-tesorow/internal/pkg/hxxp"
	"github.com/halimdotnet/grango-tesorow/internal/pkg/logger"
	"github.com/halimdotnet/grango-tesorow/internal/pkg/postgres"
	"github.com/halimdotnet/grango-tesorow/internal/pkg/validator"
)

type App struct {
	Config    *config.Config
	Logger    logger.Logger
	Server    hxxp.Server
	Router    *hxxp.Router
	Validator *validator.Validator
	Pg        *postgres.DB
}

func NewApp() *App {
	cfg, err := config.BindAllConfig()
	if err != nil {
		log.Fatal(err)
		return nil
	}

	application := &App{
		Config: cfg,
	}

	return application
}

func (a *App) Run() {
	a.providers()
	a.modules()

	defer a.cleanup()

	if err := a.Server.RunServer(); err != nil {
		log.Fatal(err)
	}
}

func (a *App) cleanup() {
	if a.Pg != nil && a.Pg.Sqlx != nil {
		if err := a.Pg.Sqlx.Close(); err != nil {
			log.Printf("Failed to close PostgreSQL connection: %v", err)
		}
	}

	if a.Logger != nil {
		defer a.Logger.Sync()
	}
}
