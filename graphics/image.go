package graphics

import (
	"fmt"
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

var AllImages = map[string]*Image{}

type Image struct {
	Image *ebiten.Image
}

func (i *Image) Draw(screen *ebiten.Image, x, y float64) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(x, y)
	screen.DrawImage(i.Image, options)
}

func LoadImages() error {
	if err := loadImage("background", "resources/background.png"); err != nil {
		return err
	}
	if err := loadImage("player", "resources/player.png"); err != nil {
		return err
	}
	if err := loadImage("bullet", "resources/bullet.png"); err != nil {
		return err
	}
	if err := loadImage("enemy", "resources/enemy_1.png"); err != nil {
		return err
	}

	return nil
}

func loadImage(name, file string) error {
	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("Error opening image file [name:%s] [file:%s]: %v", name, file, err)
	}

	i, _, err := image.Decode(f)
	if err != nil {
		return fmt.Errorf("Error decoding image file [name:%s] [file:%s]: %v", name, file, err)
	}

	AllImages[name] = &Image{
		Image: ebiten.NewImageFromImage(i),
	}

	return nil
}
