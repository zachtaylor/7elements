package game

import "time"

type Settings struct {
	Timeout time.Duration
}

func NewDefaultSettings() *Settings {
	settings := &Settings{}
	settings.Timeout = 7 * time.Minute
	return settings
}
