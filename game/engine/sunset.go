package engine

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/game"
)

func Sunset(game *game.T) game.Event {
	game.Log().Info("sunset")
	return new(SunsetEvent)
}

type SunsetEvent struct {
}

func (event *SunsetEvent) Name() string {
	return "sunset"
}

// // OnActivate implements game.ActivateEventer
// func (event *SunsetEvent) OnActivate(game *game.T) {
// }

// // OnConnect implements game.ConnectEventer
// func (event *SunsetEvent) OnConnect(*game.T, *game.Seat) {
// }

// // // Finish implements game.FinishEventer
// func (event *SunsetEvent) Finish(game *game.T) {
// 	// game.State.Seat = game.GetOpponentSeat(game.State.Seat).Username
// }

// // GetStack implements game.StackEventer
// func (event *SunsetEvent) GetStack(game *game.T) *game.State {
// 	return nil
// }

// // Request implements game.RequestEventer
// func (event *SunsetEvent) Request(game *game.T, seat *game.Seat, json vii.Json) {
// }

func (event *SunsetEvent) GetNext(game *game.T) *game.State {
	return game.NewState(game.GetOpponentSeat(game.State.Seat).Username, Sunrise(game))
}

func (event *SunsetEvent) Json(game *game.T) vii.Json {
	return nil
}
