package postgres

import (
	"context"
	"gogle-class/backend/internal/models"
	"gorm.io/gorm"
)

type AuthRepo struct {
	db *gorm.DB
}

func NewAuthRepo(db *gorm.DB) *AuthRepo {
	return &AuthRepo{db: db}
}

func (r *AuthRepo) Create(ctx context.Context, user *models.User) error {
	return r.db.WithContext(ctx).Create(user).Error
}

func (r *AuthRepo) GetUserByID(ctx context.Context, userID uint) (*models.User, error) {
	var user models.User
	err := r.db.WithContext(ctx).Where("id = ?", userID).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *AuthRepo) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User

	err := r.db.WithContext(ctx).Where("email = ?", email).First(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil

}

func (r *AuthRepo) GetUserByRefresh(ctx context.Context, refreshToken string) (*models.User, error) {
	session, err := r.GetSessionByRefresh(ctx, refreshToken)
	if err != nil {
		return nil, err
	}

	return r.GetUserByID(ctx, session.UserID)
}

func (r *AuthRepo) SetSession(ctx context.Context, userID uint, session *models.Session) error {
	user, err := r.GetUserByID(ctx, userID)
	if err != nil {
		return err
	}

	err = r.db.WithContext(ctx).Model(&user).Association("Session").Replace(session)
	if err != nil {
		return err
	}

	return nil
}

func (r *AuthRepo) GetSessionByRefresh(ctx context.Context, refreshToken string) (*models.Session, error) {
	var session models.Session
	err := r.db.WithContext(ctx).Where("refresh_token = ?", refreshToken).First(&session).Error
	if err != nil {
		return nil, err
	}

	return &session, nil
}
