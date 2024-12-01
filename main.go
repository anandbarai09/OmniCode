package main

import (
	"context"
	"embed"
	"fmt"
	"log"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

// Embed the frontend build output
//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create the app instance
	app := &App{}

	// Run the Wails application
	err := wails.Run(&options.App{
		Title:  "OmniCode",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets, // Pass the embedded assets directly
		},
		OnStartup:  app.startup,
		OnShutdown: app.shutdown,
		Bind: []interface{}{
			app,
		},
	})
	if err != nil {
		log.Fatal(err)
	}
}

// App struct
type App struct {
	ctx context.Context
}

// startup is called when the application starts
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// shutdown is called when the application shuts down
func (a *App) shutdown(ctx context.Context) {
	// pass
}

// Greet returns a greeting message
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s!", name)
}
