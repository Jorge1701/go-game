package fonts

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font"
)

var allFontFaces = map[string]*font.Face{}

type FontManager struct{}

func NewFontManager() (*FontManager, error) {
	// Load all font files
	for _, fontFile := range allFontFiles {
		if err := loadFonts(fontFile.alias, fontFile.file, fontFile.size); err != nil {
			return nil, err
		}
	}

	return &FontManager{}, nil
}

func (fm *FontManager) ShowText(screen *ebiten.Image, textToShow, fontAlias string, x, y int, color color.Color) {
	text.Draw(screen, textToShow, *allFontFaces[fontAlias], x, y, color)
}
