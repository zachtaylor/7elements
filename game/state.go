package game

import (
	"time"

	"github.com/zachtaylor/7elements"
)

type State struct {
	id     string
	Stack  *State
	Seat   string
	Timer  time.Duration
	Reacts map[string]string
	Event  Event
}

func (s *State) ID() string {
	return s.id
}

func (s *State) EventName() string {
	return s.Event.Name()
}

func (s State) String() string {
	return `state#` + s.id + `:` + s.EventName()
}

func (s *State) Json(game *T) vii.Json {
	return vii.Json{
		"id":     s.id,
		"gameid": game.id,
		"event": vii.Json{
			"name": s.Event.Name(),
			"data": s.Event.Json(game),
		},
		"seat":   s.Seat,
		"timer":  int(s.Timer.Seconds()),
		"reacts": s.Reacts,
	}
}
