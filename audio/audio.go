package audio

import (
	"fmt"
	"game/configuration"
	"io"

	"github.com/hajimehoshi/ebiten/v2/audio"
)

var allAudioReaders = map[string]io.Reader{}
var allAudioBytes = map[string][]byte{}

type AudioPlayer struct {
	audioContext *audio.Context
}

func NewAudioPlayer() (*AudioPlayer, error) {
	// Creat audio context
	audioContext := audio.NewContext(configuration.SampleRate)

	// Load all configured audio files
	for _, audioFile := range allAudioFiles {
		if err := loadAudio(audioFile.alias, audioFile.file, audioFile.audioType); err != nil {
			return nil, err
		}
	}

	return &AudioPlayer{
		audioContext: audioContext,
	}, nil
}

// PlayFromBytes allows to play the same audio multiple times
// Good for quick effects that will overlap with multiple instances of the same audio
func (ap *AudioPlayer) PlayFromBytes(audioName string) {
	player := ap.audioContext.NewPlayerFromBytes(allAudioBytes[audioName])
	player.Play()
}

// PlayFromReader can play the same audio multiple times
// but if executed before the previous finishes it will stop the it and start from the beginning
// Good for background ambient and music
func (ap *AudioPlayer) PlayFromReader(audioName string) error {
	player, err := ap.audioContext.NewPlayer(allAudioReaders[audioName])
	if err != nil {
		return fmt.Errorf("Error playing from audio reader [name:%s]: %v", audioName, err)
	}

	player.Play()

	return nil
}
