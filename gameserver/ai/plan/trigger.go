package plan

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/game/token"
	"github.com/zachtaylor/7elements/gameserver/ai/aim"
	"github.com/zachtaylor/7elements/power"
	"taylz.io/http/websocket"
)

// Trigger is a plan to trigger an ability
type Trigger struct {
	TokenID string
	PowerID int
	Target  interface{}
	score   int
}

func (trigger *Trigger) Score() int {
	return trigger.score
}

func (trigger *Trigger) Submit(request RequestFunc) {
	request("trigger", websocket.MsgData{
		"id":      trigger.TokenID,
		"powerid": float64(trigger.PowerID),
		"target":  trigger.Target,
	})
}

func (trigger *Trigger) String() string {
	return "Trigger " + trigger.TokenID
}

func ParseTrigger(game *game.T, seat *seat.T) (trigger *Trigger) {
	for _, token := range seat.Present {
		if t := ParseTriggerWith(game, seat, token); t == nil {
		} else if trigger == nil || trigger.score < t.score {
			trigger = t
		}
	}
	return
}

func ParseTriggerWith(game *game.T, seat *seat.T, token *token.T) (trigger *Trigger) {
	powers := token.Powers.GetTrigger("")
	if len(powers) < 1 {
		return
	}
	for _, power := range powers {
		if !seat.Karma.Active().Test(power.Costs) {
		} else if power.UsesTurn && !token.IsAwake {
		} else if t := ParseTriggerWithWith(game, seat, token, power); t == nil {
		} else if trigger == nil || trigger.score < t.score {
			trigger = t
		}
	}
	return
}

func ParseTriggerWithWith(game *game.T, seat *seat.T, token *token.T, power *power.T) (trigger *Trigger) {
	if target, score := triggerPowerScore(game, seat, token, power); score > 0 {
		trigger = &Trigger{
			TokenID: token.ID,
			PowerID: power.ID,
			Target:  target,
			score:   score,
		}
	}
	return
}

// triggerPowerScore picks best Target for given Token and Power, and returns score for the choice
func triggerPowerScore(game *game.T, seat *seat.T, t *token.T, p *power.T) (target interface{}, score int) {
	switch t.Card.Proto.ID {
	case 1:
		target, score = t.ID, 10
	case 2:
		// if ai.Settings.Aggro {
		// target, score = game.Seats.GetOpponent(seat.Username).Username, 10
		// } else if target, score = ai.TargetEnemyBeing("damage"); score < 1 {
		target, score = game.Seats.GetOpponent(seat.Username).Username, 1
		// }
	case 5:
		// if ai.Settings.Aggro {
		if game.Phase() != `sunrise` || game.State.Phase.Seat() != seat.Username {
		} else {
			target, score = aim.EnemyBeing(game, seat, "sleep")
		}
		// } else {
		// if ai.Game.State.Phase.Name() != `main` || !ai.myturn() {
		// } else {
		// target, score = ai.TargetEnemyBeing("sleep")
		// }
		// }
	case 7:
		target, score = aim.EnemyBeing(game, seat, "damage")
	case 15:
		target, score = aim.EnemyPastBeingItem(game, seat)
	case 20:
		target, score = seat.Username, 8-seat.Life
		if targetT, scoreT := aim.MyPresentBeing(game, seat, "health"); scoreT > score {
			target, score = targetT, scoreT
		}
	case 21:
		target, score = aim.MyPastBeing(game, seat)
	}
	return
}
