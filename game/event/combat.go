package event

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/trigger"
	"ztaylor.me/cast"
)

type CombatEvent struct {
	Event
	A *game.Card
	B *game.Card
}

func NewCombatEvent(seat string, acard *game.Card, dcard *game.Card) game.Event {
	return &CombatEvent{
		Event: Event(seat),
		A:     acard,
		B:     dcard,
	}
}

func (event *CombatEvent) Name() string {
	return "combat"
}

// // OnActivate implements game.ActivateEventer
// func (event *CombatEvent) OnActivate(*game.T) {
// }

// // OnConnect implements game.ConnectEventer
// func (event *CombatEvent) OnConnect(*game.T, *game.Seat) {
// }

// // GetStack implements game.StackEventer
// func (event *CombatEvent) GetStack(g *game.T) *game.State {
// 	return nil
// }

// // Request implements game.RequestEventer
// func (event *CombatEvent) Request(g*game.T, seat *game.Seat, json cast.JSON) {
// }

// Finish implements game.FinishEventer
func (event *CombatEvent) Finish(g *game.T) []game.Event {
	var events []game.Event
	if event.B != nil {
		if e := trigger.Damage(g, event.B, event.A.Body.Attack); e != nil {
			events = append(events, e...)
		}
		g.SendCardUpdate(event.B)
		if e := trigger.Damage(g, event.A, event.B.Body.Attack); e != nil {
			events = append(events, e...)
		}
		g.SendCardUpdate(event.A)
	} else if enemyseat := g.GetOpponentSeat(event.A.Username); enemyseat == nil {

	} else if dmgEvents := trigger.DamageSeat(g, event.A, enemyseat, event.A.Body.Attack); len(dmgEvents) > 0 {
		events = append(events, dmgEvents...)
	}

	return events
}
func (event *CombatEvent) finishEventer() game.FinishEventer {
	return event
}

func (event *CombatEvent) GetNext(g *game.T) game.Event {
	return nil
}

func (event *CombatEvent) JSON() cast.JSON {
	return cast.JSON{
		"attack": event.A.JSON(),
		"block":  event.B.JSON(),
	}
}
