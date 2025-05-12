package repository

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	UserId uint   `gorm:"type:int(11);not null"`
	Role   string `gorm:"type:varchar(255);not null"`
}
