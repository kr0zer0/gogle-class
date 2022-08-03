package service

import "gogle-class/backend/internal/repository"

// all service interfaces here...

type Service struct {
	// include all service interfaces
}

func NewService(repository *repository.Repository) *Service {
	return &Service{}
}
