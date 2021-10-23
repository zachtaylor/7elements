package engine

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/engine/script"
	"github.com/zachtaylor/7elements/game/engine/trigger"
	"github.com/zachtaylor/7elements/game/phase"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/game/token"
	"github.com/zachtaylor/7elements/power"
	"taylz.io/log"
)

// T is a game engine
type T struct{}

// New returns a new game engine
func New() game.Engine { return &T{} }

func (t *T) Run(logger *log.T, game *game.T) {
	t.run(logger, game)
}

func (t *T) NewState(game *game.T, phase game.Phaser) *game.State {
	state := game.NewState(game.Rules().Timeout, phase)
	game.RegisterState(state)
	return state
}

func (t *T) NewEnding(game *game.T, results game.Resulter) game.Phaser {
	return phase.NewEnd(results)
}

func (t *T) NewTrigger(game *game.T, seat *seat.T, token *token.T, power *power.T) game.Phaser {
	if power.Target == "self" {
		return phase.NewTrigger(seat.Username, token, power, token.ID)
	}
	return phase.NewTarget(seat.Username, power, token)
}

func (t *T) NewToken(game *game.T, seat *seat.T, token *token.T) (rs []game.Phaser) {
	rs = trigger.NewToken(game, seat, token)
	return
}

func (t *T) RemoveToken(game *game.T, token *token.T) (rs []game.Phaser) {
	return trigger.RemoveToken(game, token)
}

func (t *T) WakeToken(game *game.T, token *token.T) (rs []game.Phaser) {
	return trigger.WakeToken(game, token)
}

func (t *T) SleepToken(game *game.T, token *token.T) (rs []game.Phaser) {
	return trigger.SleepToken(game, token)
}

func (t *T) HealToken(game *game.T, token *token.T, n int) []game.Phaser {
	return trigger.HealToken(game, token, n)
}

func (t *T) DamageToken(game *game.T, token *token.T, n int) []game.Phaser {
	return trigger.DamageToken(game, token, n)
}

func (t *T) HealSeat(game *game.T, seat *seat.T, n int) []game.Phaser {
	return trigger.HealSeat(game, seat, n)
}

func (t *T) DamageSeat(game *game.T, seat *seat.T, n int) (rs []game.Phaser) {
	return trigger.DamageSeat(game, seat, n)
}

func (t *T) DrawCard(game *game.T, seat *seat.T, n int) []game.Phaser {
	return trigger.DrawCard(game, seat, n)
}

func (t *T) Script(game *game.T, seat *seat.T, scriptName string, me interface{}, args []string) []game.Phaser {
	return script.Run(game, seat, scriptName, me, args)
}

// func (t *T) Target(seat *game.Seat, target string, text string, finish func(val string) []game.Phaser) game.Phaser {
// 	return state.NewTarget(seat.Username, target, text, finish)
// }

// func (t *T) TriggerTokenEvent(g *game.T, seat *seat.T, token *token.T, name string) (rs []game.Phaser) {
// 	powers := token.Powers.GetTrigger(name)
// 	if len(powers) < 1 {
// 		return
// 	}
// 	for _, p := range powers {
// 		if p.Target != "self" {
// 			power := p
// 			rs = append(rs, phase.NewTarget(
// 				seat.Username,
// 				power,
// 				token,
// 			))
// 		} else {
// 			rs = append(rs, state.NewTrigger(seat.Username, token, p, token))
// 		}
// 	}
// 	return
// }

// func (t *T) TriggerTokenPower(g *game.T, seat *game.Seat, token *game.Token, power *power.T, arg interface{}) []game.Phaser {
// 	dirty := false
// 	if power.Costs.Total() > 0 {
// 		dirty = true
// 		seat.Karma.Deactivate(power.Costs)
// 	}
// 	if power.UsesTurn {
// 		token.IsAwake = false
// 		out.GameToken(g, token.JSON())
// 	}
// 	if power.UsesKill {
// 		dirty = true
// 		delete(seat.Present, token.ID)
// 	}
// 	if dirty {
// 		out.GameSeat(g, seat.JSON())
// 	}

// 	if power.Target == "self" {
// 		return []game.Phaser{state.NewTrigger(seat.Username, token, power, token)}
// 	}
// 	return []game.Phaser{state.NewTrigger(seat.Username, token, power, arg)}
// }

// func (*T) Target(seat *seat.T, target string, text string, finish func(val string) []state.Phaser) state.Phaser {
// 	return phase.NewTarget(seat.Username, target, text, finish)
// }
