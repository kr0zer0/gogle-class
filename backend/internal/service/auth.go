package service

import (
	"context"
	"errors"
	"gogle-class/backend/config"
	"gogle-class/backend/internal/domain"
	"gogle-class/backend/internal/models"
	"gogle-class/backend/internal/repository"
	"gogle-class/backend/pkg/auth"
	"gogle-class/backend/pkg/hash"
	"strconv"
	"time"
)

type AuthService struct {
	authRepo     repository.Auth
	hasher       hash.PasswordHasher
	tokenManager auth.TokenManager
	cfg          *config.Config
}

func NewAuthService(authRepo repository.Auth, hasher hash.PasswordHasher, tokenManager auth.TokenManager, cfg *config.Config) *AuthService {
	return &AuthService{
		authRepo:     authRepo,
		hasher:       hasher,
		tokenManager: tokenManager,
		cfg:          cfg,
	}
}

func (s *AuthService) Registration(ctx context.Context, user *domain.User) error {
	user.Password = s.hasher.Hash(user.Password)

	err := s.authRepo.Create(ctx, domainUserToDB(user))
	if err != nil {
		return err
	}

	return nil
}

type Tokens struct {
	AccessToken  string
	RefreshToken string
}

func (s *AuthService) Login(ctx context.Context, email string, password string) (tokens Tokens, err error) {
	user, err := s.authRepo.GetUserByEmail(ctx, email)
	if err != nil {
		return
	}

	if user.Password != s.hasher.Hash(password) {
		return tokens, errors.New("emails or password is not correct")
	}

	tokens, err = s.createSession(ctx, user.ID)
	if err != nil {
		return
	}

	return
}

func (s *AuthService) createSession(ctx context.Context, userID uint) (tokens Tokens, err error) {
	tokens.AccessToken, err = s.tokenManager.GenerateToken(strconv.Itoa(int(userID)), s.cfg.Auth.JWT.AccessTokenTTL)
	if err != nil {
		return
	}

	tokens.RefreshToken, err = s.tokenManager.NewRefreshToken()
	if err != nil {
		return
	}

	session := models.Session{
		RefreshToken: tokens.RefreshToken,
		ExpiresAt:    time.Now().Add(s.cfg.Auth.JWT.AccessTokenTTL),
	}

	err = s.authRepo.SetSession(ctx, userID, &session)
	if err != nil {
		return
	}

	return
}

func (s *AuthService) RefreshUserTokens(ctx context.Context, refreshToken string) (tokens Tokens, err error) {
	user, err := s.authRepo.GetUserByRefresh(ctx, refreshToken)
	if err != nil {
		return
	}

	return s.createSession(ctx, user.ID)
}

func domainUserToDB(user *domain.User) *models.User {
	return &models.User{
		Email:    user.Email,
		Password: user.Password,
	}
}
