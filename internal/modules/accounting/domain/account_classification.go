package domain

import "time"

type AccountType struct {
	ID        int
	Code      string
	Name      string
	DCPattern string
	CreatedAt time.Time
	CreatedBy *string
	UpdatedAt time.Time
	UpdatedBy *string
	DeletedAt *time.Time
	DeletedBy *string
}
