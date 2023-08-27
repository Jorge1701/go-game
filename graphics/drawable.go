package graphics

import (
	"github.com/hajimehoshi/ebiten/v2"
)

type Drawable interface {
	GetX() float64
	GetY() float64
	GetImage() *Image
}

func Draw(drawable Drawable, screen *ebiten.Image) {
	drawable.GetImage().Draw(screen, drawable.GetX(), drawable.GetY())
}
