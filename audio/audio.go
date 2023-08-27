package audio

import (
	"bytes"
	"fmt"
	"game/configuration"
	"io"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
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

type AudioType int

const (
	WAV AudioType = 0
	MP3 AudioType = 1
)

func loadAllAudios() error {
	if err := loadAudio("enemy_dead", "resources/enemy_dead.mp3", MP3); err != nil {
		return err
	}
	if err := loadAudio("player_hit", "resources/player_hit.wav", WAV); err != nil {
		return err
	}
	if err := loadAudio("game_over", "resources/game_over.wav", WAV); err != nil {
		return err
	}

	return nil
}

func loadAudio(name, file string, audioType AudioType) error {
	fileBytes, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("Error loading audio file [name:%s] [file:%s]: %v", name, file, err)
	}

	var reader io.Reader

	switch audioType {
	case WAV:
		stream, err := wav.DecodeWithSampleRate(configuration.SampleRate, bytes.NewReader(fileBytes))
		if err != nil {
			return fmt.Errorf("Error decoding wav bytes [name:%s] [file:%s]: %v", name, file, err)
		}
		reader = stream
	case MP3:
		stream, err := mp3.DecodeWithSampleRate(configuration.SampleRate, bytes.NewBuffer(fileBytes))
		if err != nil {
			return fmt.Errorf("Error decoding mp3 bytes [name:%s] [file:%s]: %v", name, file, err)
		}
		reader = stream
	}

	allAudioReaders[name] = reader

	bytesFromStream, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("Error reading bytes from audio stream [name:%s] [file:%s]: %v", name, file, err)
	}
	allAudioBytes[name] = bytesFromStream

	return nil
}
