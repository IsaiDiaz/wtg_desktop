package db

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"wtg_desktop/internal/domain/auth"
	"wtg_desktop/internal/domain/category"
	"wtg_desktop/internal/domain/device"
	"wtg_desktop/internal/domain/employee"
	"wtg_desktop/internal/domain/project"
	"wtg_desktop/internal/domain/projectemployee"
	"wtg_desktop/internal/domain/rfidcard"
	"wtg_desktop/internal/domain/rfidcardhistory"
)

func SetupDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./internal/db/wtg.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(
		&employee.Employee{},
		&category.Category{},
		&device.Device{},
		&rfidcard.RfidCard{},
		&rfidcardhistory.RfidCardHistory{},
		&project.Project{},
		&projectemployee.ProjectEmployee{},
		&auth.AuthUser{},
		&auth.AuthUserRole{},
		&auth.AuthRole{},
		&auth.AuthRolePermission{},
		&auth.AuthPermission{},
	)

	if err != nil {
		panic("migration failed: " + err.Error())
	}

	Seed(db)

	return db
}
