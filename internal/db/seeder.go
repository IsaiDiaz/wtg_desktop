package db

import (
	"wtg_desktop/internal/domain/auth"
	"wtg_desktop/internal/logger"

	"gorm.io/gorm"
)

func Seed(db *gorm.DB) {
	var count int64
	db.Model(&auth.AuthRole{}).Count(&count)
	if count == 0 {
		roles := []auth.AuthRole{
			{Name: "ADMIN"},
			{Name: "EMPLOYEE"},
			{Name: "AUDITOR"},
		}
		if err := db.Create(&roles).Error; err != nil {
			logger.Error("Error al hacer seed de empleados:", err)
		} else {
			logger.Info("Datos de empleados sembrados correctamente.")
		}
	} else {
		logger.Info("Empleados ya existen, se omite seed.")
	}
}
