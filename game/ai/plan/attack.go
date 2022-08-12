package plan

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/ai/inspection"
	"github.com/zachtaylor/7elements/game/ai/view"
)

// Attack is a plan to make an attack
type Attack struct {
	id    string // token id
	score int    // plan value
}

func (attack *Attack) Score() int {
	return attack.score
}

func (attack *Attack) Submit(request RequestFunc) {
	request("attack", map[string]any{
		"id": attack.id,
	})
}

func (attack *Attack) String() string {
	return "Attack " + attack.id
}

func ParseAttack(v view.T) (attack *Attack) {
	if !v.IsMyMain {
		v.Game.Log().Add("Phase", v.State.T.Phase.Type()).Add("Seat", v.State.Player()).Trace("skip")
		return nil
	}

	insme := inspection.Parse(v.Game, v.Self)
	insop := inspection.Parse(v.Game, v.Enemy)
	for tokenID := range v.Self.T.Present {
		token := v.Game.Token(tokenID)
		a := ParseAttackWith(v, insme, insop, token)
		if a == nil {
			continue
		}

		v.Game.Log().Trace("potential", a)

		if attack == nil {
			attack = a
		} else if a.score > attack.score {
			attack = a
		}
	}

	return
}

func ParseAttackWith(view view.T, insme, insop inspection.T, token *game.Token) *Attack {
	if token.T.Body == nil {
		return nil
	} else if !token.T.Awake {
		return nil
	}
	score := 0
	score += token.T.Body.Attack
	score += 2 * (view.Self.T.Life - insop.BeingsAttack)
	score -= 2 * (view.Enemy.T.Life - insme.AwakeBeingsAttack)

	view.Game.Log().With(map[string]any{
		"Score": score,
		"TID":   token.ID,
	}).Trace("potential")

	return &Attack{
		id:    token.ID(),
		score: score,
	}
}
