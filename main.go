package main

import (
	"fmt"
	"game/configuration"
	"game/game"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	// Create window
	ebiten.SetWindowTitle(configuration.Title)
	ebiten.SetWindowSize(configuration.Width, configuration.Height)

	g, err := game.NewGame()
	if err != nil {
		fmt.Println("Error creating new game:", err)
		os.Exit(1)
	}

	// Run game loop
	if err := ebiten.RunGame(g); err != nil {
		fmt.Println("Error running game loop:", err)
		os.Exit(1)
	}
}
