package event

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/trigger"
	"ztaylor.me/cast"
)

func NewSunsetEvent(seat string) game.Event {
	return &SunsetEvent{
		Event: Event(seat),
	}
}

type SunsetEvent struct {
	Event
}

func (event *SunsetEvent) Name() string {
	return "sunset"
}

// OnActivate implements game.ActivateEventer
func (event *SunsetEvent) OnActivate(g *game.T) []game.Event {
	seat := g.GetSeat(event.Seat())
	return trigger.SeatPresent(g, seat, "sunset")
}
func _sunsetIsActivateEventer(event *SunsetEvent) game.ActivateEventer {
	return event
}

// // OnConnect implements game.ConnectEventer
// func (event *SunsetEvent) OnConnect(*game.T, *game.Seat) {
// }

// // // Finish implements game.FinishEventer
// func (event *SunsetEvent) Finish(g *game.T) {
// 	// game.State.Seat = game.GetOpponentSeat(game.State.Seat).Username
// }

// // GetStack implements game.StackEventer
// func (event *SunsetEvent) GetStack(g *game.T) *game.State {
// 	return nil
// }

// // Request implements game.RequestEventer
// func (event *SunsetEvent) Request(g*game.T, seat *game.Seat, json cast.JSON) {
// }

func (event *SunsetEvent) GetNext(g *game.T) game.Event {
	return Sunrise(g.GetOpponentSeat(event.Seat()).Username)
}

func (event *SunsetEvent) JSON() cast.JSON {
	return nil
}
