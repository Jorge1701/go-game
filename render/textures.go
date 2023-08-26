package render

import (
	"fmt"

	"github.com/veandco/go-sdl2/sdl"
)

var AllTextures = map[string]*Texture{}

type Texture struct {
	T      *sdl.Texture
	Width  float64
	Height float64
}

func LoadAllTextures(renderer *Renderer) error {
	playerTexture, err := NewTexture(renderer, "resources/player.bmp", 19, 40)
	if err != nil {
		return fmt.Errorf("Error loading player texture: %v", err)
	}
	AllTextures["player"] = playerTexture

	enemyTexture, err := NewTexture(renderer, "resources/enemy_1.bmp", 19, 17)
	if err != nil {
		return fmt.Errorf("Error loading enemy texture: %v", err)
	}
	AllTextures["enemy"] = enemyTexture

	return nil
}

func NewTexture(render *Renderer, file string, width, height float64) (*Texture, error) {
	img, err := sdl.LoadBMP(file)
	if err != nil {
		return &Texture{}, fmt.Errorf("Error loading image (%s): %v", file, err)
	}
	defer img.Free()

	texture, err := render.renderer.CreateTextureFromSurface(img)
	if err != nil {
		return &Texture{}, fmt.Errorf("Error creating texture from image (%s): %v", file, err)
	}

	return &Texture{
		T:      texture,
		Width:  width,
		Height: height,
	}, nil
}

func ClearAllTextures() {
	for _, texture := range AllTextures {
		texture.Destroy()
	}
}

func (t *Texture) Destroy() {
	t.T.Destroy()
}
