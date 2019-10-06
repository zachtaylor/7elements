package event

import (
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
)

func Main(seat string) game.Event {
	return &MainEvent{
		Event: Event(seat),
	}
}

type MainEvent struct {
	Event
}

func (event *MainEvent) Name() string {
	return "main"
}

// // OnActivate implements game.ActivateEventer
// func (event *MainEvent) OnActivate(g *game.T) {
// }

// // OnConnect implements game.ConnectEventer
// func (event *MainEvent) OnConnect(*game.T, *game.Seat) {
// }

// // Finish implements game.FinishEventer
// func (event *MainEvent) Finish(*game.T) {
// }

// // Request implements game.RequestEventer
// func (event *MainEvent) Request(g*game.T, seat *game.Seat, json js.Object) {
// }

// // GetStack implements game.StackEventer
// func (event *MainEvent) GetStack(*game.T) *game.State {
// 	return nil
// }

func (event *MainEvent) GetNext(g *game.T) game.Event {
	return NewSunsetEvent(event.Seat())
}

func (event *MainEvent) JSON() cast.JSON {
	return nil
}
