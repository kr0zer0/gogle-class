package app

import (
	"gogle-class/backend/internal/handler"
	"gogle-class/backend/internal/repository"
	"gogle-class/backend/internal/service"
)

func Run() error {

	repo := repository.NewRepository()
	servs := service.NewService(repo)
	hand := handler.NewHandler(servs)

	err := hand.InitRouter(":8080")
	if err != nil {
		return err
	}

	return nil
}
