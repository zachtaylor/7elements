package engine

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/log"
)

func Attack(game *game.T) game.Event {
	game.Log().Info("attack")
	return &AttackEvent{}
}

type AttackEvent struct {
	Card *game.Card
}

func (event *AttackEvent) Name() string {
	return "attack"
}

// // OnActivate implements game.ActivateEventer
// func (event *AttackEvent) OnActivate(*game.T) {
// }

// // OnConnect implements game.ConnectEventer
// func (event *AttackEvent) OnConnect(*game.T, *game.Seat) {
// }

// // Finish implements game.FinishEventer
// func (event *AttackEvent) Finish(*game.T) {
// }

// // GetStack implements game.StackEventer
// func (event *AttackEvent) GetStack(game *game.T) *game.State {
// 	return nil
// }

func (event *AttackEvent) GetNext(game *game.T) *game.State {
	if event.Card != nil {
		return game.NewState(game.State.Seat, Defend(game, event))
	}
	return game.NewState(game.State.Seat, Sunset(game))
}

func (event *AttackEvent) Json(game *game.T) vii.Json {
	if event.Card == nil {
		return nil
	}
	return event.Card.Json()
}

// Request implements game.RequestEventer
func (event *AttackEvent) Request(game *game.T, seat *game.Seat, json vii.Json) {
	if seat.Username != game.State.Seat {
		game.Log().WithFields(log.Fields{
			"Seat": seat,
			"json": json,
		}).Warn("engine/attack: receive")
		return
	}

	gcid := json.Sval("gcid")
	if len(gcid) < 1 {
		game.Log().Error("engine/attack: gcid missing")
		return
	}

	log := game.Log().Add("Seat", seat).Add("gcid", gcid)

	if gc := seat.Present[gcid]; gc == nil {
		log.Error("engine/attack: gcid not found")
	} else if !gc.IsAwake {
		log.Warn("engine/attack: card is not awake")
		animate.GameError(seat, game, gc.Card.Name, "not awake")
	} else {
		event.Card = gc
		game.State.Reacts[seat.Username] = "attack" // triggers len(game.State.Reacts) == len(game.Seats)
	}
}
