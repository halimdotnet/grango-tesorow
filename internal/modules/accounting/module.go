package accounting

import (
	"github.com/halimdotnet/grango-tesorow/internal/modules/accounting/handler"
	"github.com/halimdotnet/grango-tesorow/internal/modules/accounting/repository"
	"github.com/halimdotnet/grango-tesorow/internal/modules/accounting/service"
	"github.com/halimdotnet/grango-tesorow/internal/pkg/hxxp"
	"github.com/halimdotnet/grango-tesorow/internal/pkg/logger"
	pgx "github.com/halimdotnet/grango-tesorow/internal/pkg/postgres"
)

type Module struct {
	AccountClassificationHandler *handler.AccountClassificationHandler
}

func NewModule(db *pgx.DB, router *hxxp.Router, log logger.Logger) *Module {
	accountTypeRepo := repository.NewAccountTypeRepository(db, log)
	categoryRepo := repository.NewAccountCategoryRepository(db, log)

	accountClassificationSvc := service.NewAccountClassificationService(log, accountTypeRepo, categoryRepo)

	accountClassificationHandler := handler.NewAccountClassificationHandler(
		router,
		accountClassificationSvc,
	)

	return &Module{
		AccountClassificationHandler: accountClassificationHandler,
	}
}

func (m *Module) Register() {
	m.AccountClassificationHandler.RegisterRoutes()
}
