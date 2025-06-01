package main

import (
	"embed"

	"github.com/joho/godotenv"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"wtg_desktop/internal/api/server"
	"wtg_desktop/internal/bootstrap"
	"wtg_desktop/internal/config"
	"wtg_desktop/internal/db"
	"wtg_desktop/internal/logger"
)

//go:embed all:frontend/dist
var assets embed.FS

func init() {
	err := godotenv.Load()
	if err != nil {
		logger.Info("Could not load .env file; using default values")
	}
}

func main() {

	environment := config.GetEnv("ENVIRONMENT", "DEVELOPMENT")
	logger.Init(environment == "DEVELOPMENT")
	defer logger.Sync()

	logger.Info("Initializing application...")

	server.InitServer()
	db := db.SetupDatabase()

	container := bootstrap.InitAppContainer(db)
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "wireless time guardian",
		Width:  1200,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
			container.EmployeeHandler,
			container.CategoryHandler,
			container.DeviceHandler,
			container.RFIDCardHandler,
			container.RFIDCardHistoryHandler,
			container.ProjectHandler,
			container.ProjectEmployeeHandler,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
