package game

import (
	"time"

	"ztaylor.me/cast"
)

type State struct {
	id     string
	Timer  time.Duration
	Reacts map[string]string
	Event  Event
	Stack  *State
}

func (s *State) ID() string {
	return s.id
}

func (s *State) EventName() string {
	return s.Event.Name()
}

func (s *State) String() string {
	return `game.State{` + s.Print() + `}`
}

// Print returns a detailed compressed string representation
func (s *State) Print() string {
	return s.Event.Name() + `#` + s.id + `(` + s.Event.Seat() + `)`
}

// JSON returns a a representation of game state
func (s *State) JSON() cast.JSON {
	return cast.JSON{
		"id":     s.ID(),
		"name":   s.Event.Name(),
		"seat":   s.Event.Seat(),
		"data":   s.Event.JSON(),
		"timer":  int(s.Timer.Seconds()),
		"reacts": s.Reacts,
	}
}
