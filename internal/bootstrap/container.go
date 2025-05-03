package bootstrap

import (
	"gorm.io/gorm"

	"wtg_desktop/internal/api/desktop"
	"wtg_desktop/internal/domain/category"
	"wtg_desktop/internal/domain/employee"
)

type AppContainer struct {
	EmployeeHandler *desktop.EmployeeHandler
	CategoryHandler *desktop.CategoryHandler
}

func InitAppContainer(db *gorm.DB) *AppContainer {
	// --- Employee wiring ---
	employeeRepo := employee.NewRepository(db)
	employeeService := employee.NewService(employeeRepo)
	employeeHandler := desktop.NewEmployeeHandler(employeeService)

	// --- Category wiring ---
	categoryRepo := category.NewRepository(db)
	categoryService := category.NewService(categoryRepo)
	categoryHandler := desktop.NewCategoryHandler(categoryService)

	return &AppContainer{
		EmployeeHandler: employeeHandler,
		CategoryHandler: categoryHandler,
	}
}
