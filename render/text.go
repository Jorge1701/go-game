package render

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
	"github.com/veandco/go-sdl2/ttf"
)

func NewTextTexture(render *Renderer, text string, width, height float64) (*Texture, error) {
	// Loading font
	font, err := ttf.OpenFont("resources/font.ttf", 32)
	if err != nil {
		return &Texture{}, fmt.Errorf("Error loading font: %v", err)
	}
	defer font.Close()

	// Creating textSurface with the font
	textSurface, err := font.RenderUTF8Blended(text, sdl.Color{R: 255, G: 0, B: 0, A: 255})
	if err != nil {
		return &Texture{}, fmt.Errorf("Error creating text surface: %v", err)
	}
	defer textSurface.Free()

	// Creating texture from text surface
	texture, err := render.renderer.CreateTextureFromSurface(textSurface)
	if err != nil {
		return &Texture{}, fmt.Errorf("Error creating texture from surface: %v", err)
	}

	return &Texture{
		T:      texture,
		Width:  width,
		Height: height,
	}, nil
}
