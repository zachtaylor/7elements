package engine

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func Defend(game *vii.Game, past *Timeline, a AttackEvent) Event {
	if tname := past.Name(); tname != "attack" {
		game.Log().Add("Timeline", tname).Error("defend can only follow attack")
		return nil
	}
	game.Log().Info("defend")

	return &DefendEvent{a, DefendOptions{}}
}

type DefendOptions map[string]string

type DefendEvent struct {
	AttackEvent
	DefendOptions
}

func (event *DefendEvent) Name() string {
	return "defend"
}

func (event *DefendEvent) Priority(game *vii.Game, t *Timeline) bool {
	return t.HasPause() || t.Reacts[t.HotSeat] != "pass"
}

func (event *DefendEvent) Receive(game *vii.Game, t *Timeline, seat *vii.GameSeat, json js.Object) {
	if seat.Username != t.HotSeat {
		game.Log().WithFields(log.Fields{
			"Seat":  seat,
			"Event": json["event"],
		}).Warn("engine-defend: receive")
		return
	} else if json["event"] != "defend" {
		return
	}

	gcid := json.Sval("gcid")
	if gcid == "" {
		game.Log().Error("engine-defend: gcid missing")
		return
	}

	target := json.Sval("target")
	if target != "" {
		log.Error("engine-defend: receive target")
		return
	}

	log := game.Logger.WithFields(log.Fields{
		"Seat":   seat,
		"Event":  json["event"],
		"gcid":   gcid,
		"Target": target,
	})

	if event.DefendOptions[gcid] != "" {
		delete(event.DefendOptions, gcid)
	} else if gc := game.Cards[gcid]; gc == nil {
		log.Error("engine-defend: gcid not found")
	} else if !gc.IsAwake {
		log.Warn("engine-defend: card is not awake")

		animate.GameError(seat, game, gc.Card.Name, "not awake")
	} else {
		event.DefendOptions[gcid] = target
	}

	seat.WriteJson(animate.Build("/game/"+event.Name(), event.Json(game, t)))
}

func (event *DefendEvent) OnReconnect(*vii.Game, *Timeline, *vii.GameSeat) {
}

func (event *DefendEvent) OnStart(*vii.Game, *Timeline) {
}

func (event *DefendEvent) OnStop(game *vii.Game, t *Timeline) *Timeline {
	for gcid, name := range event.AttackEvent {
		seat := game.GetSeat(name)
		acard := game.Cards[gcid]
		isBlocked := false
		for gcid_d1, gcid_d2 := range event.DefendOptions {
			if gcid == gcid_d2 {
				isBlocked = true
				dcard := game.Cards[gcid_d1]
				Combat(game, acard, dcard)
			}
		}
		if !isBlocked {
			if seat.Life > acard.Attack {
				seat.Life -= acard.Attack
			} else {
				seat.Life = 0
				game.Results = &vii.GameResults{
					Loser:  seat.Username,
					Winner: game.GetOpponentSeat(seat.Username).Username,
				}
			}
		}
	}

	return t.Fork(game, Sunset(game, t))
}

func (event *DefendEvent) Json(game *vii.Game, t *Timeline) js.Object {
	return js.Object{
		"gameid":   game,
		"username": t.HotSeat,
		"timer":    t.Lifetime.Seconds(),
		"attacks":  event.AttackEvent.Json(game, t),
		"defends":  event.DefendOptions,
	}
}
