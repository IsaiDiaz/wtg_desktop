package bootstrap

import (
	"wtg_desktop/internal/api/desktop"
	"wtg_desktop/internal/domain/category"

	"gorm.io/gorm"
)

func InitCategoryHandler(db *gorm.DB) *desktop.CategoryHandler {
	repo := category.NewRepository(db)
	service := category.NewService(repo)
	return desktop.NewCategoryHandler(service)
}
