package fonts

import (
	"fmt"
	"os"

	"golang.org/x/image/font"
	"golang.org/x/image/font/opentype"
)

var Font font.Face

func LoadFonts() {
	file, err := os.ReadFile("resources/font.ttf")
	if err != nil {
		fmt.Println("Error loading font")
		os.Exit(1)
	}
	tt, err := opentype.Parse(file)
	if err != nil {
		fmt.Println("Error loading font")
		os.Exit(1)
	}

	font, err := opentype.NewFace(tt, &opentype.FaceOptions{
		Size:    24,
		DPI:     72,
		Hinting: font.HintingVertical,
	})
	if err != nil {
		fmt.Println("Error loading font")
		os.Exit(1)
	}
	Font = font
}
