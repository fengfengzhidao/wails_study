package main

import (
	"fmt"
	"github.com/wailsapp/wails/v3/pkg/application"
	"time"
)

type GreetService struct{}

func (g *GreetService) Greet(name string) string {

	app := application.Get()
	window := app.Window.NewWithOptions(application.WebviewWindowOptions{
		Title:            "Window " + name,
		BackgroundColour: application.NewRGB(255, 255, 255),
		HTML:             fmt.Sprintf("hello %s", time.Now().Format(time.DateTime)),
	})

	go func() {
		time.Sleep(2 * time.Second)
		window.Close()
	}()

	return "Hello " + name + "!"
}
