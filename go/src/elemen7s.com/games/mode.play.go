package games

import (
	"elemen7s.com/cards/types"
	"ztaylor.me/js"
)

type PlayMode struct {
	Card  *Card
	Stack *Event
}

func (m *PlayMode) Name() string {
	return "play"
}

func (m *PlayMode) OnActivate(e *Event, g *Game) {
	log := g.Log().Add("Username", e.Username).Add("Elements", g.GetSeat(e.Username).Elements.String()).Add("gcid", m.Card.Id)
	if err := m.removeCardAndElements(g.GetSeat(e.Username)); err != nil {
		log.Add("Error", err).Error("play: activate failed")
		e.Timeout()
	} else {
		log.Add("Name", m.Card.Text.Name).Info("play")
	}
}

func (m *PlayMode) Json(e *Event, g *Game, seat *Seat) js.Object {
	return js.Object{
		"gameid":   g.Id,
		"timer":    int(e.Duration.Seconds()),
		"username": seat.Username,
		"elements": seat.Elements,
		"hand":     len(seat.Hand),
		"card":     m.Card.Json(),
	}
}

func (m *PlayMode) OnResolve(e *Event, g *Game) {
	log := g.Log().Add("Username", e.Username).Add("gcid", m.Card.Id).Add("CardId", m.Card.Card.Id).Add("CardType", m.Card.Card.CardType)
	seat := g.GetSeat(e.Username)

	g.Broadcast("resolve", js.Object{
		"gameid":   g.Id,
		"username": e.Username,
		"card":     m.Card.Json(),
	})

	if m.Card.Card.CardType == ctypes.Body || m.Card.Card.CardType == ctypes.Item {
		seat.Alive[m.Card.Id] = m.Card
		g.Broadcast("spawn", js.Object{
			"gameid":   g.Id,
			"username": e.Username,
			"card":     m.Card.Json(),
		})
	} else if m.Card.Card.CardType == ctypes.Spell {
		if power := m.Card.Card.Powers[0]; power == nil {
			g.Broadcast("alert", js.Object{
				"class":    "error",
				"gameid":   g.Id,
				"username": e.Username,
				"message":  m.Card.Text.Name + " does not work yet",
			})
			log.Warn("play: resolve; card does not work")
		} else if script := Scripts[power.Script]; script == nil {
			g.Broadcast("alert", js.Object{
				"class":    "error",
				"gameid":   g.Id,
				"username": e.Username,
				"message":  m.Card.Text.Name + " does not work yet",
			})
			log.Warn("play: resolve; card does not work")
		} else {
			log.Info("play")
			script(g, seat)

			if g.Results != nil {
				return
			}
		}
	} else {
		log.Warn("play: resolve; cannot resolve cardtype")
	}

	g.Active = m.Stack
	m.Stack.Activate(g)
}

func (m *PlayMode) OnReceive(event *Event, g *Game, s *Seat, j js.Object) {
	g.Log().Add("Username", s.Username).Add("EventName", j["event"]).Error("play: receive")
}

func (m *PlayMode) removeCardAndElements(seat *Seat) error {
	seat.Elements.Deactivate(m.Card.Card.Costs)
	delete(seat.Hand, m.Card.Id)
	return nil
}

func Play(stack *Event, g *Game, c *Card, seat *Seat) {
	e := NewEvent(seat.Username)
	e.Target = c.Text.Name
	e.EMode = &PlayMode{
		Card:  c,
		Stack: stack,
	}
	g.TimelineJoin(e)
}
