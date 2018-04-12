package games

import (
	"elemen7s.com"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

type TriggerMode struct {
	*Card
	*vii.Power
	Stack  *Event
	Target interface{}
}

func (m *TriggerMode) Name() string {
	return "trigger"
}

func (m *TriggerMode) OnActivate(e *Event, g *Game) {
	g.Log().WithFields(log.Fields{
		"Username":  e.Username,
		"Elements":  g.GetSeat(e.Username).Elements,
		"gcid":      m.Card.Id,
		"CardId":    m.Card.Card.Id,
		"CardType":  m.Card.Card.CardType,
		"Name":      m.Card.CardText.Name,
		"PowerId":   m.Power.Id,
		"StackMode": m.Stack.Name(),
	}).Info("trigger activate")
}

func (m *TriggerMode) OnSendCatchup(*Event, *Game, *Seat) {
}

func (m *TriggerMode) Json(e *Event, g *Game, seat *Seat) js.Object {
	return js.Object{
		"gameid":   g.Id,
		"timer":    int(e.Duration.Seconds()),
		"username": seat.Username,
		"elements": seat.Elements,
		"hand":     len(seat.Hand),
		"gcid":     m.Card.Id,
	}
}

func (m *TriggerMode) OnResolve(e *Event, g *Game) {
	g.Log().WithFields(log.Fields{
		"Username":  e.Username,
		"gcid":      m.Card.Id,
		"CardId":    m.Card.Card.Id,
		"Name":      m.Card.CardText.Name,
		"PowerId":   m.Power.Id,
		"UsesTurn":  m.Power.UsesTurn,
		"StackMode": m.Stack.Name(),
	}).Info("trigger resolve")

	if m.Power.UsesTurn {
		if !m.Card.Awake {
			return
		}
		m.Card.Awake = false
		BroadcastAnimateCardUpdate(g, m.Card)
	}

	g.Active = m.Stack
	g.PowerScript(g.GetSeat(e.Username), m.Power, m.Target)
}

func (m *TriggerMode) OnReceive(event *Event, g *Game, s *Seat, j js.Object) {
	g.Log().Add("Username", s.Username).Add("EventName", j["event"]).Error("trigger: receive")
}

func Trigger(g *Game, s *Seat, c *Card, p *vii.Power, target interface{}) {
	if p.UsesTurn && !c.Awake {
		return
	} else if !s.Elements.GetActive().Test(p.Costs) {
		return
	}
	s.Elements.Deactivate(p.Costs)

	e := NewEvent(s.Username)
	e.Target = c.CardText.Name
	e.EMode = &TriggerMode{
		Card:   c,
		Power:  p,
		Target: target,
		Stack:  g.Active,
	}
	g.TimelineJoin(e)
}
