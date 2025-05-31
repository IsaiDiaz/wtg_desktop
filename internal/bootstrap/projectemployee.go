package bootstrap

import (
	"wtg_desktop/internal/api/desktop"
	"wtg_desktop/internal/domain/projectemployee"

	"gorm.io/gorm"
)

func InitProjectEmployeeHandler(db *gorm.DB) *desktop.ProjectEmployeeHandler {
	repo := projectemployee.NewRepository(db)
	service := projectemployee.NewService(repo)
	return desktop.NewProjectEmployeeHandler(service)
}
