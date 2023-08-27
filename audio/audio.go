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
	audioContext := audio.NewContext(configuration.SampleRate)

	if err := loadAllAudios(); err != nil {
		return nil, fmt.Errorf("Error loading audios: %v", err)
	}

	return &AudioPlayer{
		audioContext: audioContext,
	}, nil
}

func (ap *AudioPlayer) PlayFromBytes(audioName string) {
	player := ap.audioContext.NewPlayerFromBytes(allAudioBytes[audioName])
	player.Play()
}

func (ap *AudioPlayer) PlayFromReader(audioName string) error {
	player, err := ap.audioContext.NewPlayer(allAudioReaders[audioName])
	if err != nil {
		return fmt.Errorf("Error playing from audio reader [name:%s]: %v", audioName, err)
	}

	player.Play()

	return nil
}
