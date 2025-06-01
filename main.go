package main

import (
	"embed"

	"github.com/joho/godotenv"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"

	"wtg_desktop/internal/app"
	"wtg_desktop/internal/bootstrap/desktop"
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

	environment := config.GetEnvironment()
	logger.Init(environment == "development")
	defer logger.Sync()

	logger.Info("Initializing application...")

	db := db.SetupDatabase()

	container := desktop.InitAppContainer(db)
	webServer := app.StartWebApp(db)

	go func() {
		if err := webServer.Run(":" + config.GetServerPort()); err != nil {
			logger.Fatal(err)
		}
	}()

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
		Bind:             append([]interface{}{app}, container.AllHandlers()...),
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
