package usecase

import (
	"context"
	"gogle-class/backend/config"
	"gogle-class/backend/internal/domain"
	"gogle-class/backend/internal/usecase/usecase_contracts"
	"gogle-class/backend/pkg/auth"
)

type RefreshTokensUseCase struct {
	authRepo     usecase_contracts.AuthRepo
	tokenManager auth.TokenManager
	cfg          *config.Config
}

func NewRefreshTokensUseCase(authRepo usecase_contracts.AuthRepo, tokenManager auth.TokenManager, cfg *config.Config) *RefreshTokensUseCase {
	return &RefreshTokensUseCase{authRepo: authRepo, tokenManager: tokenManager, cfg: cfg}
}

func (u *RefreshTokensUseCase) RefreshUserTokens(ctx context.Context, refreshToken string) (tokens domain.Tokens, err error) {
	user, err := u.authRepo.GetUserByRefresh(ctx, refreshToken)
	if err != nil {
		return
	}

	return createSession(ctx, user.ID, u.tokenManager, u.authRepo, u.cfg)
}
