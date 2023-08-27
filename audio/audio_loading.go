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

var allAudioFiles = []*audioFile{
	{alias: "enemy_dead", file: "resources/audios/enemy_dead.wav", audioType: WAV},
	{alias: "player_hit", file: "resources/audios/player_hit.wav", audioType: WAV},
	{alias: "game_over", file: "resources/audios/game_over.wav", audioType: WAV},
	{alias: "shot", file: "resources/audios/shot.wav", audioType: WAV},
}

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

func loadAllAudios() error {
	for _, audioFile := range allAudioFiles {
		if err := loadAudio(audioFile.alias, audioFile.file, audioFile.audioType); err != nil {
			return err
		}
	}

	return nil
}

func loadAudio(name, file string, audioType audioType) error {
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
