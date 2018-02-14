package games

import (
	"ztaylor.me/js"
)

type MainMode bool

func (m MainMode) Name() string {
	return "main"
}

func (m MainMode) OnActivate(e *Event, g *Game) {
	g.Log().Add("Username", e.Username).Add("Elements", g.GetSeat(e.Username).Elements.String()).Info("main")
}

func (m MainMode) Json(event *Event, game *Game, seat *Seat) js.Object {
	return js.Object{
		"gameid":   game.Id,
		"timer":    int(event.Duration.Seconds()),
		"username": event.Username,
		"life":     seat.Life,
		"hand":     len(seat.Hand),
		"deck":     len(seat.Deck.Cards),
		"elements": seat.Elements,
		"spent":    len(seat.Graveyard),
	}
}

func (m MainMode) OnResolve(event *Event, g *Game) {
	if g.Results != nil {
		go End(g)
		return
	}

	Attack(g)
}

func (m MainMode) OnReceive(e *Event, g *Game, s *Seat, j js.Object) {
	if j["event"] == "main" {
		TryPlay(e, g, s, j, false)
	} else {
		g.Log().Add("Username", s.Username).Add("EventName", j["event"]).Error("main: receive")
	}
}

func Main(g *Game) {
	e := NewEvent(g.TurnClock.Username)
	e.EMode = MainMode(true)
	go func() {
		g.Timeline <- e
	}()
}
