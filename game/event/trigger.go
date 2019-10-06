package event

import (
	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
	"ztaylor.me/log"
)

func NewTriggerEvent(seat string, card *game.Card, p *vii.Power, target interface{}) game.Event {
	return &TriggerEvent{
		Event:  Event(seat),
		Card:   card,
		Power:  p,
		Target: target,
	}
}

type TriggerEvent struct {
	Event
	Card   *game.Card
	Power  *vii.Power
	Target interface{}
}

func (event *TriggerEvent) Name() string {
	return "trigger"
}

func (event *TriggerEvent) GetNext(g *game.T) game.Event {
	return nil
}

func (event *TriggerEvent) JSON() cast.JSON {
	json := cast.JSON{
		"card":  event.Card.JSON(),
		"power": event.Power.JSON(),
	}
	if c, ok := event.Target.(*game.Card); ok {
		json["target"] = c.JSON()
	} else {
		json["target"] = event.Target
	}
	return json
}

// OnActivate implements game.ActivateEventer
func (event *TriggerEvent) OnActivate(g *game.T) []game.Event {
	go g.GetChat().AddMessage(chat.NewMessage(event.Seat(), "trigger: "+event.Card.Card.Name))
	return nil
}
func (event *TriggerEvent) activateEventer() game.ActivateEventer {
	return event
}

// Finish implements game.FinishEventer
func (event *TriggerEvent) Finish(g *game.T) []game.Event {
	seat := g.GetSeat(event.Seat())
	log := g.Log().With(log.Fields{
		"Username": seat.Username,
		"Card":     event.Card.Print(),
		"PowerId":  event.Power.Id,
		"UsesTurn": event.Power.UsesTurn,
		"Stack":    g.State.Stack,
	}).Tag("engine/trigger")

	if script := game.Scripts[event.Power.Script]; script == nil {
		log.Add("Script", event.Power.Script).Warn("script missing")
	} else {
		log.Info("scripting")
		return script(g, seat, event.Target)
	}
	return nil
}
func (event *TriggerEvent) finishEventer() game.FinishEventer {
	return event
}

// // OnConnect implements game.ConnectEventer
// func (event *TriggerEvent) OnConnect(*game.T, *game.Seat) {
// }

// // Request implements game.RequestEventer
// func (event *TriggerEvent) Request(g*game.T, seat *game.Seat, json js.Object) {
// }
