package phase

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/engine/script"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/game/token"
	"github.com/zachtaylor/7elements/power"
)

func NewTrigger(seat string, token *token.T, p *power.T, target string) game.Phaser {
	return &Trigger{
		R:      R(seat),
		Token:  token,
		Power:  p,
		Target: target,
	}
}

type Trigger struct {
	R
	Token  *token.T
	Power  *power.T
	Target string
}

func (r *Trigger) Name() string {
	return "trigger"
}

func (r *Trigger) String() string {
	return "trigger (" + r.Seat() + ":" + r.Token.ID + ")"
}

func (r *Trigger) GetNext(game *game.T) game.Phaser {
	return nil
}

func (r *Trigger) Data() map[string]interface{} {
	return map[string]interface{}{
		"token":  r.Token.Data(),
		"power":  r.Power.Data(),
		"target": r.Target,
	}
}

// // OnActivate implements game.OnActivatePhaser
// func (r *Trigger) OnActivate(game *T) []Phaser {
// 	return nil
// }
// func (r *Trigger) activateEventer() game.OnActivatePhaser {
// 	return r
// }

// OnConnect implements OnConnectPhaser
func (r *Trigger) OnConnect(game *game.T, seat *seat.T) {
	if seat == nil {
		go game.Chat(game.State.Phase.Seat(), "Trigger "+r.Token.Card.Proto.Name)
	}
}
func (r *Trigger) onConnectPhaser() game.OnConnectPhaser { return r }

// Finish implements OnFinishPhaser
func (r *Trigger) OnFinish(game *game.T) []game.Phaser {
	seat := game.Seats.Get(r.Seat())
	game.Log().With(map[string]interface{}{
		"Username": seat.Username,
		"Token":    r.Token,
		"Stack":    game.State.Stack,
	}).Debug("engine/trigger: finish")
	return script.Run(game, seat, r.Power, r.Token, []string{r.Target})
}
func (r *Trigger) onFinishPhaser() game.OnFinishPhaser { return r }

// // OnConnect implements OnConnectPhaser
// func (r *Trigger) OnConnect(*T, *seat.T) {
// }

// // Request implements Requester
// func (r *Trigger) Request(g*T, seat *seat.T, json js.Object) {
// }
