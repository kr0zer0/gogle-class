package repository

import (
	"context"
	"gogle-class/backend/internal/models"
	"gorm.io/gorm"
)

// all repos' interfaces here...

type Auth interface {
	Create(ctx context.Context, user *models.User) error
	GetUserByID(ctx context.Context, userID uint) (*models.User, error)
	GetUserByEmail(ctx context.Context, email string) (*models.User, error)
	SetSession(ctx context.Context, userID uint, session *models.Session) error
	GetUserByRefresh(ctx context.Context, refreshToken string) (*models.User, error)
}

type Repository struct {
	// include all repos' interfaces

	Auth
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		Auth: NewAuthRepo(db),
	}
}
