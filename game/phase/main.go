package phase

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/seat"
	"taylz.io/http/websocket"
)

func NewMain(seat string) game.Phaser {
	return &Main{
		R: R(seat),
	}
}

type Main struct{ R }

func (r *Main) Name() string { return "main" }

func (r *Main) String() string { return "main (" + r.Seat() + ")" }

// OnConnect implements game.OnConnectPhaser
func (r *Main) OnConnect(game *game.T, seat *seat.T) {
	if seat == nil {
		// go game.Chat("main", r.Seat())
	}
}
func (r *Main) onConnectPhaser() game.OnConnectPhaser { return r }

func (r *Main) GetNext(game *game.T) game.Phaser { return NewSunset(r.Seat()) }

func (r *Main) Data() websocket.MsgData { return nil }
