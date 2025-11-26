package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/halimdotnet/grango-tesorow/internal/modules/accounting/domain"
	pgx "github.com/halimdotnet/grango-tesorow/internal/pkg/postgres"
)

type accountTypeRepository struct {
	db *pgx.DB
}

type AccountTypeRepository interface {
	List(ctx context.Context) ([]*domain.AccountType, error)
	Find(ctx context.Context, code string) (*domain.AccountType, error)
}

func NewAccountTypeRepository(db *pgx.DB) AccountTypeRepository {
	return &accountTypeRepository{db: db}
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
		return nil, fmt.Errorf("failed to query account type: %w", err)
	}
	defer rows.Close()

	var accountTypes []*domain.AccountType
	for rows.Next() {
		accountType := &domain.AccountType{}
		err := rows.Scan(
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
			return nil, fmt.Errorf("failed to scan account type row: %w", err)
		}
		accountTypes = append(accountTypes, accountType)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
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
			return nil, fmt.Errorf("account type with code %s not found", code)
		}
		return nil, fmt.Errorf("failed to find account type: %w", err)
	}

	return accountType, nil
}
