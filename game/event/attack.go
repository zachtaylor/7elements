package event

import (
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
	"ztaylor.me/log"
)

func NewAttackEvent(seat string, card *game.Card) game.Event {
	return &AttackEvent{
		Event:      Event(seat),
		AttackCard: card,
	}
}

type AttackEvent struct {
	Event
	AttackCard *game.Card
	DefendCard *game.Card
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

// Finish implements game.FinishEventer
func (event *AttackEvent) Finish(*game.T) []game.Event {
	return []game.Event{NewCombatEvent(event.Seat(), event.AttackCard, event.DefendCard)}
}
func _defendFinishEventer(event *AttackEvent) game.FinishEventer {
	return event
}

// // GetStack implements game.StackEventer
// func (event *AttackEvent) GetStack(g *game.T) *game.State {
// 	return nil
// }

// GetNext implements game.Event
func (event *AttackEvent) GetNext(_ *game.T) game.Event {
	return nil
}

func (event *AttackEvent) JSON() cast.JSON {
	return event.AttackCard.JSON()
}

// Request implements game.RequestEventer
func (event *AttackEvent) Request(g *game.T, seat *game.Seat, json cast.JSON) {
	log := g.Log().With(log.Fields{
		"Seat":  seat.Username,
		"Event": json["event"],
	}).Tag("attack")
	if seat.Username == event.Seat() {
		log.Add("Priority", event.Seat()).Warn("seat mismatch")
	} else if gcid := json.GetS("gcid"); gcid == "" {
		log.Add("GCID", json["gcid"]).Warn("gcid missing")
	} else if c := seat.Present[gcid]; c == nil {
		log.Add("GCID", json["gcid"]).Error("gcid not found")
	} else if !c.IsAwake {
		log.Warn("card asleep")
		seat.Send(game.BuildErrorUpdate(c.Card.Name, "not awake"))
	} else {
		event.DefendCard = c
	}

	g.State.Reacts[seat.Username] = "defend" // trigger len(game.State.Reacts) == len(game.Seats)
}
