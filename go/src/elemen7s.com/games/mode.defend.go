package games

import (
	"ztaylor.me/js"
	"ztaylor.me/log"
)

type DefendOptions map[int]int

type DefendMode struct {
	AttackOptions
	DefendOptions
}

func (m *DefendMode) Name() string {
	return "defend"
}

func (m *DefendMode) OnActivate(e *Event, g *Game) {
}

func (m *DefendMode) Json(e *Event, g *Game, s *Seat) js.Object {
	return js.Object{
		"gameid":   g.Id,
		"username": e.Username,
		"timer":    int(e.Duration.Seconds()),
		"attacks":  m.AttackOptions,
		"defends":  m.DefendOptions,
	}
}

func (m *DefendMode) OnResolve(e *Event, g *Game) {
	for gcid, name := range m.AttackOptions {
		seat := g.GetSeat(name)
		acard := g.Cards[gcid]
		isBlocked := false
		for gcid_d1, gcid_d2 := range m.DefendOptions {
			if gcid == gcid_d2 {
				isBlocked = true
				dcard := g.Cards[gcid_d1]
				Combat(g, acard, dcard)
			}
		}
		if !isBlocked {
			if seat.Life <= acard.Attack {
				seat.Life = 0
				g.Lose(seat)
			} else {
				seat.Life -= acard.Attack
			}
		}
	}

	if g.Results == nil {
		Sunset(g)
	} else {
		End(g)
	}
}

func (m *DefendMode) OnReceive(e *Event, g *Game, s *Seat, j js.Object) {
	if j["event"] == "defend" && s.Username == e.Username {
		go m.defend(e, g, s, j)
	} else {
		log.Add("GameId", g.Id).Add("Seat", s.Username).Add("Event", j["event"]).Warn("games.Defend: receive sync error")
	}
}

func (m *DefendMode) defend(e *Event, g *Game, s *Seat, j js.Object) {
	gcid := j.Ival("gcid")
	if gcid < 1 {
		g.Log().Error("games.Defend: gcid missing")
		return
	}

	target := j.Ival("target")
	if target < 1 {
		log.Error("games.Defend: receive target")
		return
	}

	log := log.Add("Username", s.Username).Add("GameId", g.Id).Add("gcid", gcid).Add("Target", target)

	if m.DefendOptions[gcid] != 0 {
		delete(m.DefendOptions, gcid)
	} else if gc := g.Cards[gcid]; gc == nil {
		log.Error("games.Defend: gcid not found")
	} else if !gc.Awake {
		log.Warn("games.Defend: card is not awake")

		s.Send("alert", js.Object{
			"class":    "error",
			"gameid":   g.Id,
			"username": e.Username,
			"message":  gc.Text.Name + " is not awake",
		})
	} else {
		m.DefendOptions[gcid] = target
	}

	s.Send(m.Name(), m.Json(e, g, s))
}

func Defend(g *Game, a AttackOptions) {
	e := NewEvent(g.TurnClock.Next.Username)
	e.Target = "defend"
	e.EMode = &DefendMode{
		AttackOptions: a,
		DefendOptions: DefendOptions{},
	}
	go func() {
		g.Timeline <- e
	}()
}
