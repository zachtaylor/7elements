package games

import (
	"ztaylor.me/js"
)

type AttackOptions map[int]string

type AttackMode AttackOptions

func (a AttackMode) Name() string {
	return "attack"
}

func (a AttackMode) OnActivate(e *Event, g *Game) {
}

func (a AttackMode) Json(e *Event, g *Game, s *Seat) js.Object {
	return js.Object{
		"gameid":        g.Id,
		"username":      s.Username,
		"timer":         int(e.Duration.Seconds()),
		"attackoptions": a,
	}
}

func (a AttackMode) OnResolve(e *Event, g *Game) {
	Defend(g, AttackOptions(a))
}

func (a AttackMode) OnReceive(e *Event, g *Game, s *Seat, j js.Object) {
	if s.Username == e.Username {
		go a.attack(e, g, s, j)
	} else {
		g.Log().Add("Seat", s.Username).Add("Event", j["event"]).Warn("games.AttackMode: receive sync error")
	}
}

func (a AttackMode) attack(e *Event, g *Game, s *Seat, j js.Object) {
	gcid := j.Ival("gcid")
	if gcid < 1 {
		g.Log().Error("games.Attack: gcid missing")
		return
	}

	log := g.Log().Add("Username", s.Username).Add("gcid", gcid)

	if a[gcid] != "" {
		delete(a, gcid)
	} else if gc := g.Cards[gcid]; gc == nil {
		log.Error("games.Attack: gcid not found")
	} else if !gc.Awake {
		log.Warn("games.Attack: card is not awake")

		AnimateAlertError(s, g, gc.Text.Name, "not awake")
	} else {
		for _, s2 := range g.Seats {
			if s2 != s {
				a[gcid] = s2.Username
			}
		}
	}

	AnimateAttack(s, AttackOptions(a))
}

func Attack(g *Game) {
	e := NewEvent(g.TurnClock.Username)
	e.Target = "attack"
	e.EMode = AttackMode{}
	g.TimelineJoin(e)
}
