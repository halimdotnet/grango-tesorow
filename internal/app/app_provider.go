package app

import (
	"log"

	"github.com/halimdotnet/grango-tesorow/internal/pkg/hxxp"
	"github.com/halimdotnet/grango-tesorow/internal/pkg/logger"
	"github.com/halimdotnet/grango-tesorow/internal/pkg/postgres"
	"github.com/halimdotnet/grango-tesorow/internal/pkg/validator"
)

func (a *App) providers() {
	a.Logger = logger.New(a.Config.Logger, a.Config.Environment)

	a.Validator = validator.New()

	pg, err := postgres.Connect(a.Config.Postgres)
	if err != nil {
		log.Fatal(err)
	}
	a.Pg = pg

	a.Server = hxxp.NewServer(a.Config.Server, a.Logger)
	a.Router = a.Server.BuildRouter()
}
