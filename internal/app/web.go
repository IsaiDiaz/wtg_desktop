package app

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"wtg_desktop/internal/bootstrap/web"
)

func StartWebApp(db *gorm.DB) *gin.Engine {
	router := gin.Default()
	api := router.Group("/api")

	container := web.InitAppContainer(db)

	container.AuthHandler.RegisterRoutes(api)

	return router
}
