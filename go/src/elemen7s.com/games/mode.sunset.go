package games

import (
	"ztaylor.me/js"
)

type SunsetMode bool

func (m SunsetMode) Name() string {
	return "sunset"
}

func (m SunsetMode) OnActivate(e *Event, g *Game) {
}

func (m SunsetMode) Json(e *Event, g *Game, seat *Seat) js.Object {
	return js.Object{
		"gameid":   g.Id,
		"username": e.Username,
		"timer":    int(e.Duration.Seconds()),
	}
}

func (m SunsetMode) OnResolve(e *Event, g *Game) {
	g.TurnClock = g.TurnClock.Next
	Sunrise(g)
}

func (m SunsetMode) OnReceive(e *Event, g *Game, s *Seat, j js.Object) {
	g.Log().Add("Username", s.Username).Add("EventName", j["event"]).Error("games.Sunset: receive")
}

func Sunset(g *Game) {
	e := NewEvent("start")
	e.EMode = SunsetMode(true)
	g.TimelineJoin(e)
}
