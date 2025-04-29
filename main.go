package main

import (
	"HuaYuYue/config"
	"HuaYuYue/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	db, err := config.InitDB()
	if err != nil {
		panic("failed to connect database")
	}

	// 创建Gin实例
	router := gin.Default()

	// 注册路由
	routes.SetupRoutes(router, db)

	// 启动服务
	router.Run(":8080")
}
