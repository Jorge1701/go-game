package main

import (
	"fmt"
	"game/audio"
	"game/configuration"
	"game/game"
	"game/render"
	"os"
	"runtime"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func run() int {
	runtime.LockOSThread()

	if err := ttf.Init(); err != nil {
		fmt.Println("ERROR Initializing TTF")
		return 500
	}
	defer ttf.Quit()

	// Initialize audio
	if err := audio.Initialize(); err != nil {
		fmt.Println("ERROR Initializing audio:", err)
		return 600
	}
	defer audio.Clear()

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

	// Load texts
	gameOverTexture, err := render.NewTextTexture(r, "Game Over!", 200, 50)
	if err != nil {
		fmt.Println("ERROR Creating text texture:", err)
		return 501
	}
	defer gameOverTexture.Destroy()

	// Load textures
	render.LoadAllTextures(r)
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

		if g.IsGameOver {
			r.RenderTexture(gameOverTexture, 0, 0)
		} else {
			g.Update()
		}

		g.Render()

		r.Present()
		sdl.Delay(16)
	}

	return 0
}

func main() {
	os.Exit(run())
}
