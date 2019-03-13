package engine

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func Play(game *game.T, seat *game.Seat, card *game.Card, target interface{}) game.Event {
	game.Log().Info("play")

	return &PlayEvent{
		Stack:  game.State,
		Card:   card,
		Target: target,
	}
}

type PlayEvent struct {
	Stack       *game.State
	Card        *game.Card
	IsCancelled bool
	Target      interface{}
}

func (event *PlayEvent) Name() string {
	return "play"
}

// // OnActivate implements game.ActivateEventer
// func (event *PlayEvent) OnActivate(*game.T) {
// }

// // OnConnect implements game.ConnectEventer
// func (event *PlayEvent) OnConnect(*game.T, *game.Seat) {
// }

// // Request implements game.RequestEventer
// func (event *PlayEvent) Request(*game.T, *game.Seat, vii.Json) {
// }

// Finish implements game.FinishEventer
func (event *PlayEvent) Finish(game *game.T) {
	seat := game.GetSeat(game.State.Seat)
	log := game.Log().WithFields(log.Fields{
		"Seat": seat,
		"Card": event.Card,
		"Hand": seat.Hand,
	})

	game.WriteJson(animate.Build("/game/resolve", js.Object{
		"gameid":   game.ID(),
		"username": seat.Username,
		"card":     event.Card.Json(),
	}))

	playPower := event.Card.Card.GetPlayPower()
	if playPower == nil {
		animate.GameError(game, game, event.Card.Card.Name, "card does not work yet")
		log.Warn("engine/play: resolve: card does not work")
	}

	if event.Card.Card.Type == vii.CTYPbody || event.Card.Card.Type == vii.CTYPitem {
		seat.Present[event.Card.Id] = event.Card
		animate.GameSpawn(game, event.Card)
	} else if event.Card.Card.Type == vii.CTYPspell {
		seat.Past[event.Card.Id] = event.Card
	}

	if playPower != nil {
		Script(game, seat, playPower, event.Target)
	}
}

// GetStack implements game.StackEventer
func (event *PlayEvent) GetStack(game *game.T) *game.State {
	return event.Stack
}

func (event *PlayEvent) GetNext(game *game.T) *game.State {
	return nil
}

func (event *PlayEvent) Json(game *game.T) vii.Json {
	return js.Object{
		"stack":  event.Stack.Json(game),
		"card":   event.Card.Json(),
		"target": event.Target,
	}
}
