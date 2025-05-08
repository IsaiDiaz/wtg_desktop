package main

import (
	"context"
	"wtg_desktop/internal/domain/category"
	"wtg_desktop/internal/domain/device"
	"wtg_desktop/internal/domain/employee"
	"wtg_desktop/internal/domain/project"
	"wtg_desktop/internal/domain/rfidcard"
	"wtg_desktop/internal/domain/rfidcardhistory"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

func SetupDatabase() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./internal/db/wtg.db"), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(
		&employee.Employee{},
		&category.Category{},
		&device.Device{},
		&rfidcard.RfidCard{},
		&rfidcardhistory.RfidCardHistory{},
		&project.Project{},
	)

	return db
}
