package main

import (
    "embed"
    "github.com/wailsapp/wails/v2"
    "github.com/wailsapp/wails/v2/pkg/options"
    "github.com/wailsapp/wails/v2/pkg/options/assetserver"
    "ToDoAmi2/internal"
    "context"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
    app := internal.NewApp()

    err := wails.Run(&options.App{
        Title:  "ToDoAmi",
        Width:  1024,
        Height: 768,
        AssetServer: &assetserver.Options{
        Assets: assets,
    },
        BackgroundColour: &options.RGBA{R: 100, G: 0, B: 0, A: 1},
        OnStartup:        func(ctx context.Context) { app.Startup(ctx) },
        Bind: []interface{}{
        app,
    },
 })

    if err != nil {
        println("Error:", err.Error())
    }
}