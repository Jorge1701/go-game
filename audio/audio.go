package audio

import (
	"fmt"

	"github.com/veandco/go-sdl2/mix"
	"github.com/veandco/go-sdl2/sdl"
)

var AllAudios = map[string]*Audio{}

type Audio struct {
	chunk *mix.Chunk
}

func (a *Audio) Play() {
	a.chunk.Play(1, 0)
}

func Initialize() error {
	// Initialize audio
	if err := sdl.Init(sdl.INIT_AUDIO); err != nil {
		return err
	}

	// Open audio device
	if err := mix.OpenAudio(44100, mix.DEFAULT_FORMAT, 2, 8); err != nil {
		return err
	}

	// Load audio files
	if err := loadAudio("shot", "resources/shot.wav"); err != nil {
		return err
	}

	return nil
}

func loadAudio(name, file string) error {
	// Load chunk from file
	chunk, err := mix.LoadWAV(file)

	if err != nil {
		return fmt.Errorf("Error loading audio (%s, %s): %v", name, file, err)
	}

	// Save loaded chunk as audio in memory
	AllAudios[name] = &Audio{
		chunk: chunk,
	}

	return nil
}

func Clear() {
	for _, a := range AllAudios {
		a.chunk.Free()
	}

	sdl.Quit()
	mix.CloseAudio()
}
