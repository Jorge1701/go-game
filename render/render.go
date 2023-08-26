package render

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

type Drawable interface {
	GetX() float64
	GetY() float64
	GetTexture() *Texture
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

func (r *Renderer) RenderTexture(texture *Texture, x, y int32) {
	r.renderer.Copy(
		texture.T,
		&sdl.Rect{X: 0, Y: 0, W: int32(texture.Width), H: int32(texture.Height)},
		&sdl.Rect{X: x, Y: y, W: int32(texture.Width), H: int32(texture.Height)},
	)
}

func (r *Renderer) RenderDrawable(drawable Drawable) {
	r.renderer.Copy(
		drawable.GetTexture().T,
		&sdl.Rect{X: 0, Y: 0, W: int32(drawable.GetTexture().Width), H: int32(drawable.GetTexture().Height)},
		&sdl.Rect{X: int32(drawable.GetX()), Y: int32(drawable.GetY()), W: int32(drawable.GetTexture().Width), H: int32(drawable.GetTexture().Height)},
	)
}

func (r *Renderer) Present() {
	r.renderer.Present()
}

func (r *Renderer) Destroy() {
	r.renderer.Destroy()
}
