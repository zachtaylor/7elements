package state

import (
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/trigger"
	"github.com/zachtaylor/7elements/out"
	"ztaylor.me/cast"
)

type Combat struct {
	R
	A *game.Token
	B *game.Token
}

func NewCombat(seat string, t1 *game.Token, t2 *game.Token) game.Stater {
	return &Combat{
		R: R(seat),
		A: t1,
		B: t2,
	}
}

func (r *Combat) Name() string {
	return "combat"
}

// OnActivate implements game.ActivateStater
func (r *Combat) OnActivate(g *game.T) []game.Stater {
	if r.B != nil {
		go g.Settings.Chat.AddMessage(chat.NewMessage(r.A.Card.Proto.Name, "vs "+r.B.Card.Proto.Name))
	} else if enemyseat := g.GetOpponentSeat(r.A.Username); enemyseat == nil {
	} else {
		go g.Settings.Chat.AddMessage(chat.NewMessage(r.A.Card.Proto.Name, "vs "+enemyseat.Username))
	}
	return nil
}
func (r *Combat) _isActivateEventer() game.ActivateStater {
	return r
}

// // OnConnect implements game.ConnectStater
// func (r *Combat) OnConnect(*game.T, *game.Seat) {
// }

// // GetStack implements game.StackEventer
// func (r *Combat) GetStack(g *game.T) *game.State {
// 	return nil
// }

// // Request implements game.RequestStater
// func (r *Combat) Request(g*game.T, seat *game.Seat, json cast.JSON) {
// }

// Finish implements game.FinishStater
func (r *Combat) Finish(g *game.T) []game.Stater {
	var rs []game.Stater
	if r.B != nil {
		if e := trigger.Damage(g, r.B, r.A.Body.Attack); e != nil {
			rs = append(rs, e...)
		}
		out.GameToken(g, r.B.JSON())
		if e := trigger.Damage(g, r.A, r.B.Body.Attack); e != nil {
			rs = append(rs, e...)
		}
		out.GameToken(g, r.A.JSON())
	} else if enemyseat := g.GetOpponentSeat(r.A.Username); enemyseat == nil {

	} else if dmgEvents := trigger.DamageSeat(g, r.A.Card, enemyseat, r.A.Body.Attack); len(dmgEvents) > 0 {
		rs = append(rs, dmgEvents...)
	}

	return rs
}
func (r *Combat) _finishStater() game.FinishStater {
	return r
}

func (r *Combat) GetNext(g *game.T) game.Stater {
	return nil
}

func (r *Combat) JSON() cast.JSON {
	return cast.JSON{
		"attack": r.A.JSON(),
		"block":  r.B.JSON(),
	}
}
