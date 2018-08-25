package engine

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func Trigger(game *vii.Game, past *Timeline, seat *vii.GameSeat, card *vii.GameCard, p *vii.Power, target interface{}) Event {
	game.Log().Info("trigger")

	return &TriggerEvent{
		Stack:  past,
		Card:   card,
		Power:  p,
		Target: target,
	}
}

type TriggerEvent struct {
	Stack *Timeline
	Card  *vii.GameCard
	*vii.Power
	Target interface{}
}

func (event *TriggerEvent) Name() string {
	return "trigger"
}

func (event *TriggerEvent) Priority(game *vii.Game, t *Timeline) bool {
	return t.Reacts[game.GetOpponentSeat(t.HotSeat).Username] != "pass"
}

func (event *TriggerEvent) Receive(game *vii.Game, t *Timeline, seat *vii.GameSeat, json js.Object) {
	game.Log().WithFields(log.Fields{
		"Seat":  seat,
		"Event": json["event"],
	}).Warn("trigger: receive")
}

func (event *TriggerEvent) OnStart(game *vii.Game, t *Timeline) {
	seat := game.GetSeat(t.HotSeat)
	game.Log().WithFields(log.Fields{
		"Username":  seat.Username,
		"Elements":  seat.Elements,
		"Card":      event.Card,
		"PowerId":   event.Power.Id,
		"StackMode": event.Stack.Name(),
	}).Debug("engine-trigger: OnStart")
	animate.BroadcastCardUpdate(game, event.Card)
}

func (event *TriggerEvent) OnReconnect(*vii.Game, *Timeline, *vii.GameSeat) {
}

func (event *TriggerEvent) OnStop(game *vii.Game, t *Timeline) *Timeline {
	seat := game.GetSeat(t.HotSeat)
	game.Log().WithFields(log.Fields{
		"Username":  seat.Username,
		"gcid":      event.Card.Id,
		"CardId":    event.Card.Card.Id,
		"Name":      event.Card.Name,
		"PowerId":   event.Power.Id,
		"UsesTurn":  event.Power.UsesTurn,
		"StackMode": event.Stack.Name(),
	}).Info("trigger resolve")

	Script(game, t, seat, event.Power, event.Target)
	return event.Stack
}

func (event *TriggerEvent) Json(game *vii.Game, t *Timeline) js.Object {
	seat := game.GetSeat(t.HotSeat)
	return js.Object{
		"gameid":   game,
		"timer":    t.Lifetime.Seconds(),
		"username": seat.Username,
		"elements": seat.Elements,
		"hand":     len(seat.Hand),
		"gcid":     event.Card.Id,
	}
}
