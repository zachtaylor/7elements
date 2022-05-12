package phase

import (
	"github.com/zachtaylor/7elements/game/trigger"
	"github.com/zachtaylor/7elements/game/v2"
)

type Combat struct {
	game.PriorityContext
	A *game.Token
	B *game.Token
}

func NewCombat(pc game.PriorityContext, t1 *game.Token, t2 *game.Token) game.Phaser {
	return &Combat{
		PriorityContext: pc,
		A:               t1,
		B:               t2,
	}
}

func (r *Combat) Type() string { return "combat" }

// OnActivate implements game.OnActivatePhaser
func (r *Combat) OnActivate(g *game.G) []game.Phaser {
	log := g.Log().With(map[string]any{
		"AttackID": r.A.ID,
		"A♠":       r.A.T.Body.Attack,
		"A♥":       r.A.T.Body.Life,
	})
	if r.B != nil {
		log.With(map[string]any{
			"DefendID": r.B.ID,
			"B♠":       r.B.T.Body.Attack,
			"B♥":       r.B.T.Body.Life,
		}).Info("block")
		// go game.Chat(r.A.Card.Proto.Name, "vs "+r.B.Card.Proto.Name)
	} else {
		priority := r.Priority().Unique()
		enemy := g.Player(priority[1])
		log.Add("Life", enemy.T.Life).Info("going face")
	}

	return nil
}

// // OnConnect implements game.OnConnectPhaser
// func (r *Combat) OnConnect(*game.T, *seat.T) {
// }

// // GetStack implements game.StackEventer
// func (r *Combat) GetStack(game *game.G) *game.State {
// 	return nil
// }

// // Request implements Requester
// func (r *Combat) Request(g*game.T, player *game.Player, json map[string]any) {
// }

// Finish implements game.OnFinishPhaser
func (r *Combat) OnFinish(g *game.G, _ *game.State) (rs []game.Phaser) {
	if r.B != nil {
		if triggered := trigger.TokenDamage(g, r.B, r.A.T.Body.Attack); len(triggered) > 0 {
			rs = append(rs, triggered...)
		}
		g.MarkUpdate(r.B.ID())

		if triggered := trigger.TokenDamage(g, r.A, r.B.T.Body.Attack); len(triggered) > 0 {
			rs = append(rs, triggered...)
		}
		g.MarkUpdate(r.A.ID())

	} else {
		priority := r.Priority().Unique()
		enemy := g.Player(priority[1])

		if triggered := trigger.PlayerDamage(g, enemy, r.A.T.Body.Attack); len(triggered) > 0 {
			rs = append(rs, triggered...)
		}
	}
	return
}

func (r *Combat) JSON() map[string]any {
	return map[string]any{
		"attack": r.A.T.JSON(),
		"block":  r.B.T.JSON(),
	}
}
