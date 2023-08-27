package image

import (
	"game/graphics"

	"github.com/hajimehoshi/ebiten/v2"
)

var allImages = map[string]*ebiten.Image{}

type ImageManager struct{}

func NewImageManager() (*ImageManager, error) {
	// Load all configured image files
	for _, imageFile := range allImageFiles {
		if err := loadImage(imageFile.alias, imageFile.file); err != nil {
			return nil, err
		}
	}

	return &ImageManager{}, nil
}

func (im *ImageManager) DrawImage(screen *ebiten.Image, alias string, x, y float64) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(x, y)
	screen.DrawImage(allImages[alias], options)
}

func (im *ImageManager) Draw(screen *ebiten.Image, drawable *graphics.Drawable) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(drawable.Rect.Position.X, drawable.Rect.Position.Y)
	screen.DrawImage(allImages[drawable.ImageAlias], options)
}
