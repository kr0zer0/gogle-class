package usecase

import (
	"context"
	"gogle-class/backend/internal/domain"
	"gogle-class/backend/internal/models"
	"gogle-class/backend/internal/usecase/usecase_contracts"
	"gogle-class/backend/pkg/hash"
)

type RegisterUseCase struct {
	authRepo usecase_contracts.AuthRepo
	hasher   hash.PasswordHasher
}

func NewRegisterUseCase(authRepo usecase_contracts.AuthRepo, hasher hash.PasswordHasher) *RegisterUseCase {
	return &RegisterUseCase{authRepo: authRepo, hasher: hasher}
}

func (u *RegisterUseCase) Register(ctx context.Context, user *domain.User) error {
	user.Password = u.hasher.Hash(user.Password)

	err := u.authRepo.Create(ctx, domainUserToDB(user))
	if err != nil {
		return err
	}

	return nil
}

func domainUserToDB(user *domain.User) *models.User {
	return &models.User{
		Email:    user.Email,
		Password: user.Password,
	}
}
