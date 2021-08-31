package phase

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/seat"
	"taylz.io/http/websocket"
)

func NewSunset(seat string) game.Phaser {
	return &Sunset{
		R: R(seat),
	}
}

type Sunset struct {
	R
}

func (r *Sunset) Name() string {
	return "sunset"
}

func (r *Sunset) String() string {
	return "sunset (" + r.Seat() + ")"
}

// OnConnect implements game.OnConnectPhaser
func (r *Sunset) OnConnect(game *game.T, seat *seat.T) {
	if seat == nil {
		go game.Chat("sunset", r.Seat())
	}
}
func (r *Sunset) onConnectPhaser() game.OnConnectPhaser { return r }

// // // Finish implements game.OnFinishPhaser
// func (r *Sunset) Finish(game *game.T) {
// 	// game.State.Seat = game.GetOpponentSeat(game.State.Seat).Username
// }

// // GetStack implements game.StackEventer
// func (r *Sunset) GetStack(game *game.T) *game.State {
// 	return nil
// }

// // Request implements Requester
// func (r *Sunset) Request(g*game.T, seat *seat.T, json websocket.MsgData) {
// }

func (r *Sunset) GetNext(game *game.T) game.Phaser {
	return NewSunrise(game.Seats.GetOpponent(r.Seat()).Username)
}

func (r *Sunset) Data() websocket.MsgData { return nil }
