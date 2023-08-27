package audio

import (
	"bytes"
	"fmt"
	"game/configuration"
	"io"
	"os"

	"github.com/hajimehoshi/ebiten/v2/audio/mp3"
	"github.com/hajimehoshi/ebiten/v2/audio/wav"
)

type audioType int

const (
	WAV audioType = 0
	MP3 audioType = 1
)

type audioFile struct {
	alias string
	file  string
	audioType
}

func loadAudio(alias, file string, audioType audioType) error {
	// Get bytes from audio file
	fileBytes, err := os.ReadFile(file)
	if err != nil {
		return fmt.Errorf("Error loading audio file [alias:%s] [file:%s]: %v", alias, file, err)
	}

	var reader io.Reader

	// Depending on the type we get a different reader
	switch audioType {
	case WAV:
		stream, err := wav.DecodeWithSampleRate(configuration.SampleRate, bytes.NewReader(fileBytes))
		if err != nil {
			return fmt.Errorf("Error decoding wav bytes [alias:%s] [file:%s]: %v", alias, file, err)
		}
		reader = stream
	case MP3:
		stream, err := mp3.DecodeWithSampleRate(configuration.SampleRate, bytes.NewBuffer(fileBytes))
		if err != nil {
			return fmt.Errorf("Error decoding mp3 bytes [alias:%s] [file:%s]: %v", alias, file, err)
		}
		reader = stream
	}

	// Save the reader that is going to be used with NewPlayer in PlayFromReader function
	// Multiple players can't use the same reader at the same time
	allAudioReaders[alias] = reader

	// Read all the bytes from the audio stream
	// If fileBytes is used directly then the audio will sound bad
	bytesFromStream, err := io.ReadAll(reader)
	if err != nil {
		return fmt.Errorf("Error reading bytes from audio stream [alias:%s] [file:%s]: %v", alias, file, err)
	}
	// Save the bytes that are going to be used with NewPlayerFromBytes in PlayFromBytes function
	// Multiple samples can be played at the same time
	allAudioBytes[alias] = bytesFromStream

	return nil
}
