package accounting

import (
	"context"

	"github.com/halimdotnet/grango-tesorow/internal/modules/accounting/dto"
	"github.com/halimdotnet/grango-tesorow/internal/modules/accounting/repository"
	"github.com/halimdotnet/grango-tesorow/internal/modules/accounting/service"
	"github.com/halimdotnet/grango-tesorow/internal/pkg/logger"
	pgx "github.com/halimdotnet/grango-tesorow/internal/pkg/postgres"
)

type AccountClassification interface {
	ListAccountType(ctx context.Context) ([]*dto.AccountTypeResponse, error)
}

func NewAccountClassification(db *pgx.DB, log logger.Logger) AccountClassification {
	accountTypeRepo := repository.NewAccountTypeRepository(db, log)
	categoryRepo := repository.NewAccountCategoryRepository(db, log)

	return service.NewAccountClassificationService(log, accountTypeRepo, categoryRepo)
}
