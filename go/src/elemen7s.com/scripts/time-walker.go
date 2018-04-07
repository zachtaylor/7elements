package scripts

import (
	"elemen7s.com"
	"elemen7s.com/games"
	"ztaylor.me/js"
)

func init() {
	games.Scripts["time-walker"] = TimeWalker
}

type TimeWalkerMode struct {
	vii.Element
	Stack *games.Event
}

func (mode *TimeWalkerMode) Name() string {
	return "choice"
}

func (mode *TimeWalkerMode) Json(e *games.Event, g *games.Game, s *games.Seat) js.Object {
	return js.Object{
		"gameid":   g.Id,
		"choice":   "Time Walker",
		"username": s.Username,
		"timer":    int(e.Duration.Seconds()),
	}
}

func (mode *TimeWalkerMode) OnActivate(e *games.Event, g *games.Game) {
	games.AnimateNewElementChoice(g.GetSeat(e.Username), g)
}

func (mode *TimeWalkerMode) OnSendCatchup(e *games.Event, g *games.Game, s *games.Seat) {
	if e.Username == s.Username {
		games.AnimateNewElementChoice(s, g)
	}
}

func (mode *TimeWalkerMode) OnResolve(e *games.Event, g *games.Game) {
	seat := g.GetSeat(e.Username)
	seat.Elements.Append(mode.Element)
	games.BroadcastAnimateAddElement(g, e.Username, int(mode.Element))
	mode.Stack.Activate(g)
}

func (mode *TimeWalkerMode) OnReceive(e *games.Event, g *games.Game, s *games.Seat, json js.Object) {
	if s.Username != e.Username {
		g.Log().Add("Username", s.Username).Add("HotSeat", e.Username).Warn("games.TimeWalkerMode: not your choice")
		return
	}

	if e := json.Ival("choice"); e < 1 || e > 7 {
		g.Log().Add("Username", s.Username).Add("Element", e).Warn("games.TimeWalkerMode: invalid element")
		return
	} else {
		mode.Element = vii.Element(e)
		g.Log().Add("Username", s.Username).Add("Element", mode.Element).Info("games.TimeWalker: confirmed element")
	}

	g.TimelineJoin(nil)
}

func TimeWalker(g *games.Game, s *games.Seat, target interface{}) {
	event := games.NewEvent(s.Username)
	event.EMode = &TimeWalkerMode{
		Stack: g.Active,
	}
	g.TimelineJoin(event)
}
