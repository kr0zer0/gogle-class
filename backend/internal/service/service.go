package service

import (
	"context"
	"gogle-class/backend/config"
	"gogle-class/backend/internal/domain"
	"gogle-class/backend/internal/repository"
	"gogle-class/backend/pkg/auth"
	"gogle-class/backend/pkg/hash"
)

// all service interfaces here...

type Auth interface {
	Registration(ctx context.Context, user *domain.User) error
	Login(ctx context.Context, email string, password string) (tokens Tokens, err error)
	RefreshUserTokens(ctx context.Context, refreshToken string) (tokens Tokens, err error)
}

type Service struct {
	Auth
}

func NewService(repository *repository.Repository, hasher hash.PasswordHasher, tokenManager auth.TokenManager, cfg *config.Config) *Service {
	return &Service{
		Auth: NewAuthService(repository.Auth, hasher, tokenManager, cfg),
	}
}
