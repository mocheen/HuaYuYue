package repository

import "fmt"

func migration() {
	err := DB.Set("gorm:table_options", "CHARSET=utf8mb4").
		AutoMigrate(
			&AdminAPL{},
			&UserRole{},
		)
	if err != nil {
		fmt.Println("数据库迁移失败")
	}
}
