package main

import (
	"fmt"
	"game/configuration"
	"game/images"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var img *ebiten.Image

type Game struct{}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "Hello, World!")

	images.AllImages["background"].Draw(screen, 0, 0)
	images.AllImages["player"].Draw(screen, 10, 10)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return configuration.Width, configuration.Height
}

func main() {
	// Load all images
	if err := images.LoadImages(); err != nil {
		fmt.Println("Error loading images:", err)
		os.Exit(1)
	}

	// Create window
	ebiten.SetWindowTitle(configuration.Title)
	ebiten.SetWindowSize(configuration.Width, configuration.Height)

	// Run game loop
	if err := ebiten.RunGame(&Game{}); err != nil {
		fmt.Println("Error running game loop:", err)
		os.Exit(1)
	}
}
