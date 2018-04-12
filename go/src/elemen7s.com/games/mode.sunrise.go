package games

import (
	"elemen7s.com"
	"ztaylor.me/js"
)

type SunriseMode struct {
	vii.Element
}

func (m *SunriseMode) Name() string {
	return "sunrise"
}

func (m *SunriseMode) OnActivate(e *Event, g *Game) {
}

func (m *SunriseMode) OnSendCatchup(*Event, *Game, *Seat) {
}

func (m *SunriseMode) Json(e *Event, g *Game, s *Seat) js.Object {
	return js.Object{
		"gameid":   g.Id,
		"timer":    int(e.Duration.Seconds()),
		"username": e.Username,
	}
}

func (m *SunriseMode) OnResolve(e *Event, g *Game) {
	if m.Element == vii.ELEMnull {
		g.Log().Warn("games.Sunrise: !resolve forfeit")
		g.Results = &Results{
			Losers:  []string{g.TurnClock.Username},
			Winners: []string{g.TurnClock.Next.Username},
		}
		End(g)
	} else if seat := g.GetSeat(e.Username); seat == nil {
		g.Log().Warn("sunrise: !resolve seat missing")
	} else {
		card := seat.Deck.Draw()
		seat.Hand[card.Id] = card
		seat.Reactivate()
		AnimateHand(g, seat)
		BroadcastAnimateSeatUpdate(g, seat)
		Main(g)
	}
}

func (m *SunriseMode) OnReceive(e *Event, g *Game, s *Seat, j js.Object) {
	log := g.Log().Add("Username", s.Username)
	if m.Element != vii.ELEMnull {
		log.Add("SavedEl", m.Element).Add("ElementId", j["elementid"]).Warn("sunrise: choice already saved")
		return
	} else if j["event"] == "element" {
		log.Warn("element: event unrecognized")
		return
	} else if s.Username != e.Username {
		log.Add("ExpectedUsername", e.Username).Warn("sunrise: username rejected")
	}
	elementId := j.Ival("elementid")
	if elementId < 1 || elementId > 7 {
		log.Add("ElementId", elementId).Warn("sunrise: elementid out of bounds")
		return
	}
	m.Element = vii.Elements[int(elementId)]
	s.Elements.Append(m.Element)
	log.Add("Element", m.Element).Info("sunrise: confirm element choice")
	e.Timeout()
}

func Sunrise(g *Game) {
	e := NewEvent(g.TurnClock.Username)
	e.Target = "sunrise"
	e.EMode = &SunriseMode{}
	g.TimelineJoin(e)
}
