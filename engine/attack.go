package engine

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func Attack(game *vii.Game) vii.GameEvent {
	game.Log().Info("attack")
	return AttackEvent{}
}

type AttackEvent map[string]string

func (event AttackEvent) Name() string {
	return "attack"
}

func (event AttackEvent) Receive(game *vii.Game, seat *vii.GameSeat, json vii.Json) {
	if seat.Username != game.State.Seat {
		game.Log().WithFields(log.Fields{
			"Seat":  seat,
			"Event": json["event"],
		}).Warn("engine-attack: receive")
		return
	}
	gcid := json.Sval("gcid")
	if len(gcid) < 1 {
		game.Log().Error("games.Attack: gcid missing")
		return
	}

	log := game.Log().Add("Seat", seat).Add("gcid", gcid)

	if event[gcid] != "" {
		delete(event, gcid)
	} else if gc := game.Cards[gcid]; gc == nil {
		log.Error("games.Attack: gcid not found")
	} else if !gc.IsAwake {
		log.Warn("games.Attack: card is not awake")

		animate.GameError(seat, game, gc.Card.Name, "not awake")
	} else {
		for _, s2 := range game.Seats {
			if s2 != seat {
				event[gcid] = s2.Username
			}
		}
	}

	seat.WriteJson(animate.Build("/animate", js.Object{
		"animate":       "attack options",
		"attackoptions": event,
	}))
}

func (event AttackEvent) OnReconnect(*vii.Game, *vii.GameSeat) {
}

func (event AttackEvent) OnStart(*vii.Game) {
}

func (event AttackEvent) NextEvent(game *vii.Game) vii.GameEvent {
	return Defend(game, event)
}

func (event AttackEvent) Json(game *vii.Game) vii.Json {
	json := vii.Json{}
	for k, v := range event {
		json[k] = v
	}
	return json
}
