package bootstrap

import (
	"gorm.io/gorm"

	"wtg_desktop/internal/api/desktop"
)

type AppContainer struct {
	EmployeeHandler        *desktop.EmployeeHandler
	CategoryHandler        *desktop.CategoryHandler
	DeviceHandler          *desktop.DeviceHandler
	RFIDCardHandler        *desktop.RfidCardHandler
	RFIDCardHistoryHandler *desktop.RfidCardHistoryHandler
	ProjectHandler         *desktop.ProjectHandler
	ProjectEmployeeHandler *desktop.ProjectEmployeeHandler
}

func InitAppContainer(db *gorm.DB) *AppContainer {

	return &AppContainer{
		EmployeeHandler:        InitEmployeeHanlder(db),
		CategoryHandler:        InitCategoryHandler(db),
		DeviceHandler:          InitDeviceHandler(db),
		RFIDCardHandler:        InitRfidCardHandler(db),
		RFIDCardHistoryHandler: InitRfidCardHistoryHandler(db),
		ProjectHandler:         InitProjectHandler(db),
		ProjectEmployeeHandler: InitProjectEmployeeHandler(db),
	}
}
