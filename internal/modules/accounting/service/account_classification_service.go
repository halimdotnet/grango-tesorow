package service

import (
	"context"

	"github.com/halimdotnet/grango-tesorow/internal/modules/accounting/dto"
	"github.com/halimdotnet/grango-tesorow/internal/modules/accounting/repository"
	"github.com/halimdotnet/grango-tesorow/internal/pkg/logger"
)

type accountClassificationService struct {
	accountType repository.AccountTypeRepository
	category    repository.AccountCategoryRepository
	log         logger.Logger
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

func (s *accountClassificationService) ListCategory(ctx context.Context) ([]*dto.AccountCategoryResponse, error) {
	categories, err := s.category.List(ctx)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}

	var result []*dto.AccountCategoryResponse
	for _, category := range categories {
		result = append(result, &dto.AccountCategoryResponse{
			ID:                   category.ID,
			Code:                 category.Code,
			Name:                 category.Name,
			Classification:       category.Classification,
			IsActive:             category.IsActive,
			AccountTypeCode:      category.AccountTypeCode,
			AccountTypeName:      category.AccountTypeName,
			AccountTypeDCPattern: category.AccountTypeDCPattern,
		})
	}

	return result, nil
}

func (s *accountClassificationService) FindCategory(ctx context.Context, code string) (*dto.AccountCategoryResponse, error) {
	category, err := s.category.Find(ctx, code)
	if err != nil {
		s.log.Error(err)
		return nil, err
	}

	if category == nil {
		return nil, nil
	}

	return &dto.AccountCategoryResponse{
		ID:                   category.ID,
		Code:                 category.Code,
		Name:                 category.Name,
		Classification:       category.Classification,
		IsActive:             category.IsActive,
		AccountTypeCode:      category.AccountTypeCode,
		AccountTypeName:      category.AccountTypeName,
		AccountTypeDCPattern: category.AccountTypeDCPattern,
	}, nil
}

type AccountClassificationService interface {
	ListAccountType(ctx context.Context) ([]*dto.AccountTypeResponse, error)

	ListCategory(ctx context.Context) ([]*dto.AccountCategoryResponse, error)
	FindCategory(ctx context.Context, code string) (*dto.AccountCategoryResponse, error)
}

func NewAccountClassificationService(
	log logger.Logger,
	accountType repository.AccountTypeRepository,
	category repository.AccountCategoryRepository,
) AccountClassificationService {
	return &accountClassificationService{
		log:         log,
		accountType: accountType,
		category:    category,
	}
}
