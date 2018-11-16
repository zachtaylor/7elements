package engine

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func Trigger(game *vii.Game, seat *vii.GameSeat, card *vii.GameCard, p *vii.Power, target interface{}) vii.GameEvent {
	game.Log().Info("trigger")

	return &TriggerEvent{
		Stack:  game.State.Event,
		Card:   card,
		Power:  p,
		Target: target,
	}
}

type TriggerEvent struct {
	Stack vii.GameEvent
	Card  *vii.GameCard
	*vii.Power
	Target interface{}
}

func (event *TriggerEvent) Name() string {
	return "trigger"
}

func (event *TriggerEvent) Priority(game *vii.Game) bool {
	return game.State.Reacts[game.GetOpponentSeat(game.State.Seat).Username] != "pass"
}

func (event *TriggerEvent) Receive(game *vii.Game, seat *vii.GameSeat, json js.Object) {
	game.Log().WithFields(log.Fields{
		"Seat":  seat,
		"Event": json["event"],
	}).Warn("trigger: receive")
}

func (event *TriggerEvent) OnStart(game *vii.Game) {
	seat := game.GetSeat(game.State.Seat)
	game.Log().WithFields(log.Fields{
		"Username":  seat.Username,
		"Elements":  seat.Elements,
		"Card":      event.Card,
		"PowerId":   event.Power.Id,
		"StackMode": event.Stack.Name(),
	}).Debug("engine-trigger: OnStart")
	animate.GameCard(game, event.Card)
}

func (event *TriggerEvent) OnReconnect(*vii.Game, *vii.GameSeat) {
}

func (event *TriggerEvent) NextEvent(game *vii.Game) vii.GameEvent {
	seat := game.GetSeat(game.State.Seat)
	game.Log().WithFields(log.Fields{
		"Username":  seat.Username,
		"gcid":      event.Card.Id,
		"CardId":    event.Card.Card.Id,
		"Name":      event.Card.Card.Name,
		"PowerId":   event.Power.Id,
		"UsesTurn":  event.Power.UsesTurn,
		"StackMode": event.Stack.Name(),
	}).Info("trigger resolve")

	Script(game, seat, event.Power, event.Target)
	return event.Stack
}

func (event *TriggerEvent) Json(game *vii.Game) js.Object {
	seat := game.GetSeat(game.State.Seat)
	return js.Object{
		"gameid":   game.Key,
		"timer":    game.State.Timer.Seconds(),
		"username": seat.Username,
		"elements": seat.Elements,
		"hand":     len(seat.Hand),
		"gcid":     event.Card.Id,
	}
}
