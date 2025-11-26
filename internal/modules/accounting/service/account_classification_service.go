package service

import (
	"context"
	"fmt"

	"github.com/halimdotnet/grango-tesorow/internal/modules/accounting/dto"
	"github.com/halimdotnet/grango-tesorow/internal/modules/accounting/repository"
	pgx "github.com/halimdotnet/grango-tesorow/internal/pkg/postgres"
)

type accountClassification struct {
	accountType repository.AccountTypeRepository
}

type AccountClassification interface {
	ListAccountType(ctx context.Context) ([]*dto.AccountTypeResponse, error)
}

func NewAccountClassificationService(db *pgx.DB, accountType repository.AccountTypeRepository) AccountClassification {
	return &accountClassification{
		accountType: accountType,
	}
}

func (a *accountClassification) ListAccountType(ctx context.Context) ([]*dto.AccountTypeResponse, error) {
	accountTypes, err := a.accountType.List(ctx)
	if err != nil {
		return nil, fmt.Errorf("failed to list account types: %w", err)
	}

	var result []*dto.AccountTypeResponse
	for _, accountType := range accountTypes {
		result = append(result, &dto.AccountTypeResponse{
			ID:        accountType.ID,
			Code:      accountType.Code,
			Name:      accountType.Name,
			DCPattern: accountType.DCPattern,
		})
	}
	return result, nil
}
