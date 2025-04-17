package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"wtg_desktop/internal/api/desktop"
	"wtg_desktop/internal/domain/employee"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	db, _ := gorm.Open(sqlite.Open("./internal/db/wtg.db"), &gorm.Config{})

	db.AutoMigrate(&employee.Employee{})

	repo := employee.NewRepository(db)
	service := employee.NewService(repo)
	handler := desktop.NewEmployeeHandler(service)
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "wireless time guardian",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			handler,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
