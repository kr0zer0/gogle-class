package usecase

import (
	"context"
	"errors"
	"gogle-class/backend/config"
	"gogle-class/backend/internal/domain"
	"gogle-class/backend/internal/models"
	"gogle-class/backend/internal/usecase/usecase_contracts"
	"gogle-class/backend/pkg/auth"
	"gogle-class/backend/pkg/hash"
	"strconv"
	"time"
)

type LoginUseCase struct {
	authRepo     usecase_contracts.AuthRepo
	tokenManager auth.TokenManager
	hasher       hash.PasswordHasher
	cfg          *config.Config
}

func NewLoginUseCase(authRepo usecase_contracts.AuthRepo, tokenManager auth.TokenManager, hasher hash.PasswordHasher, cfg *config.Config) *LoginUseCase {
	return &LoginUseCase{authRepo: authRepo, tokenManager: tokenManager, hasher: hasher, cfg: cfg}
}

func (u *LoginUseCase) Login(ctx context.Context, email string, password string) (tokens domain.Tokens, err error) {
	user, err := u.authRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return
	}

	if user.Password != u.hasher.Hash(password) {
		return tokens, errors.New("emails or password is not correct")
	}

	tokens, err = createSession(ctx, user.ID, u.tokenManager, u.authRepo, u.cfg)
	if err != nil {
		return
	}

	return
}

func createSession(ctx context.Context, userID uint, tokenManager auth.TokenManager, authRepo usecase_contracts.AuthRepo, cfg *config.Config) (tokens domain.Tokens, err error) {
	tokens.AccessToken, err = tokenManager.GenerateToken(strconv.Itoa(int(userID)), cfg.Auth.JWT.AccessTokenTTL)
	if err != nil {
		return
	}

	tokens.RefreshToken, err = tokenManager.NewRefreshToken()
	if err != nil {
		return
	}

	session := models.Session{
		RefreshToken: tokens.RefreshToken,
		ExpiresAt:    time.Now().Add(cfg.Auth.JWT.RefreshTokenTTL),
	}

	err = authRepo.SetSession(ctx, userID, &session)
	if err != nil {
		return
	}

	return
}
