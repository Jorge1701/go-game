package render

import (
	"fmt"
	"game/images"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/veandco/go-sdl2/sdl"
)

type Drawable interface {
	GetX() float64
	GetY() float64
	GetImage() *images.Image
}

func Draw(drawable Drawable, screen *ebiten.Image) {
	drawable.GetImage().Draw(screen, drawable.GetX(), drawable.GetY())
}

type Renderer struct {
	renderer *sdl.Renderer
}

func NewRenderer(window *sdl.Window) (*Renderer, error) {
	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)

	if err != nil {
		return &Renderer{}, fmt.Errorf("Error creating new renderer: %v", err)
	}

	return &Renderer{renderer: renderer}, nil
}

func (r *Renderer) Present() {
	r.renderer.Present()
}

func (r *Renderer) Destroy() {
	r.renderer.Destroy()
}
