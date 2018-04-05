package games

import (
	"elemen7s.com"
	"ztaylor.me/js"
)

type PlayMode struct {
	*Card
	*Stack
	Target interface{}
}

func (m *PlayMode) Name() string {
	return "play"
}

func (m *PlayMode) OnActivate(e *Event, g *Game) {
	g.Log().Add("Username", e.Username).Add("Elements", g.GetSeat(e.Username).Elements).Add("gcid", m.Card.Id).Add("Name", m.Card.CardText.Name).Add("Target", m.Target).Info("play activate")
}

func (m *PlayMode) OnSendCatchup(*Event, *Game, *Seat) {
}

func (m *PlayMode) Json(e *Event, g *Game, seat *Seat) js.Object {
	return js.Object{
		"gameid":   g.Id,
		"timer":    int(e.Duration.Seconds()),
		"username": seat.Username,
		"elements": seat.Elements,
		"hand":     len(seat.Hand),
		"card":     m.Card.Json(),
		"target":   m.Target,
	}
}

func (m *PlayMode) OnResolve(e *Event, g *Game) {
	log := g.Log().Add("Username", e.Username).Add("gcid", m.Card.Id).Add("CardId", m.Card.Card.Id).Add("CardType", m.Card.Card.CardType).Add("Name", m.Card.CardText.Name)
	seat := g.GetSeat(e.Username)

	g.Broadcast("resolve", js.Object{
		"gameid":   g.Id,
		"username": e.Username,
		"card":     m.Card.Json(),
	})

	if m.Card.Card.CardType == vii.CTYPbody || m.Card.Card.CardType == vii.CTYPitem {
		seat.Alive[m.Card.Id] = m.Card
		BroadcastAnimateSpawn(g, m.Card)
	} else if m.Card.Card.CardType == vii.CTYPspell {
		if power := m.Card.Card.Powers[0]; power == nil {
			BroadcastAnimateAlertError(g, m.Card.CardText.Name+" does not work yet")
			log.Warn("play: resolve; card does not work")
		} else {
			g.PowerScript(m.Card.Username, power, m.Target)
		}
	} else {
		log.Warn("play: resolve; cannot resolve cardtype")
	}

	m.Stack.OnResolve(e, g)
}

func (m *PlayMode) OnReceive(event *Event, g *Game, s *Seat, j js.Object) {
	g.Log().Add("Username", s.Username).Add("EventName", j["event"]).Error("play: receive")
}

func Play(g *Game, s *Seat, c *Card, target interface{}) {
	if !s.RemoveHandAndElements(c.Id) {
		return
	}

	e := NewEvent(s.Username)
	e.Target = c.CardText.Name
	e.EMode = &PlayMode{
		Card:   c,
		Target: target,
		Stack:  StackEvent(g.Active),
	}
	g.TimelineJoin(e)
}
