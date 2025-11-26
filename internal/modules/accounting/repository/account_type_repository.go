package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/halimdotnet/grango-tesorow/internal/modules/accounting/domain"
	"github.com/halimdotnet/grango-tesorow/internal/pkg/logger"
	pgx "github.com/halimdotnet/grango-tesorow/internal/pkg/postgres"
)

type accountTypeRepository struct {
	db  *pgx.DB
	log logger.Logger
}

func (r *accountTypeRepository) List(ctx context.Context) ([]*domain.AccountType, error) {
	query := `
        SELECT id, code, name, dc_pattern, 
               created_at, created_by, updated_at, updated_by, 
               deleted_at, deleted_by
        FROM accounting.account_types
        WHERE deleted_at IS NULL
        ORDER BY code ASC
    `

	rows, err := r.db.Sqlx.QueryContext(ctx, query)
	if err != nil {
		r.log.Errorf("Failed to query account types: %v", err)
		return nil, err
	}
	defer rows.Close()

	var accountTypes []*domain.AccountType
	for rows.Next() {
		accountType := &domain.AccountType{}
		err = rows.Scan(
			&accountType.ID,
			&accountType.Code,
			&accountType.Name,
			&accountType.DCPattern,
			&accountType.CreatedAt,
			&accountType.CreatedBy,
			&accountType.UpdatedAt,
			&accountType.UpdatedBy,
			&accountType.DeletedAt,
			&accountType.DeletedBy,
		)
		if err != nil {
			r.log.Errorf("Failed to scan account type: %v", err)
			return nil, err
		}
		accountTypes = append(accountTypes, accountType)
	}

	if err = rows.Err(); err != nil {
		r.log.Errorf("Account type row interation error: %v", err)
		return nil, err
	}

	return accountTypes, nil
}

func (r *accountTypeRepository) Find(ctx context.Context, code string) (*domain.AccountType, error) {
	query := `
        SELECT id, code, name, dc_pattern, 
               created_at, created_by, updated_at, updated_by, 
               deleted_at, deleted_by
        FROM accounting.account_types
        WHERE code = $1 AND deleted_at IS NULL
    `

	accountType := &domain.AccountType{}
	err := r.db.Sqlx.QueryRowContext(ctx, query, code).Scan(
		&accountType.ID,
		&accountType.Code,
		&accountType.Name,
		&accountType.DCPattern,
		&accountType.CreatedAt,
		&accountType.CreatedBy,
		&accountType.UpdatedAt,
		&accountType.UpdatedBy,
		&accountType.DeletedAt,
		&accountType.DeletedBy,
	)

	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			r.log.Errorf("Account type with code %q not found", code)
			return nil, err
		}

		r.log.Errorf("Failed to find account type: %v", err)
		return nil, err
	}

	return accountType, nil
}

type AccountTypeRepository interface {
	List(ctx context.Context) ([]*domain.AccountType, error)
	Find(ctx context.Context, code string) (*domain.AccountType, error)
}

func NewAccountTypeRepository(db *pgx.DB, log logger.Logger) AccountTypeRepository {
	return &accountTypeRepository{
		db:  db,
		log: log,
	}
}
