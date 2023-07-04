package main

import (
	"embed"

	"wails-entire/backend"
	"wails-entire/backend/api"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

// TODO cc
func main() {
	// Create an instance of the app structure
	app := backend.NewApp()
	// Create application with options
	err := wails.Run(&options.App{
		Title:  "helloworld",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.Startup,
		// Frameless:        true, // 无边框模式
		Bind: []interface{}{
			app,
			&api.HelloApi{},
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
