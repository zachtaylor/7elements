package phase

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/token"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

type Combat struct {
	R
	A *token.T
	B *token.T
}

func NewCombat(seat string, t1 *token.T, t2 *token.T) game.Phaser {
	return &Combat{
		R: R(seat),
		A: t1,
		B: t2,
	}
}

func (r *Combat) Name() string {
	return "combat"
}

func (r *Combat) String() string {
	return "combat (" + r.A.String() + "v" + r.B.String() + ")"
}

// OnActivate implements game.OnActivatePhaser
func (r *Combat) OnActivate(game *game.T) []game.Phaser {
	if r.B != nil {
		go game.Chat(r.A.Card.Proto.Name, "vs "+r.B.Card.Proto.Name)
	} else if enemyseat := game.Seats.GetOpponent(r.A.User); enemyseat == nil {
	} else {
		go game.Chat(r.A.Card.Proto.Name, "vs "+enemyseat.Username)
	}
	return nil
}
func (r *Combat) onActivatePhaser() game.OnActivatePhaser { return r }

// // OnConnect implements game.OnConnectPhaser
// func (r *Combat) OnConnect(*game.T, *seat.T) {
// }

// // GetStack implements game.StackEventer
// func (r *Combat) GetStack(game *game.T) *game.State {
// 	return nil
// }

// // Request implements Requester
// func (r *Combat) Request(g*game.T, seat *seat.T, json websocket.MsgData) {
// }

// Finish implements game.OnFinishPhaser
func (r *Combat) OnFinish(game *game.T) (rs []game.Phaser) {
	if r.B != nil {
		if e := game.Engine().DamageToken(game, r.B, r.A.Body.Attack); e != nil {
			rs = append(rs, e...)
		}
		game.Seats.Write(wsout.GameTokenJSON(r.B.Data()))
		if e := game.Engine().DamageToken(game, r.A, r.B.Body.Attack); e != nil {
			rs = append(rs, e...)
		}
		game.Seats.Write(wsout.GameTokenJSON(r.A.Data()))
	} else if enemyseat := game.Seats.GetOpponent(r.A.User); enemyseat == nil {

	} else if e := game.Engine().DamageSeat(game, enemyseat, r.A.Body.Attack); len(e) > 0 {
		rs = append(rs, e...)
	}
	return
}
func (r *Combat) onFinishPhaser() game.OnFinishPhaser { return r }

func (r *Combat) GetNext(game *game.T) game.Phaser {
	return nil
}

func (r *Combat) Data() websocket.MsgData {
	return websocket.MsgData{
		"attack": r.A.Data(),
		"block":  r.B.Data(),
	}
}
