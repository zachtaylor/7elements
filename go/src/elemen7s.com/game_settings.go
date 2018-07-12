package vii

import (
	"time"
)

type GameSettings struct {
	Timeout time.Duration
	Teams   [][]string
}

func NewGameSettings() *GameSettings {
	return &GameSettings{
		Teams: make([][]string, 0),
	}
}

func NewDefaultGameSettings() *GameSettings {
	settings := NewGameSettings()
	settings.Timeout = 7 * time.Minute
	return settings
}
