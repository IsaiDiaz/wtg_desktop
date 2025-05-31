package bootstrap

import (
	"wtg_desktop/internal/api/desktop"
	"wtg_desktop/internal/domain/rfidcardhistory"

	"gorm.io/gorm"
)

func InitRfidCardHistoryHandler(db *gorm.DB) *desktop.RfidCardHistoryHandler {
	repo := rfidcardhistory.NewRepository(db)
	service := rfidcardhistory.NewService(repo)
	return desktop.NewRfidCardHistoryHandler(service)
}
