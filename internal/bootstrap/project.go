package bootstrap

import (
	"wtg_desktop/internal/api/desktop"
	"wtg_desktop/internal/domain/project"

	"gorm.io/gorm"
)

func InitProjectHandler(db *gorm.DB) *desktop.ProjectHandler {
	repo := project.NewRepository(db)
	service := project.NewService(repo)
	return desktop.NewProjectHandler(service)
}
