package state

import (
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/out"
	"ztaylor.me/cast"
)

func NewAttack(seat string, token *game.Token) game.Stater {
	return &Attack{
		R: R(seat),
		A: token,
	}
}

type Attack struct {
	R
	A *game.Token
	B *game.Token
}

func (r *Attack) Name() string {
	return "attack"
}

// OnActivate implements game.ActivateStater
func (r *Attack) OnActivate(g *game.T) []game.Stater {

	go g.Settings.Chat.AddMessage(chat.NewMessage(r.A.Card.Proto.Name, "attack"))
	return nil
}
func (e *Attack) _isActivateRer() game.ActivateStater {
	return e
}

// // OnConnect implements game.ConnectStater
// func (r *Attack) OnConnect(*game.T, *game.Seat) {
// }

// Finish implements game.FinishStater
func (r *Attack) Finish(*game.T) []game.Stater {
	return []game.Stater{NewCombat(r.Seat(), r.A, r.B)}
}
func (r *Attack) _finishStater() game.FinishStater {
	return r
}

// // GetStack implements game.StackRer
// func (r *Attack) GetStack(g *game.T) *game.State {
// 	return nil
// }

// GetNext implements game.Stater
func (r *Attack) GetNext(_ *game.T) game.Stater {
	return nil
}

func (r *Attack) JSON() cast.JSON {
	return r.A.JSON()
}

// Request implements game.RequestStater
func (r *Attack) Request(g *game.T, seat *game.Seat, json cast.JSON) {
	log := g.Log().Add("Seat", seat.Username)
	if seat.Username == r.Seat() {
		log.Add("Priority", r.Seat()).Warn("seat mismatch")
	} else if id := json.GetS("id"); id == "" {
		log.Add("ID", json["id"]).Warn("id missing")
	} else if t := seat.Present[id]; t == nil {
		log.Add("ID", id).Error("id not found")
	} else if !t.IsAwake {
		log.Warn("token asleep")
		out.GameError(seat.Player, t.Card.Proto.Name, "not awake")
	} else {
		r.B = t
	}

	g.State.Reacts[seat.Username] = "defend" // trigger len(game.State.Reacts) == len(game.Seats)
}
