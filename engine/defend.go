package engine

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func Defend(game *vii.Game, a AttackEvent) vii.GameEvent {
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

func (event *DefendEvent) Priority(game *vii.Game) bool {
	return len(game.State.Reacts) < len(game.Seats)
}

func (event *DefendEvent) Receive(game *vii.Game, seat *vii.GameSeat, json js.Object) {
	if seat.Username != game.State.Seat {
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

	seat.WriteJson(animate.Build("/game/"+event.Name(), game.State.Event.Json(game)))
}

func (event *DefendEvent) OnReconnect(*vii.Game, *vii.GameSeat) {
}

func (event *DefendEvent) OnStart(*vii.Game) {
}

func (event *DefendEvent) NextEvent(game *vii.Game) vii.GameEvent {
	for gcid, name := range event.AttackEvent {
		seat := game.GetSeat(name)
		acard := game.Cards[gcid]
		isBlocked := false
		for dcarddid, dcardaid := range event.DefendOptions {
			if gcid == dcardaid {
				isBlocked = true
				dcard := game.Cards[dcarddid]
				Combat(game, acard, dcard)
			}
		}
		if !isBlocked {
			if seat.Life > acard.Body.Attack {
				seat.Life -= acard.Body.Attack
			} else {
				seat.Life = 0
				return End(game, acard.Username, seat.Username)
			}
		}
	}

	return Main(game)
}

func (event *DefendEvent) Json(game *vii.Game) vii.Json {
	return js.Object{
		"gameid":   game.Key,
		"username": game.State.Seat,
		"timer":    game.State.Timer.Seconds(),
		"attacks":  event.AttackEvent.Json(game),
		"defends":  event.DefendOptions,
	}
}
