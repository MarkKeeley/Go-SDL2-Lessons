package main

import (
	"fmt"
	"os"

	"github.com/veandco/go-sdl2/sdl"
)

type Application struct {
	renderer *sdl.Renderer
	window   *sdl.Window
	event    sdl.Event
	err      error
	running  bool
}

func (a *Application) Setup(title string, xposition int32, yposition int32,
	width int32, height int32, fullscreen bool) {
	// initialize SDL
	sdl.Init(sdl.INIT_EVERYTHING)

	var flags uint32 = sdl.WINDOW_SHOWN
	if fullscreen {
		flags = sdl.WINDOW_FULLSCREEN_DESKTOP
	}

	a.window, a.err = sdl.CreateWindow(title, xposition, yposition,
		width, height, flags)
	if a.err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create window: %s\n", a.err)
		os.Exit(1)
	}
	a.renderer, a.err = sdl.CreateRenderer(a.window, -1, 0)
	if a.err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create renderer: %s\n", a.err)
		os.Exit(2)
	}

	//if the code got to here, everything seems to be working
	a.running = true

}

func (a *Application) HandleEvents() {
	for a.event = sdl.PollEvent(); a.event != nil; a.event = sdl.PollEvent() {
		switch t := a.event.(type) {
		case *sdl.QuitEvent:
			a.running = false
		case *sdl.KeyboardEvent:
			if t.Keysym.Sym == sdl.K_ESCAPE {
				a.running = false
			}
		}
	}
}

func (a *Application) Update() {

}

func (a *Application) Render() {
	a.renderer.SetDrawColor(255, 255, 255, 255)
	a.renderer.Clear()

	a.renderer.Present()
}
