package web

import (
	"reflect"

	"github.com/gin-gonic/gin"
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

func (c *AppContainer) RegisterAllRoutes(router *gin.RouterGroup) {
	val := reflect.ValueOf(c).Elem()

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)

		if field.Kind() == reflect.Ptr && !field.IsNil() {
			if handler, ok := field.Interface().(Routable); ok {
				handler.RegisterRoutes(router)
			}
		}
	}
}
