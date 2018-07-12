package engine

import (
	"elemen7s.com"
	"time"
	"ztaylor.me/js"
	"ztaylor.me/keygen"
)

type Timeline struct {
	key      string
	HotSeat  string
	event    Event
	Lifetime time.Duration
	Reacts   map[string]string
}

func NewTimeline(hotseat string, d time.Duration, e Event) *Timeline {
	return &Timeline{
		key:      keygen.NewVal(),
		HotSeat:  hotseat,
		event:    e,
		Lifetime: d,
		Reacts:   make(map[string]string),
	}
}

func (t *Timeline) Fork(game *vii.Game, e Event) *Timeline {
	return NewTimeline(t.HotSeat, game.Settings.Timeout, e)
}

func (t Timeline) Key() string {
	return t.key
}

func (t Timeline) Name() string {
	return t.event.Name()
}

func (t Timeline) String() string {
	return t.Key() + ":" + t.Name()
}

// Receive will handle the underlying GameEvent type requests
func (t *Timeline) Receive(game *vii.Game, seat *vii.GameSeat, json js.Object) {
	t.event.Receive(game, t, seat, json)
}

// Receive will check the underlying GameEvent Priority
func (t *Timeline) Priority(game *vii.Game) bool {
	return t.event.Priority(game, t)
}

// OnReconnect will call the underlying GameEvent OnReconnect
func (t *Timeline) OnReconnect(game *vii.Game, seat *vii.GameSeat) {
	t.event.OnReconnect(game, t, seat)
}

// OnStop will call the underlying GameEvent OnStop
func (t *Timeline) OnStop(game *vii.Game) *Timeline {
	return t.event.OnStop(game, t)
}

func (t *Timeline) HasPass() bool {
	return t.HasReact("pass")
}
func (t *Timeline) HasPause() bool {
	return t.HasReact("pause")
}
func (t *Timeline) HasReact(react string) bool {
	for _, v := range t.Reacts {
		if v == react {
			return true
		}
	}
	return false
}
