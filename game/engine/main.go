package engine

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/game"
)

func Main(game *game.T) game.Event {
	return new(MainEvent)
}

type MainEvent bool

func (event *MainEvent) Name() string {
	return "main"
}

// // OnActivate implements game.ActivateEventer
// func (event *MainEvent) OnActivate(game *game.T) {
// }

// // OnConnect implements game.ConnectEventer
// func (event *MainEvent) OnConnect(*game.T, *game.Seat) {
// }

// // Finish implements game.FinishEventer
// func (event *MainEvent) Finish(*game.T) {
// }

// // Request implements game.RequestEventer
// func (event *MainEvent) Request(game *game.T, seat *game.Seat, json js.Object) {
// }

// // GetStack implements game.StackEventer
// func (event *MainEvent) GetStack(*game.T) *game.State {
// 	return nil
// }

func (event *MainEvent) GetNext(game *game.T) *game.State {
	return game.NewState(game.State.Seat, Attack(game))
}

func (event *MainEvent) Json(game *game.T) vii.Json {
	return nil
}
