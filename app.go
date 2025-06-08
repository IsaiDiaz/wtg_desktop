package main

import (
	"context"
	"wtg_desktop/internal/discovery"
	"wtg_desktop/internal/logger"
)

// App struct
type App struct {
	ctx              context.Context
	discoveryService *discovery.DiscoveryService
}

// NewApp creates a new App application struct
func NewApp(ds *discovery.DiscoveryService) *App {
	return &App{
		discoveryService: ds,
	}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	logger.Info("Wails application started.")
}

func (a *App) shutdown(ctx context.Context) {
	logger.Info("Wails application shutting down.")
	if a.discoveryService != nil {
		a.discoveryService.Stop() // Detener el servicio de descubrimiento
	}
	logger.Info("Wails application shutdown complete.")
	// Aquí puedes agregar cualquier otra lógica de limpieza necesaria al cerrar la aplicación.
}
