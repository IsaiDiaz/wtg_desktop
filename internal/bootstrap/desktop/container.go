package desktop

import (
	"reflect"

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

func (c *AppContainer) AllHandlers() []interface{} {
	var handlers []interface{}
	val := reflect.ValueOf(c).Elem() // Desreferencia el puntero *AppContainer

	for i := 0; i < val.NumField(); i++ {
		field := val.Field(i)
		if field.Kind() == reflect.Ptr && !field.IsNil() {
			handlers = append(handlers, field.Interface())
		}
	}
	return handlers
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
