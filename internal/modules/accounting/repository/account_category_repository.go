package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/halimdotnet/grango-tesorow/internal/modules/accounting/domain"
	"github.com/halimdotnet/grango-tesorow/internal/pkg/logger"
	pgx "github.com/halimdotnet/grango-tesorow/internal/pkg/postgres"
)

type accountCategoryRepository struct {
	db  *pgx.DB
	log logger.Logger
}

func (r *accountCategoryRepository) List(ctx context.Context) ([]*domain.AccountCategory, error) {
	query := `
        SELECT 
            ac.id, ac.account_type_id, ac.code, ac.name, 
            ac.classification, ac.is_active,
            ac.created_at, ac.created_by, ac.updated_at, ac.updated_by, 
            ac.deleted_at, ac.deleted_by,
            at.code, at.name, at.dc_pattern
        FROM accounting.account_categories ac
        INNER JOIN accounting.account_types at ON ac.account_type_id = at.id
        WHERE ac.deleted_at IS NULL AND at.deleted_at IS NULL
        ORDER BY ac.code ASC
    `

	rows, err := r.db.Sqlx.QueryContext(ctx, query)
	if err != nil {
		r.log.Errorf("Failed to query list account categories: %v", err)
		return nil, err
	}
	defer rows.Close()

	var categories []*domain.AccountCategory
	for rows.Next() {
		cat := &domain.AccountCategory{}
		err = rows.Scan(
			&cat.ID,
			&cat.AccountTypeID,
			&cat.Code,
			&cat.Name,
			&cat.Classification,
			&cat.IsActive,
			&cat.CreatedAt,
			&cat.CreatedBy,
			&cat.UpdatedAt,
			&cat.UpdatedBy,
			&cat.DeletedAt,
			&cat.DeletedBy,
			&cat.AccountTypeCode,
			&cat.AccountTypeName,
			&cat.AccountTypeDCPattern,
		)
		if err != nil {
			r.log.Errorf("Failed to scan account category: %v", err)
			return nil, err
		}
		categories = append(categories, cat)
	}

	if err = rows.Err(); err != nil {
		r.log.Errorf("Account category row iteration error: %v", err)
		return nil, err
	}

	return categories, nil
}

func (r *accountCategoryRepository) Find(ctx context.Context, code string) (*domain.AccountCategory, error) {
	query := `
	   SELECT
	       ac.id, ac.account_type_id, ac.code, ac.name,
	       ac.classification, ac.is_active,
	       ac.created_at, ac.created_by, ac.updated_at, ac.updated_by,
	       ac.deleted_at, ac.deleted_by,
	       at.code, at.name, at.dc_pattern
	   FROM accounting.account_categories ac
	   INNER JOIN accounting.account_types at ON ac.account_type_id = at.id
	   WHERE ac.code = $1 AND ac.deleted_at IS NULL AND at.deleted_at IS NULL
	`

	cat := &domain.AccountCategory{}
	err := r.db.Sqlx.QueryRowContext(ctx, query, code).Scan(
		&cat.ID,
		&cat.AccountTypeID,
		&cat.Code,
		&cat.Name,
		&cat.Classification,
		&cat.IsActive,
		&cat.CreatedAt,
		&cat.CreatedBy,
		&cat.UpdatedAt,
		&cat.UpdatedBy,
		&cat.DeletedAt,
		&cat.DeletedBy,
		&cat.AccountTypeCode,
		&cat.AccountTypeName,
		&cat.AccountTypeDCPattern,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		r.log.Errorf("failed to find account category: %v", err)
		return nil, err
	}

	return cat, nil
}

type AccountCategoryRepository interface {
	List(ctx context.Context) ([]*domain.AccountCategory, error)
	Find(ctx context.Context, code string) (*domain.AccountCategory, error)
}

func NewAccountCategoryRepository(db *pgx.DB, log logger.Logger) AccountCategoryRepository {
	return &accountCategoryRepository{
		db:  db,
		log: log,
	}
}
