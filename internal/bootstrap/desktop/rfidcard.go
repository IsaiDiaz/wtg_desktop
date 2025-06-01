package desktop

import (
	"wtg_desktop/internal/api/desktop"
	"wtg_desktop/internal/domain/rfidcard"

	"gorm.io/gorm"
)

func InitRfidCardHandler(db *gorm.DB) *desktop.RfidCardHandler {
	repo := rfidcard.NewRepository(db)
	service := rfidcard.NewService(repo)
	return desktop.NewRfidCardHandler(service)
}
