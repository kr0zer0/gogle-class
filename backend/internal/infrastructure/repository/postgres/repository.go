package postgres

import (
	"gogle-class/backend/internal/usecase/usecase_contracts"
)

func NewRepository(authRepo usecase_contracts.AuthRepo) *usecase_contracts.Repositories {
	return &usecase_contracts.Repositories{
		AuthRepo: authRepo,
	}
}
