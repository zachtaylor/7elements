package plan

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/ai/aim"
	"github.com/zachtaylor/7elements/game/ai/view"
	"github.com/zachtaylor/7elements/power"
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
	request("trigger", map[string]any{
		"id":      trigger.TokenID,
		"powerid": float64(trigger.PowerID),
		"target":  trigger.Target,
	})
}

func (trigger *Trigger) String() string {
	return "Trigger " + trigger.TokenID
}

func ParseTrigger(view view.T) (trigger *Trigger) {
	for tokenID := range view.Self.T.Present {
		token := view.Game.Token(tokenID)
		if t := ParseTriggerWith(view, token); t == nil {
		} else if trigger == nil || trigger.score < t.score {
			trigger = t
		}
	}
	return
}

func ParseTriggerWith(view view.T, token *game.Token) (trigger *Trigger) {
	powers := token.T.Powers.GetTrigger("")
	if len(powers) < 1 {
		return
	}
	for _, power := range powers {
		if !view.Self.T.Karma.Active().Test(power.Costs) {
		} else if power.UsesTurn && !token.T.Awake {
		} else if t := ParseTriggerWithWith(view, token, power); t == nil {
		} else if trigger == nil || trigger.score < t.score {
			trigger = t
		}
	}
	return
}

func ParseTriggerWithWith(view view.T, token *game.Token, power *power.T) (trigger *Trigger) {
	if target, score := triggerPowerScore(view, token, power); score > 0 {
		trigger = &Trigger{
			TokenID: token.ID(),
			PowerID: power.ID,
			Target:  target,
			score:   score,
		}
	}
	return
}

// triggerPowerScore picks best Target for given Token and Power, and returns score for the choice
func triggerPowerScore(view view.T, t *game.Token, p *power.T) (target interface{}, score int) {
	castingCard := view.Game.Card(t.T.Card)
	switch castingCard.T.ID {
	case 1:
		target, score = t.ID, 10
	case 2:
		// if ai.Settings.Aggro {
		// target, score = g.PlayerOpponent(seat.Username).Username, 10
		// } else if target, score = ai.TargetEnemyBeing("damage"); score < 1 {
		target, score = view.Enemy.ID(), 1
		// }
	case 5:
		// if ai.Settings.Aggro {
		if !view.IsMyMain {
		} else {
			target, score = aim.EnemyBeing(view, "sleep")
		}
		// } else {
		// if ai.Game.State.Phase.Name() != `main` || !ai.myturn() {
		// } else {
		// target, score = ai.TargetEnemyBeing("sleep")
		// }
		// }
	case 7:
		target, score = aim.EnemyBeing(view, "damage")
	case 15:
		target, score = aim.EnemyPastBeingItem(view)
	case 20:
		target, score = view.Self.ID(), 8-view.Self.T.Life
		if targetT, scoreT := aim.MyPresentBeing(view, "health"); scoreT > score {
			target, score = targetT, scoreT
		}
	case 21:
		target, score = aim.MyPastBeing(view)
	}
	return
}
