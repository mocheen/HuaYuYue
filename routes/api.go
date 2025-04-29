package routes

import (
	"HuaYuYue/controllers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func SetupRoutes(router *gin.Engine, db *gorm.DB) {
	healthCtrl := &controllers.HealthController{DB: db}

	api := router.Group("/api")
	{
		api.GET("/ping", healthCtrl.Ping)
	}
}
