package repository

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"type:varchar(255);not null"`
	Password string `gorm:"type:varchar(255);not null"`
	UserName string `gorm:"type:varchar(255);not null"`
}
