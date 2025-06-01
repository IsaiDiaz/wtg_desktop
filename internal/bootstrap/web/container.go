package web

import (
	"gorm.io/gorm"

	"wtg_desktop/internal/api/web"
	"wtg_desktop/internal/domain/auth"
)

type AppContainer struct {
	AuthHandler *web.AuthHandler
}

func InitAppContainer(db *gorm.DB) *AppContainer {
	return &AppContainer{
		AuthHandler: InitAuthHandler(db),
	}
}

func InitAuthHandler(db *gorm.DB) *web.AuthHandler {
	repo := auth.NewRepository(db)
	service := auth.NewService(repo)
	return web.NewAuthHandler(service)
}
