package desktop

import (
	"wtg_desktop/internal/api/desktop"
	"wtg_desktop/internal/domain/employee"

	"gorm.io/gorm"
)

func InitEmployeeHanlder(db *gorm.DB) *desktop.EmployeeHandler {
	repo := employee.NewRepository(db)
	service := employee.NewService(repo)
	return desktop.NewEmployeeHandler(service)
}
