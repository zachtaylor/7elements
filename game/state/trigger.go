package state

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/trigger"
	"github.com/zachtaylor/7elements/power"
	"ztaylor.me/cast"
)

func NewTrigger(seat string, token *game.Token, p *power.T, target interface{}) game.Stater {
	return &Trigger{
		R:      R(seat),
		Token:  token,
		Power:  p,
		Target: target,
	}
}

type Trigger struct {
	R
	Token  *game.Token
	Power  *power.T
	Target interface{}
}

func (r *Trigger) Name() string {
	return "trigger"
}

func (r *Trigger) GetNext(g *game.T) game.Stater {
	return nil
}

func (r *Trigger) JSON() cast.JSON {
	json := cast.JSON{
		"token": r.Token.JSON(),
		"power": r.Power.JSON(),
	}
	if c, ok := r.Target.(*card.T); ok {
		json["target"] = c.JSON()
	} else if t, ok := r.Target.(*game.Token); ok {
		json["target"] = t.JSON()
	} else {
		json["target"] = cast.String(r.Target)
	}
	return json
}

// // OnActivate implements game.ActivateStater
// func (r *Trigger) OnActivate(g *game.T) []game.Stater {
// 	return nil
// }
// func (r *Trigger) activateEventer() game.ActivateStater {
// 	return r
// }

// OnConnect implements game.ConnectStater
func (r *Trigger) OnConnect(g *game.T, seat *game.Seat) {
	if seat == nil {
		go g.GetChat().AddMessage(chat.NewMessage("trigger", r.Token.Card.Proto.Name))
	}
}
func (r *Trigger) _isConnectStater() game.ConnectStater {
	return r
}

// Finish implements game.FinishStater
func (r *Trigger) Finish(g *game.T) []game.Stater {
	seat := g.GetSeat(r.Seat())
	g.Log().With(cast.JSON{
		"Username": seat.Username,
		"Token":    r.Token,
		"Stack":    g.State.Stack,
	}).Debug("engine/trigger: finish")
	return trigger.Power(g, seat, r.Power, r.Token, cast.NewArray(r.Target))
}
func (r *Trigger) _finishStater() game.FinishStater {
	return r
}

// // OnConnect implements game.ConnectStater
// func (r *Trigger) OnConnect(*game.T, *game.Seat) {
// }

// // Request implements game.RequestStater
// func (r *Trigger) Request(g*game.T, seat *game.Seat, json js.Object) {
// }
