package bootstrap

import (
	"gorm.io/gorm"

	"wtg_desktop/internal/api/desktop"
	"wtg_desktop/internal/domain/category"
	"wtg_desktop/internal/domain/device"
	"wtg_desktop/internal/domain/employee"
)

type AppContainer struct {
	EmployeeHandler *desktop.EmployeeHandler
	CategoryHandler *desktop.CategoryHandler
	DeviceHandler   *desktop.DeviceHandler
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

	// --- Device wiring ---
	deviceRepo := device.NewRepository(db)
	deviceService := device.NewService(deviceRepo)
	deviceHandler := desktop.NewDeviceHandler(deviceService)

	return &AppContainer{
		EmployeeHandler: employeeHandler,
		CategoryHandler: categoryHandler,
		DeviceHandler:   deviceHandler,
	}
}
