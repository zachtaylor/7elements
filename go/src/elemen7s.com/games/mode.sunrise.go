package games

import (
	"elemen7s.com"
	"ztaylor.me/js"
)

type SunriseMode struct {
	*ElementMode
}

func (m *SunriseMode) Name() string {
	return "sunrise"
}

func (m *SunriseMode) OnActivate(e *Event, g *Game) {
	m.ElementMode.OnActivate(e, g)
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
	m.ElementMode.OnResolve(e, g)
	if m.ElementMode.Element == vii.ELEMnull {
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
		AnimateAddCard(seat, card)
		Main(g)
	}
}

func (m *SunriseMode) OnReceive(e *Event, g *Game, s *Seat, j js.Object) {
	m.ElementMode.OnReceive(e, g, s, j)
}

func Sunrise(g *Game) {
	e := NewEvent(g.TurnClock.Username)
	e.Target = "sunrise"
	e.EMode = &SunriseMode{&ElementMode{}}
	g.TimelineJoin(e)
}
