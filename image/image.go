package image

import (
	"game/graphics"

	"github.com/hajimehoshi/ebiten/v2"
)

var allImages = map[string]*ebiten.Image{}

type ImageManager struct {
	XOffSet float64
	YOffSet float64
}

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
	options.GeoM.Translate(x-im.XOffSet, y-im.YOffSet)
	screen.DrawImage(allImages[alias], options)
}

func (im *ImageManager) Draw(screen *ebiten.Image, drawable *graphics.Drawable) {
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Translate(drawable.Rect.CenterX()-im.XOffSet, drawable.Rect.CenterY()-im.YOffSet)
	screen.DrawImage(allImages[drawable.ImageAlias], options)
}
