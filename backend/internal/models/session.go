package models

import (
	"gorm.io/gorm"
	"time"
)

type Session struct {
	gorm.Model
	RefreshToken string `gorm:"type:varchar(200); not null"`
	ExpiresAt    time.Time
	UserID       uint
}
