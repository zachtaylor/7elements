package scripts

import (
	"elemen7s.com"
	"elemen7s.com/games"
	"ztaylor.me/js"
)

func init() {
	games.Scripts["new-element"] = NewElement
}

type NewElementMode struct {
	vii.Element
	*games.Stack
}

func (mode *NewElementMode) Name() string {
	return "choice"
}

func (mode *NewElementMode) Json(e *games.Event, g *games.Game, s *games.Seat) js.Object {
	return js.Object{
		"gameid":   g.Id,
		"choice":   "New Element",
		"username": s.Username,
		"timer":    int(e.Duration.Seconds()),
	}
}

func (mode *NewElementMode) OnActivate(e *games.Event, g *games.Game) {
	games.AnimateNewElementChoice(g.GetSeat(e.Username), g)
}

func (mode *NewElementMode) OnSendCatchup(e *games.Event, g *games.Game, s *games.Seat) {
	if e.Username == s.Username {
		games.AnimateNewElementChoice(s, g)
	}
}

func (mode *NewElementMode) OnResolve(e *games.Event, g *games.Game) {
	seat := g.GetSeat(e.Username)
	seat.Elements.Append(mode.Element)
	games.BroadcastAnimateAddElement(g, e.Username, int(mode.Element))
	mode.Stack.OnResolve(e, g)
}

func (mode *NewElementMode) OnReceive(e *games.Event, g *games.Game, s *games.Seat, json js.Object) {
	if s.Username != e.Username {
		g.Log().Add("Username", s.Username).Add("HotSeat", e.Username).Warn("games.NewElementMode: not your choice")
		return
	}

	if e := json.Ival("choice"); e < 1 || e > 7 {
		g.Log().Add("Username", s.Username).Add("Element", e).Warn("games.NewElementMode: invalid element")
		return
	} else {
		mode.Element = vii.Element(e)
		g.Log().Add("Username", s.Username).Add("Element", mode.Element).Info("games.NewElement: confirmed element")
	}

	g.TimelineJoin(nil)
}

func NewElement(g *games.Game, s *games.Seat, target interface{}) {
	event := games.NewEvent(s.Username)
	event.EMode = &NewElementMode{
		Stack: games.StackEvent(g.Active),
	}
	g.TimelineJoin(event)
}
