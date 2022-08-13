package app

import (
	"fmt"
	"gogle-class/backend/config"
	"gogle-class/backend/internal/handler"
	"gogle-class/backend/internal/repository"
	"gogle-class/backend/internal/service"
	"gogle-class/backend/pkg/auth"
	"gogle-class/backend/pkg/database"
	"gogle-class/backend/pkg/hash"
)

func Run() error {
	cfg := config.GetConfig()

	fmt.Println(cfg.Auth.JWT.AccessTokenTTL)

	db, err := database.NewPostgresDB(cfg)
	if err != nil {
		return err
	}

	repo := repository.NewRepository(db)

	hasher := hash.NewSHA256Hasher(cfg.Auth.Salt)
	tokenManager := auth.NewManager(cfg.Auth.JWT.SigningKey)
	serv := service.NewService(repo, hasher, tokenManager, cfg)
	hand := handler.NewHandler(serv, tokenManager)

	err = hand.InitRouter(cfg.App.Port)
	if err != nil {
		return err
	}

	return nil
}
