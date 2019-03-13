package engine

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func Defend(game *game.T, a *AttackEvent) game.Event {
	game.Log().Info("defend")
	return &DefendEvent{a, nil}
}

type DefendEvent struct {
	*AttackEvent
	DefendCard *game.Card
}

func (event *DefendEvent) Name() string {
	return "defend"
}

// // OnActivate implements game.ActivateEventer
// func (event *DefendEvent) OnActivate(*game.T) {
// }

// // OnConnect implements game.ConnectEventer
// func (event *DefendEvent) OnConnect(*game.T, *game.Seat) {
// }

// // Finish implements game.FinishEventer
// func (event *DefendEvent) Finish(*game.T) {
// }

// // GetStack implements game.StackEventer
// func (event *DefendEvent) GetStack(game *game.T) *game.State {
// 	return nil
// }

// GetStack implements game.StackEventer
func (event *DefendEvent) GetNext(game *game.T) *game.State {
	return game.NewState(game.State.Seat, Combat(game, event.AttackEvent.Card, event.DefendCard))
}

func (event *DefendEvent) Json(game *game.T) vii.Json {
	return js.Object{
		"gameid":   game.ID(),
		"username": game.State.Seat,
		"timer":    game.State.Timer.Seconds(),
		"attack":   event.AttackEvent.Card.Json(),
		"defend":   event.DefendCard.Json(),
	}
}

// Request implements game.RequestEventer
func (event *DefendEvent) Request(game *game.T, seat *game.Seat, json js.Object) {
	if seat.Username == game.State.Seat {
		game.Log().WithFields(log.Fields{
			"Seat":  seat,
			"Event": json["event"],
		}).Warn("engine/defend: receive")
		return
	}

	gcid := json.Sval("gcid")
	if gcid == "" {
		game.Log().Warn("engine/defend: gcid missing")
	}

	target := json.Sval("target")
	if target != "" {
		game.Log().Warn("engine/defend: target missing")
	}

	log := game.Logger.WithFields(log.Fields{
		"Seat":   seat,
		"Event":  json["event"],
		"gcid":   gcid,
		"Target": target,
	})

	if gc := seat.Present[gcid]; gc == nil {
		log.Error("engine/defend: gcid not found")
	} else if !gc.IsAwake {
		log.Warn("engine/defend: card is not awake")
		animate.GameError(seat, game, gc.Card.Name, "not awake")
	} else {
		event.DefendCard = gc
	}

	game.State.Reacts[seat.Username] = "defend" // triggers len(game.State.Reacts) == len(game.Seats)
}
