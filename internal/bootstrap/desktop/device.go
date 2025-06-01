package desktop

import (
	"wtg_desktop/internal/api/desktop"
	"wtg_desktop/internal/domain/device"

	"gorm.io/gorm"
)

func InitDeviceHandler(db *gorm.DB) *desktop.DeviceHandler {
	repo := device.NewRepository(db)
	service := device.NewService(repo)
	return desktop.NewDeviceHandler(service)
}
