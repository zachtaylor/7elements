package games

import (
	"time"
	"ztaylor.me/events"
	"ztaylor.me/js"
)

func init() {
	gameIdGen.Seed(time.Now().Unix())
	events.On("GameStart", func(args ...interface{}) {
		Start(args[0].(*Game))
	})
}

type StartMode bool

func (m StartMode) Name() string {
	return "start"
}

func (m StartMode) OnActivate(e *Event, g *Game) {
	names := make([]string, 0)
	for _, seat := range g.Seats {
		seat.Start()
		names = append(names, seat.Username)
	}
	g.TurnClock = BuildTurnClock(names)
}

func (m StartMode) OnSendCatchup(*Event, *Game, *Seat) {
}

func (m StartMode) Json(e *Event, g *Game, s *Seat) js.Object {
	return js.Object{
		"gameid": g.Id,
		"timer":  int(e.Duration.Seconds()),
	}
}

func (m StartMode) OnResolve(e *Event, g *Game) {
	Sunrise(g)
}

func (m StartMode) OnReceive(e *Event, g *Game, s *Seat, j js.Object) {
	log := g.Log().Add("Username", s.Username).Add("Resp", j["choice"])

	if e.Resp[s.Username] != "" {
		log.Add("Val", e.Resp[s.Username]).Warn("start: receive already recorded")
		return
	} else if j["choice"] == "keep" {
		e.Resp[s.Username] = "keep"
	} else if j["choice"] == "mulligan" {
		e.Resp[s.Username] = "mulligan"
		s.DiscardHand()
		s.DrawCard(3)
		AnimateHand(g, s)
		BroadcastAnimateSeatUpdate(g, s)
	} else {
		log.Warn("start: receive unrecognized")
		return
	}

	log.Debug("start: receive")
	g.Broadcast("alert", js.Object{
		"class":    "tip",
		"gameid":   g.Id,
		"username": s.Username,
		"message":  j["choice"],
		"timer":    1000,
	})
	g.Broadcast("pass", js.Object{
		"gameid":   g.Id,
		"username": s.Username,
	})

	for _, seatx := range g.Seats {
		if e.Resp[seatx.Username] == "" {
			return
		}
	}

	e.Timeout()
}

func Start(g *Game) {
	if len(g.Events) > 0 {
		g.Log().Error("games.Start: already started")
		return
	}

	Cache.Add(g)

	e := NewEvent("start")
	e.EMode = StartMode(true)
	// special handling because we don't have previous active event
	g.Active = e
	g.Start()
	e.Activate(g)
}
