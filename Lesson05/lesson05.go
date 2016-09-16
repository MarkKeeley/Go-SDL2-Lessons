package main

import (
	"fmt"
	"os"
	"time"

	"github.com/veandco/go-sdl2/sdl"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

// Globals - terrible for serious programs, great for short examples
var (
	window    *sdl.Window
	renderer  *sdl.Renderer
	event     sdl.Event
	quit      bool
	err       error
	ticker    *time.Ticker
	box1      *sdl.Rect
	box2      *sdl.Rect
	box3      sdl.Rect
	velX      int32
	intersect bool
)

// Setup program
func Setup() (successful bool) {
	err = sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		fmt.Printf("Failed to initialize sdl: %s\n", err)
		return false
	}
	window, err = sdl.CreateWindow("Go + SDL2 Lesson 5", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Printf("Failed to create renderer: %s\n", err)
		return false
	}

	renderer, err = sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Printf("Failed to create renderer: %s\n", err)
		return false
	}

	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")

	box1 = &sdl.Rect{X: (screenWidth / 2) - 50,
		Y: (screenHeight / 2) - 50,
		W: 100, H: 100}

	box2 = &sdl.Rect{X: 0,
		Y: (screenHeight / 2),
		W: 100, H: 100}

	velX = 5

	return true
}

// Shutdown program
func Shutdown() {
	renderer.Destroy()
	window.Destroy()
	sdl.Quit()
}

// HandleEvents input mostly
func HandleEvents() {
	for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch t := event.(type) {
		case *sdl.QuitEvent:
			quit = true
		case *sdl.KeyDownEvent:
			if t.Keysym.Sym == sdl.K_ESCAPE {
				quit = true
			}
		}
	}
}

// Draw to screen
func Draw() {
	//clear the screen
	renderer.SetDrawColor(255, 255, 55, 255)
	renderer.Clear()
	//draw box 1 (blue)
	renderer.SetDrawColor(0, 0, 255, 255)
	renderer.FillRect(box1)
	//draw box 2 (green)
	renderer.SetDrawColor(0, 255, 0, 255)
	renderer.FillRect(box2)
	//if there is a collision between box 1 & 2 then draw a third box (red)
	//where the two boxes intersect
	if intersect {
		renderer.SetDrawColor(255, 0, 0, 255)
		renderer.DrawRect(&box3)
	}
	renderer.Present()
}

func main() {
	if !Setup() {
		os.Exit(1)
	}

	ticker = time.NewTicker(time.Second / 30)

	for !quit {
		HandleEvents()
		Update()
		Draw()
		<-ticker.C // wait up to 1/30th of a second

	}

	Shutdown()
}

// Update program
func Update() {
	box2.X += velX

	if box2.X+100 >= screenWidth {
		velX *= -1
	}
	if box2.X <= 0 {
		velX *= -1
	}

	/* HasIntersection is similar to Intersect except that HasIntersection
	   only returns true/false and does not return a sdl.Rect like Intersect
	   does. This is included, but commented out, for completeness.

	if box1.HasIntersection(box2) {
		fmt.Println("intersection")
	}
	*/

	/*check if box1 intersects with box2 - this function returns a bool
	and a sdl.Rect of the intersection if true
	*/
	box3, intersect = box1.Intersect(box2)
	if intersect {
		fmt.Print("Rectange Collision Detected ")
	}

}
