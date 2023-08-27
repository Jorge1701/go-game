package fonts

import (
	"fmt"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

type fontFile struct {
	alias string
	file  string
	size  float64
}

func loadFonts(alias, file string, size float64) error {
	// Load font file
	fontFile, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("Error loading font file [alias:%s] [file:%s]: %v", alias, file, err)
	}

	// Parse font file
	parsedFont, err := opentype.Parse(fontFile)
	if err != nil {
		return fmt.Errorf("Error parsing font file [alias:%s] [file:%s]: %v", alias, file, err)
	}

	// Get font face
	fontFace, err := opentype.NewFace(parsedFont, &opentype.FaceOptions{
		Size:    size,
		DPI:     72,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		return fmt.Errorf("Error creating new face from font [alias:%s] [file:%s]: %v", alias, file, err)
	}

	// Save font face to be used by the FontManager
	allFontFaces[alias] = &fontFace

	return nil
}
