package app

import (
	"gogle-class/backend/config"
	"gogle-class/backend/internal/controllers/http"
	"gogle-class/backend/internal/infrastructure/repository/postgres"
	"gogle-class/backend/internal/usecase"
	"gogle-class/backend/internal/usecase/usecase_contracts"
	"gogle-class/backend/pkg/auth"
	"gogle-class/backend/pkg/database"
	"gogle-class/backend/pkg/hash"
	"gorm.io/gorm"
)

func Run() error {
	cfg := config.GetConfig()

	hasher := hash.NewSHA256Hasher(cfg.Auth.Salt)
	tokenManager := auth.NewManager(cfg.Auth.JWT.SigningKey)

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		return err
	}

	repositories := initRepositories(db)
	useCases := initUseCases(repositories.AuthRepo, hasher, tokenManager, cfg)
	controllers := http.NewHandler(useCases, tokenManager)

	err = controllers.InitRouter(cfg.App.Port)
	if err != nil {
		return err
	}

	return nil
}

func initRepositories(db *gorm.DB) *usecase_contracts.Repositories {
	authRepo := postgres.NewAuthRepo(db)

	repositories := postgres.NewRepository(authRepo)

	return repositories
}

func initUseCases(authRepo usecase_contracts.AuthRepo, hasher hash.PasswordHasher, tokenManager auth.TokenManager, cfg *config.Config) *usecase.UseCases {
	registerUseCase := usecase.NewRegisterUseCase(authRepo, hasher)
	loginUseCase := usecase.NewLoginUseCase(authRepo, tokenManager, hasher, cfg)
	refreshTokens := usecase.NewRefreshTokensUseCase(authRepo, tokenManager, cfg)

	useCases := usecase.NewUseCases(registerUseCase, loginUseCase, refreshTokens)

	return useCases
}
