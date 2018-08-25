package engine

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func Attack(game *vii.Game, past *Timeline) Event {
	if tname := past.Name(); tname != "main" {
		game.Log().Add("Timeline", tname).Error("attack can only follow main")
		return nil
	}
	game.Log().Info("attack")

	return AttackEvent{}
}

type AttackEvent map[string]string

func (event AttackEvent) Name() string {
	return "attack"
}

func (event AttackEvent) Priority(game *vii.Game, t *Timeline) bool {
	return t.HasPause() || t.Reacts[t.HotSeat] != "pass"
}

func (event AttackEvent) Receive(game *vii.Game, t *Timeline, seat *vii.GameSeat, json js.Object) {
	if seat.Username != t.HotSeat {
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

		animate.Error(seat, game, gc.CardText.Name, "not awake")
	} else {
		for _, s2 := range game.Seats {
			if s2 != seat {
				event[gcid] = s2.Username
			}
		}
	}

	seat.Send("animate", js.Object{
		"animate":       "attack options",
		"attackoptions": event,
	})
}

func (event AttackEvent) OnReconnect(*vii.Game, *Timeline, *vii.GameSeat) {
}

func (event AttackEvent) OnStart(*vii.Game, *Timeline) {
}

func (event AttackEvent) OnStop(game *vii.Game, t *Timeline) *Timeline {
	return t.Fork(game, Defend(game, t, event))
}

func (event AttackEvent) Json(game *vii.Game, t *Timeline) js.Object {
	json := js.Object{}
	for k, v := range event {
		json[k] = v
	}
	return json
}
