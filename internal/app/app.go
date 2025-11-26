package app

import (
	"log"
	"net/http"

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

	// test route
	a.Router.Get("/", func(ctx *hxxp.Context) {
		ctx.Response(http.StatusOK, hxxp.Response{
			Error:   false,
			Message: "Hello World!",
		})
	})

	defer func() {
		if err := a.Logger.Sync(); err != nil {
			log.Printf("Failed to sync logger on shutdown: %v", err)
		}
	}()

	if err := a.Server.RunServer(); err != nil {
		log.Fatal(err)
	}

}

func (a *App) modules() {
	// Do something
}
