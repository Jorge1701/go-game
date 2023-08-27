package image

import (
	"fmt"
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

type imageFile struct {
	alias string
	file  string
}

func loadImage(alias, file string) error {
	// Open file image
	f, err := os.Open(file)
	if err != nil {
		return fmt.Errorf("Error opening image file [alias:%s] [file:%s]: %v", alias, file, err)
	}

	// Decode file into image
	image, _, err := image.Decode(f)
	if err != nil {
		return fmt.Errorf("Error decoding image file [alias:%s] [file:%s]: %v", alias, file, err)
	}

	// Save the image to memory
	allImages[alias] = ebiten.NewImageFromImage(image)

	return nil
}
