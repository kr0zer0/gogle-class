package models

import (
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(50); not null; unique"`
	Password string `gorm:"type:varchar(200); not null"`
	Session  Session
}
