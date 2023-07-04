package main

import (
	"embed"
	"wails-entire/backend"
	"wails-entire/backend/api"

	"github.com/wailsapp/wails/v2/cmd/floss"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:dist
var assets embed.FS

func main() {

	app := backend.NewApp()

	appoptions := &options.App{
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		Bind: []interface{}{
			app,
			&api.HelloApi{},
		},
	}

	floss.Start(appoptions)

}
