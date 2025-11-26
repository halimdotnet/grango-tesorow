package accounting

import (
	"context"

	"github.com/halimdotnet/grango-tesorow/internal/modules/accounting/dto"
	"github.com/halimdotnet/grango-tesorow/internal/modules/accounting/repository"
	"github.com/halimdotnet/grango-tesorow/internal/modules/accounting/service"
	pgx "github.com/halimdotnet/grango-tesorow/internal/pkg/postgres"
)

type AccountClassification interface {
	ListAccountType(ctx context.Context) ([]*dto.AccountTypeResponse, error)
}

func NewAccountClassification(db *pgx.DB) AccountClassification {
	accountTypeRepo := repository.NewAccountTypeRepository(db)

	return service.NewAccountClassificationService(accountTypeRepo)
}
