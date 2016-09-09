package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/sdl_ttf"
	"os"
	"time"
)

const (
	screenWidth  = 640
	screenHeight = 480
)

// Globals - terrible for serious programs, great for short examples
var (
	window         *sdl.Window
	renderer       *sdl.Renderer
	event          sdl.Event
	quit           bool
	err            error
	solidTexture   *sdl.Texture
	blendedTexture *sdl.Texture
	shadedTexture  *sdl.Texture
)

func Setup() (successful bool) {
	err = sdl.Init(sdl.INIT_VIDEO)
	if err != nil {
		fmt.Printf("Failed to initialize sdl: %s\n", err)
		return false
	}
	// Using the SDL_ttf library so need to initialize it before using it
	if err = ttf.Init(); err != nil {
		fmt.Printf("Failed to initialize TTF: %s\n", err)
		return false
	}

	window, err = sdl.CreateWindow("Go + SDL2 Lesson 4", sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED, screenWidth, screenHeight, sdl.WINDOW_SHOWN)
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

	return true
}

func Shutdown() {
	solidTexture.Destroy()
	shadedTexture.Destroy()
	blendedTexture.Destroy()
	ttf.Quit()
	renderer.Destroy()
	window.Destroy()
	sdl.Quit()
}

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

func Draw() {
	renderer.SetDrawColor(255, 255, 55, 255)
	renderer.Clear()

	// Draw the 3 text textures
	renderer.Copy(solidTexture, nil, &sdl.Rect{(screenWidth / 2) - 89, 60, 190, 53})
	renderer.Copy(shadedTexture, nil, &sdl.Rect{(screenWidth / 2) - 112, 130, 244, 53})
	renderer.Copy(blendedTexture, nil, &sdl.Rect{(screenWidth / 2) - 117, 200, 244, 53})

	renderer.Present()
}

func CreateFonts() (successful bool) {
	var font *ttf.Font

	if font, err = ttf.OpenFont("Roboto-Regular.ttf", 40); err != nil {
		fmt.Printf("Failed to open font: %s\n", err)
		return false
	}

	var solidSurface *sdl.Surface
	if solidSurface, err = font.RenderUTF8_Solid("Solid Text", sdl.Color{255, 0, 0, 255}); err != nil {
		fmt.Printf("Failed to render text: %s\n", err)
		return false
	}

	if solidTexture, err = renderer.CreateTextureFromSurface(solidSurface); err != nil {
		fmt.Printf("Failed to create texture: %s\n", err)
		return false
	}

	solidSurface.Free()

	var shadedSurface *sdl.Surface
	if shadedSurface, err = font.RenderUTF8_Shaded("Shaded Text",
		sdl.Color{0, 255, 0, 255}, sdl.Color{255, 0, 255, 255}); err != nil {
		fmt.Printf("Failed to render text: %s\n", err)
		return false
	}

	if shadedTexture, err = renderer.CreateTextureFromSurface(shadedSurface); err != nil {
		fmt.Printf("Failed to create texture: %s\n", err)
		return false
	}

	shadedSurface.Free()

	var blendedSurface *sdl.Surface
	if blendedSurface, err = font.RenderUTF8_Blended("Blended Text", sdl.Color{0, 0, 255, 255}); err != nil {
		fmt.Printf("Failed to render text: %s\n", err)
		return false
	}

	if blendedTexture, err = renderer.CreateTextureFromSurface(blendedSurface); err != nil {
		fmt.Printf("Failed to create texture: %s\n", err)
		return false
	}

	blendedSurface.Free()
	font.Close()

	return true
}

func main() {
	// initialize sdl and create the window & renderer
	// if there's an error terminate the program
	if !Setup() {
		os.Exit(1)
	}

	// isolate most of the new code into a seperate func for reading
	//terminate program is there is an error
	if !CreateFonts() {
		os.Exit(2)
	}

	// No more calling time.Sleep! Create a ticker that will trigger up to 30 times a second
	// for a more consistent frame rate and animation speed. Obviously more useful in demos that
	// actually have movement. Still not as good as calculating delta time for consistent framerate
	// independent animation - but this is WAY more convenient!
	ticker := time.NewTicker(time.Second / 30)

	// New main loop. Broke events and drawing into seperate functions - it was getting unwieldy
	for !quit {
		HandleEvents()
		Draw()
		<-ticker.C // wait up to 1/30th of a second

	}
	// Self explainatory
	Shutdown()
}
