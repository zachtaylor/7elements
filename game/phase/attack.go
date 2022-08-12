package phase

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/target"
)

type Attack struct {
	game.PriorityContext
	A *game.Token
	B *game.Token
}

func NewAttack(g *game.G, token *game.Token) game.Phaser {
	return &Attack{
		PriorityContext: game.PriorityContext(g.NewPriority(token.Player())),
		A:               token,
	}
}

func (*Attack) Type() string      { return "attack" }
func (*Attack) Next() game.Phaser { return nil }

// func (r *Attack) String() string { return "Attack {" + r.A.ID + "}" }

// OnActivate implements game.OnActivatePhaser
func (r *Attack) OnActivate(game *game.G) []game.Phaser {
	// go game.Chat(r.A.Card.Proto.Name, "attack")
	return nil
}

// // OnConnect implements game.OnConnectPhaser
// func (r *Attack) OnConnect(*game.T, *seat.T) {
// }

// Finish implements game.OnFinishPhaser
func (r *Attack) OnFinish(*game.G, *game.State) []game.Phaser {
	return []game.Phaser{NewCombat(r.PriorityContext, r.A, r.B)}
}

// // GetStack implements game.StackRer
// func (r *Attack) GetStack(game *game.G) *game.State {
// 	return nil
// }

// GetNext implements game.Phaser
func (r *Attack) GetNext(*game.G) game.Phaser {
	return nil
}

func (r *Attack) JSON() map[string]any { return r.A.T.Data() }

// Request implements Requester
func (r *Attack) OnRequest(g *game.G, state *game.State, player *game.Player, json map[string]any) {
	log := g.Log().Add("Player", player)
	if priority := r.Priority(); player.ID() == priority[0] {
		log.Add("Priority", priority).Warn("seat mismatch")
	} else if id, _ := json["id"].(string); id == "" {
		log.Add("ID", json["id"]).Warn("id missing")
	} else if t, err := target.MyPresentBeing(g, player, id); t == nil {
		log.Add("ID", id).Error(err)
	} else if !t.T.Awake {
		log.Warn("token asleep")
	} else {
		r.B = t
	}

	state.T.React.Add(player.ID())
}
