package main

import (
	"fmt"
	"game/configuration"
	"game/game"
	"game/render"
	"os"
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
)

func run() int {
	runtime.LockOSThread()

	// Creating window
	window, err := sdl.CreateWindow(
		configuration.Title,
		sdl.WINDOWPOS_UNDEFINED,
		sdl.WINDOWPOS_UNDEFINED,
		configuration.Width,
		configuration.Height,
		sdl.WINDOW_SHOWN,
	)
	if err != nil {
		fmt.Println("ERROR CreateWindow:", err)
		return 100
	}
	defer window.Destroy()

	// Creating renderer
	r, err := render.NewRenderer(window)
	if err != nil {
		fmt.Println("ERROR CreateRenderer:", err)
		return 200
	}
	defer r.Destroy()

	render.LoadTextures(r)
	defer render.ClearAllTextures()

	// Loading background texture
	background, err := render.NewTexture(r, "resources/background.bmp", 400, 350)
	if err != nil {
		fmt.Println("ERROR Loading background:", err)
		return 300
	}
	defer background.Destroy()

	// Creating game object
	g, err := game.NewGame(r)
	if err != nil {
		fmt.Println("ERROR Creating new game:", err)
		return 400
	}

	// Game loop
	running := true
	for running {
		// Check for quit input
		for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
			switch event.(type) {
			case sdl.QuitEvent:
				running = false
			}
		}

		r.RenderTexture(background, 0, 0)

		g.Update()
		g.Render()

		r.Present()
		sdl.Delay(16)

		if g.IsGameOver {
			running = false
		}
	}

	return 0
}

func main() {
	os.Exit(run())
}
