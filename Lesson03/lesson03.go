package main

import "fmt"
import "os"

import "time"
import "github.com/veandco/go-sdl2/sdl"

// using non bitmap images so we need sdl_image
import "github.com/veandco/go-sdl2/sdl_image"

func main() {
	screenWidth := 640
	screenHeight := 480

	err := sdl.Init(sdl.INIT_EVERYTHING)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to initialize sdl: %s\n", err)
		os.Exit(1)
	}

	window, err := sdl.CreateWindow("Go + SDL2 Lesson 3", sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		screenWidth, screenHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to create renderer: %s\n", err)
		os.Exit(2)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		fmt.Fprint(os.Stderr, "Failed to create renderer: %s\n", err)
		os.Exit(3)
	}

	renderer.Clear()

	// Unnecessary preloading of jpg and png libraries. Can be commented out and program will automatically load
	// the correct library when you use "img.Load()"
	img.Init(img.INIT_JPG | img.INIT_PNG)

	// SUGGEST to sdl that it use a certain scaling quality for images. Default is "0" a.k.a. nearest pixel sampling
	// try out settings 0, 1, 2 to see the differences with the rotating stick figure. Change the
	// time.Sleep(time.Millisecond * 10) into time.Sleep(time.Millisecond * 100) to slow down the speed of the rotating
	// stick figure and get a good look at how blocky the stick figure is at RENDER_SCALE_QUALITY 0 versus 1 or 2
	sdl.SetHint(sdl.HINT_RENDER_SCALE_QUALITY, "1")

	// Load the glorious programmer art stick figure into memory
	surfaceImg, err := img.Load("stick.png")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to load PNG: %s\n", err)
		os.Exit(4)
	}

	// This is for getting the Width and Height of surfaceImg. Once surfaceImg.Free() is called we lose the
	// ability to get information about the image we loaded into ram
	imageWidth := surfaceImg.W
	imageHeight := surfaceImg.H

	// Take the surfaceImg and use it to create a hardware accelerated textureImg. Or in other words take the image
	// sitting in ram and put it onto the graphics card.
	textureImg, err := renderer.CreateTextureFromSurface(surfaceImg)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to create texture: %s\n", err)
		os.Exit(5)
	}
	// We have the image now as a texture so we no longer have need for surface. Time to let it go
	surfaceImg.Free()

	var event sdl.Event
	isRunning := true
	// used for rotating stick figure
	var angle float64 = 0.0

	for isRunning {
		for event = sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch t := event.(type) {
			case *sdl.QuitEvent:
				isRunning = false
			case *sdl.KeyDownEvent:
				if t.Keysym.Sym == sdl.K_ESCAPE {
					isRunning = false
				}
			}
		}

		renderer.SetDrawColor(255, 255, 55, 255)
		renderer.Clear()

		// Draw the first stick figure using the simpler Copy() function. First parameter is the image we want to draw on
		// screen. Second parameter is the source sdl.Rect of what we want to draw. In this case we instead pass nil, a shortcut
		// telling sdl to draw the entire image. You could use a sdl.Rect to specify drawing only a part of the image - especially
		// useful for animation.
		//
		// The third parameter speficies where on the screen the image will go (X & Y) and how large/small it will be. Alter the
		// 50's to grow or shrink the width and height as desired - or use imageWidth and imageHeight instead to use the normal
		// size of the image.
		renderer.Copy(textureImg, nil, &sdl.Rect{0, 0, 50, 50})

		// Make the second stick figure rotate
		angle += 1.0
		if angle > 360.0 {
			angle = 0
		}

		// A different way of drawing onto the screen with more options. The first 3 parameters are the same. The fourth
		// parameter is angle of degrees - use 0 if you don't want the image angled.
		//
		// The fifth parameter is to specify a point that the image rotates around. We use nil to use the default
		// Width / 2 and Height / 2 (vertical and horizontal center of image)
		//
		// The Last parameter is the RenderFlip setting. Do you want your image looking normal? Use sdl.FLIP_NONE
		// Do you want your image looking the other way? sdl.FLIP_HORIZONTAL
		// Do you want your image upside down? sdl.SDL_FLIP_VERTICAL
		// Do you want your image upside down AND looking the other way? sdl.FLIP_HORIZONTAL | sdl.SDL_FLIP_VERTICAL
		renderer.CopyEx(textureImg, nil, &sdl.Rect{200, 200, imageWidth, imageHeight}, angle, nil, sdl.FLIP_HORIZONTAL)

		renderer.Present()

		time.Sleep(time.Millisecond * 10)

	}

	// free the texture memory
	textureImg.Destroy()
	// we may or may not use img.Init(), but it's good form to properly shut down the sdl_image library
	img.Quit()
	renderer.Destroy()
	window.Destroy()

	sdl.Quit()
}
