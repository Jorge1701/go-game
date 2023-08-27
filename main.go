package main

import (
	"fmt"
	"game/configuration"
	"game/game"
	"game/images"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// Load all images
	if err := images.LoadImages(); err != nil {
		fmt.Println("Error loading images:", err)
		os.Exit(1)
	}

	// Create window
	ebiten.SetWindowTitle(configuration.Title)
	ebiten.SetWindowSize(configuration.Width, configuration.Height)

	g := game.NewGame()

	// Run game loop
	if err := ebiten.RunGame(g); err != nil {
		fmt.Println("Error running game loop:", err)
		os.Exit(1)
	}
}
