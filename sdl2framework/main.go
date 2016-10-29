package main

import (
	"fmt"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

func main() {

	var app Application
	deltaTime := time.Now()
	ticker := time.NewTicker(time.Second / 60)

	app.Setup("SDL2 Framework", sdl.WINDOWPOS_CENTERED, sdl.WINDOWPOS_CENTERED,
		640, 480, false)

	//Main Loop
	for app.running {
		app.HandleEvents()
		app.Update()
		app.Render()

		fmt.Println(time.Since(deltaTime))
		deltaTime = time.Now()
		<-ticker.C // limit framerate
	}

	app.window.Destroy()
	app.renderer.Destroy()
	sdl.Quit()
	fmt.Println("Exit Game")
	//os.Exit(0)
}
