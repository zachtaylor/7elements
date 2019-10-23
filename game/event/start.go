package event

import (
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
)

func NewStartEvent(seat string) game.Event {
	return &StartEvent{
		Event: Event(seat),
	}
}

type StartEvent struct {
	Event
}

func (event *StartEvent) Name() string {
	return "start"
}

// OnActivate implements game.ActivateEventer
func (event *StartEvent) OnActivate(g *game.T) []game.Event {
	for _, seat := range g.Seats {
		seat.Life = 7
	}
	for _, seat := range g.Seats {
		seat.Deck.Shuffle()
		seat.DrawCard(3)
		seat.SendHandUpdate()
	}
	return nil
}
func _startIsActivator(event *StartEvent) game.ActivateEventer {
	return event
}

// // OnConnect implements game.ConnectEventer
// func (event *StartEvent) OnConnect(g *game.T, seat *game.Seat) {
// }

// // Finish implements game.FinishEventer
// func (event *StartEvent) Finish(*game.T) {
// }

// // GetStack implements game.StackEventer
// func (event *StartEvent) GetStack(g *game.T) *game.State {
// 	return nil
// }

func (event *StartEvent) GetNext(g *game.T) game.Event {
	return Sunrise(event.Seat())
}

func (event *StartEvent) JSON() cast.JSON {
	return nil
}

// Request implements game.RequestEventer
func (event *StartEvent) Request(g *game.T, seat *game.Seat, json cast.JSON) {
	choice := json.GetS("choice")
	log := g.Log().Add("Seat", seat).Add("Choice", choice).Tag("engine/start")

	if react := g.State.Reacts[seat.Username]; react != "" {
		log.Add("React", react).Warn("already recorded")
		return
	} else if choice == "keep" {
		g.State.Reacts[seat.Username] = "keep"
	} else if choice == "mulligan" {
		g.State.Reacts[seat.Username] = "mulligan"
		seat.DiscardHand()
		seat.DrawCard(3)
		seat.SendHandUpdate()
		g.SendSeatUpdate(seat)
	} else {
		log.Warn("unrecognized")
		return
	}

	g.State.Reacts[seat.Username] = choice
	g.SendReactUpdate(seat.Username)
	log.Info("confirm")
}
