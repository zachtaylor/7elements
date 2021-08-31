package phase

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/game/token"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

func NewAttack(seat string, token *token.T) game.Phaser {
	return &Attack{
		R: R(seat),
		A: token,
	}
}

type Attack struct {
	R
	A *token.T
	B *token.T
}

func (r *Attack) Name() string {
	return "attack"
}

func (r *Attack) String() string {
	return "attack (" + r.A.ID + ")"
}

// OnActivate implements game.OnActivatePhaser
func (r *Attack) OnActivate(game *game.T) []game.Phaser {
	go game.Chat(r.A.Card.Proto.Name, "attack")
	return nil
}
func (e *Attack) onActivatePhaser() game.OnActivatePhaser { return e }

// // OnConnect implements game.OnConnectPhaser
// func (r *Attack) OnConnect(*game.T, *seat.T) {
// }

// Finish implements game.OnFinishPhaser
func (r *Attack) OnFinish(*game.T) []game.Phaser {
	return []game.Phaser{NewCombat(r.Seat(), r.A, r.B)}
}
func (r *Attack) onFinishPhaser() game.OnFinishPhaser { return r }

// // GetStack implements game.StackRer
// func (r *Attack) GetStack(game *game.T) *game.State {
// 	return nil
// }

// GetNext implements game.Phaser
func (r *Attack) GetNext(*game.T) game.Phaser {
	return nil
}

func (r *Attack) Data() websocket.MsgData { return r.A.Data() }

// Request implements Requester
func (r *Attack) OnRequest(game *game.T, seat *seat.T, json websocket.MsgData) {
	log := game.Log().Add("Seat", seat.Username)
	if seat.Username == r.Seat() {
		log.Add("Priority", r.Seat()).Warn("seat mismatch")
	} else if id, _ := json["id"].(string); id == "" {
		log.Add("ID", json["id"]).Warn("id missing")
	} else if t := seat.Present[id]; t == nil {
		log.Add("ID", id).Error("id not found")
	} else if !t.IsAwake {
		log.Warn("token asleep")
		seat.Writer.Write(wsout.ErrorJSON(t.Card.Proto.Name, "not awake"))
	} else {
		r.B = t
	}

	game.State.Reacts[seat.Username] = "defend" // trigger len(game.State.Reacts) == len(game.Seats)
}
func (r *Attack) onRequestPhaser() game.OnRequestPhaser { return r }
