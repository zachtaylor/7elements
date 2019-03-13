package engine

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func Trigger(game *game.T, seat *game.Seat, card *game.Card, p *vii.Power, target interface{}) game.Event {
	game.Log().Info("trigger")

	return &TriggerEvent{
		Stack:  game.State,
		Card:   card,
		Power:  p,
		Target: target,
	}
}

type TriggerEvent struct {
	Stack  *game.State
	Card   *game.Card
	Power  *vii.Power
	Target interface{}
}

func (event *TriggerEvent) Name() string {
	return "trigger"
}

// // OnConnect implements game.ConnectEventer
// func (event *TriggerEvent) OnConnect(*game.T, *game.Seat) {
// }

// // Request implements game.RequestEventer
// func (event *TriggerEvent) Request(game *game.T, seat *game.Seat, json js.Object) {
// }

// GetStack implements game.StackEventer
func (event *TriggerEvent) GetStack(*game.T) *game.State {
	return event.Stack
}

func (event *TriggerEvent) GetNext(game *game.T) *game.State {
	return nil
}

func (event *TriggerEvent) OnActivate(game *game.T) {
	seat := game.GetSeat(game.State.Seat)
	game.Logger.WithFields(log.Fields{
		"Username": seat.Username,
		"Elements": seat.Elements,
		"Card":     event.Card,
		"PowerId":  event.Power.Id,
		"Stack":    event.Stack.EventName(),
	}).Debug("engine/trigger: OnStart")
	animate.GameCard(game, event.Card)
}

// Finish implements game.FinishEventer
func (event *TriggerEvent) Finish(game *game.T) {
	seat := game.GetSeat(game.State.Seat)
	game.Log().WithFields(log.Fields{
		"Username": seat.Username,
		"gcid":     event.Card.Id,
		"CardId":   event.Card.Card.Id,
		"Name":     event.Card.Card.Name,
		"PowerId":  event.Power.Id,
		"UsesTurn": event.Power.UsesTurn,
		"Stack":    event.Stack.EventName(),
	}).Info("trigger resolve")

	Script(game, seat, event.Power, event.Target)
}

func (event *TriggerEvent) Json(game *game.T) js.Object {
	return js.Object{
		"card":  event.Card.Json(),
		"power": event.Power.Json(),
	}
}
