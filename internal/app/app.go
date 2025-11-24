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
	Server *hxxp.Server
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

	//spew.Dump(application.Config.Logger)

	return application
}

func (a *App) Run() {
	a.providers()
	a.modules()

	if err := a.Server.Run(); err != nil {
		log.Fatal(err)
	}
}

func (a *App) providers() {
	a.Logger = logger.New(a.Config.Logger, a.Config.Environment)
	a.Server = hxxp.New(a.Config.Server, a.Logger, a.Config.Environment)
}

func (a *App) modules() {
	// Do something
}
