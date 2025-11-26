package service

import (
	"context"
	"fmt"

	"github.com/halimdotnet/grango-tesorow/internal/modules/accounting/dto"
	"github.com/halimdotnet/grango-tesorow/internal/modules/accounting/repository"
)

type accountClassificationService struct {
	accountType repository.AccountTypeRepository
}

type AccountClassificationService interface {
	ListAccountType(ctx context.Context) ([]*dto.AccountTypeResponse, error)
}

func NewAccountClassificationService(accountType repository.AccountTypeRepository) AccountClassificationService {
	return &accountClassificationService{
		accountType: accountType,
	}
}

func (a *accountClassificationService) ListAccountType(ctx context.Context) ([]*dto.AccountTypeResponse, error) {
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
