package state

import (
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
)

func NewMain(seat string) game.Stater {
	return &Main{
		R: R(seat),
	}
}

type Main struct {
	R
}

func (r *Main) Name() string {
	return "main"
}

// // OnActivate implements game.ActivateStater
// func (r *Main) OnActivate(g *game.T) []game.Stater {
// 	return nil
// }

// OnConnect implements game.ConnectStater
func (r *Main) OnConnect(g *game.T, seat *game.Seat) {
	if seat == nil {
		go g.GetChat().AddMessage(chat.NewMessage("main", r.Seat()))
	}
}
func (r *Main) _isConnectStater() game.ConnectStater {
	return r
}

// // Finish implements game.FinishStater
// func (r *Main) Finish(*game.T) {
// }

// // Request implements game.RequestStater
// func (r *Main) Request(g*game.T, seat *game.Seat, json js.Object) {
// }

// // GetStack implements game.StackEventer
// func (r *Main) GetStack(*game.T) *game.State {
// 	return nil
// }

func (r *Main) GetNext(g *game.T) game.Stater {
	return NewSunset(r.Seat())
}

func (r *Main) JSON() cast.JSON {
	return nil
}
