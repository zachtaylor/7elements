package games

import (
	"elemen7s.com/elements"
	"ztaylor.me/js"
)

type ElementMode struct {
	elements.Element
}

func (m *ElementMode) Name() string {
	return "element"
}

func (m *ElementMode) Json(e *Event, g *Game, s *Seat) js.Object {
	return js.Object{
		"name":     "open element",
		"username": e.Username,
		"element":  *m,
	}
}

func (m *ElementMode) OnActivate(e *Event, g *Game) {
}

func (m *ElementMode) OnResolve(e *Event, g *Game) {
	g.Broadcast("animate", js.Object{
		"animate":  "add element",
		"username": e.Username,
		"element":  m.Element,
	})
}

func (m *ElementMode) OnReceive(e *Event, g *Game, s *Seat, j js.Object) {
	log := g.Log().Add("Username", s.Username)
	if m.Element != elements.Null {
		log.Add("SavedEl", m.Element).Add("ElementId", j["elementid"]).Warn("element: choice already saved")
		return
	} else if j["event"] == "element" {
		log.Warn("element: event unrecognized")
		return
	} else if s.Username != e.Username {
		log.Add("ExpectedUsername", e.Username).Warn("element: username rejected")
	}
	elementId := j.Ival("elementid")
	if elementId < 1 || elementId > 7 {
		log.Add("ElementId", elementId).Warn("element: elementid out of bounds")
		return
	}
	m.Element = elements.Elements[int(elementId)]
	s.Elements.Append(m.Element)
	log.Add("Element", m.Element).Info("element")
	e.Timeout()
}

func OpenElement(g *Game) {
	e := NewEvent(g.TurnClock.Username)
	e.EMode = &ElementMode{elements.Null}
	g.TimelineJoin(e)
}
