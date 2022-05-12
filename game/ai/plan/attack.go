package plan

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/ai/inspection"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/game/token"
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

func ParseAttack(game *game.T, seat *seat.T) (attack *Attack) {
	if game.Phase() != "main" || game.State.Phase.Seat() != seat.Username {
		game.Log().Add("Phase", game.Phase()).Add("Seat", game.State.Phase.Seat()).Trace("skip")
		return nil
	}

	insme := inspection.Parse(seat)
	insop := inspection.Parse(g.PlayerOpponent(seat.Username))
	for _, token := range seat.Present {
		a := ParseAttackWith(game, seat, insme, insop, token)
		if a == nil {
			continue
		}

		game.Log().Trace("potential", a)

		if attack == nil {
			attack = a
		} else if a.score > attack.score {
			attack = a
		}
	}

	return
}

func ParseAttackWith(game *game.T, seat *seat.T, insme, insop inspection.T, token *token.T) *Attack {
	if token.Body == nil {
		return nil
	} else if !token.IsAwake {
		return nil
	}
	score := 0
	score += token.Body.Attack
	score += 2 * (seat.Life - insop.BeingsAttack)
	score -= 2 * (g.PlayerOpponent(token.User).Life - insme.AwakeBeingsAttack)

	game.Log().With(map[string]any{
		"Score": score,
		"TID":   token.ID,
	}).Trace("potential")

	return &Attack{
		id:    token.ID,
		score: score,
	}
}
