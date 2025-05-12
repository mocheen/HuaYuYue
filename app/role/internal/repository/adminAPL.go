package repository

import "gorm.io/gorm"

// 申请表
type AdminAPL struct {
	gorm.Model
	UserId uint `gorm:"type:int(11);not null"`
	//  0: 待审核 1: 拒绝 2: 同意
	status     uint   `gorm:"type:int(11);not null"`
	APLComment string `gorm:"type:varchar(255);not null"`
	REVComment string `gorm:"type:varchar(255);not null"`
}
