package service

import (
	"context"

	"github.com/halimdotnet/grango-tesorow/internal/modules/accounting/dto"
	"github.com/halimdotnet/grango-tesorow/internal/modules/accounting/repository"
	"github.com/halimdotnet/grango-tesorow/internal/pkg/logger"
)

type accountClassificationService struct {
	accountType repository.AccountTypeRepository
	log         logger.Logger
}

type AccountClassificationService interface {
	ListAccountType(ctx context.Context) ([]*dto.AccountTypeResponse, error)
}

func NewAccountClassificationService(log logger.Logger, accountType repository.AccountTypeRepository) AccountClassificationService {
	return &accountClassificationService{
		log:         log,
		accountType: accountType,
	}
}

func (s *accountClassificationService) ListAccountType(ctx context.Context) ([]*dto.AccountTypeResponse, error) {
	accountTypes, err := s.accountType.List(ctx)
	if err != nil {
		s.log.Error(err)
		return nil, err
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
