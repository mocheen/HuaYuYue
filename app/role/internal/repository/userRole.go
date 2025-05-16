package repository

import "gorm.io/gorm"

type UserRole struct {
	gorm.Model
	UserId int32 `gorm:"type:int(11);not null"`
	RoleID int32 `gorm:"type:varchar(255);not null"`
}
