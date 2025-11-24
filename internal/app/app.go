package app

import (
	"log"

	"github.com/halimdotnet/grango-tesorow/internal/pkg/config"
	"github.com/halimdotnet/grango-tesorow/internal/pkg/hxxp"
	"github.com/halimdotnet/grango-tesorow/internal/pkg/logger"
)

type App struct {
	Config *config.Config
	Logger logger.Logger
	Server hxxp.Server
	Router *hxxp.Router
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

	if err := a.Server.RunServer(); err != nil {
		log.Fatal(err)
	}

}

func (a *App) providers() {
	a.Logger = logger.New(a.Config.Logger, a.Config.Environment)

	a.Server = hxxp.NewServer(a.Config.Server, a.Logger)
	a.Router = a.Server.BuildRouter()
}

func (a *App) modules() {
	// Do something
}
