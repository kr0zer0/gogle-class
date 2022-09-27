package usecase_contracts

import (
	"context"
	"gogle-class/backend/internal/domain"
	"gogle-class/backend/internal/models"
)

type (
	RegisterUseCase interface {
		Register(ctx context.Context, user *domain.User) error
	}
	LoginUseCase interface {
		Login(ctx context.Context, email string, password string) (tokens domain.Tokens, err error)
	}
	RefreshTokensUseCase interface {
		RefreshUserTokens(ctx context.Context, refreshToken string) (tokens domain.Tokens, err error)
	}
)

type (
	AuthRepo interface {
		Create(ctx context.Context, user *models.User) error
		GetUserByID(ctx context.Context, userID uint) (*models.User, error)
		GetUserByEmail(ctx context.Context, email string) (*models.User, error)
		SetSession(ctx context.Context, userID uint, session *models.Session) error
		GetUserByRefresh(ctx context.Context, refreshToken string) (*models.User, error)
	}

	Repositories struct {
		AuthRepo
	}
)
