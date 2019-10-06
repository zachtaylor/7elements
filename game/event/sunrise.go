package event

import (
	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
	"ztaylor.me/log"
)

func Sunrise(seat string) game.Event {
	return &SunriseEvent{
		Event: Event(seat),
	}
}

type SunriseEvent struct {
	Event
	element vii.Element
}

func (event *SunriseEvent) Name() string {
	return "sunrise"
}

// OnActivate implements game.ActivateEventer
func (event *SunriseEvent) OnActivate(g *game.T) []game.Event {
	seat := g.GetSeat(event.Seat())
	seat.Reactivate()
	g.SendAll(game.BuildSeatUpdate(seat))
	return game.SeatTriggeredEvents(g, g.GetSeat(event.Seat()), "sunrise")
}
func _sunriseIsActivator(event *SunriseEvent) game.ActivateEventer {
	return event
}

// // OnConnect implements game.ConnectEventer
// func (event *SunriseEvent) OnConnect(*game.T, *game.Seat) {
// }

// // GetStack implements game.StackEventer
// func (event *SunriseEvent) GetStack(g *game.T) *game.State {
// 	return nil
// }

// Finish implements game.FinishEventer
func (event *SunriseEvent) Finish(g *game.T) []game.Event {
	seat := g.GetSeat(event.Seat())
	seat.Elements.Append(event.element)
	if card := seat.Deck.Draw(); card != nil {
		seat.Hand[card.Id] = card
		seat.Send(game.BuildHandUpdate(seat))
	}
	g.SendAll(game.BuildSeatUpdate(seat))
	g.Log().With(log.Fields{
		"Element":  event.element,
		"Karma":    seat.Elements.String(),
		"Username": seat.Username,
		"Hand":     seat.Hand.Print(),
	}).Info("engine/sunrise: finish")
	return nil
}

func (event *SunriseEvent) GetNext(g *game.T) game.Event {
	return Main(event.Seat())
}

func (event *SunriseEvent) JSON() cast.JSON {
	return nil
}

// Request implements game.RequestEventer
func (event *SunriseEvent) Request(g *game.T, seat *game.Seat, json cast.JSON) {
	log := g.Log().With(log.Fields{
		"Seat": seat.Username,
	}).Tag("engine/sunrise")
	if event.Seat() != seat.Username {
		log.Warn("wrong person")
		return
	}
	elementID := json.GetI("elementid")
	if elementID < 1 || elementID > 7 {
		log.Add("ElementId", elementID).Warn("elementid out of bounds")
		return
	}
	event.element = vii.Elements[int(elementID)]
	log.Add("Element", event.element).Info("confirm")
	g.State.Reacts[seat.Username] = "confirm"
	g.SendAll(game.BuildReactUpdate(g, seat.Username))
}
