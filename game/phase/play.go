package phase

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/element"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/trigger"
)

func NewPlay(g *game.G, playerID string, card *game.Card, karma element.Count, targets []string) game.Phaser {
	return &Play{
		PriorityContext: game.PriorityContext(g.NewPriority(playerID)),
		Card:            card,
		Karma:           karma,
		Targets:         targets,
	}
}

type Play struct {
	game.PriorityContext
	Card        *game.Card
	Karma       element.Count
	Targets     []string
	IsCancelled bool
}

func (*Play) Type() string      { return "play" }
func (*Play) Next() game.Phaser { return nil }

// OnActivate implements game.OnActivatePhaser
func (r *Play) OnActivate(game *game.G) []game.Phaser {
	// msg := r.Card.Proto.Name
	// if r.Card.Proto.Text != "" {
	// 	msg = r.Card.Proto.Text
	// }
	// go game.Chat(r.Seat(), msg)
	return nil
}

// Finish implements game.OnFinishPhaser
func (r *Play) OnFinish(g *game.G, _ *game.State) (rs []game.Phaser) {
	player := g.Player(r.Priority()[0])
	g.Log().With(map[string]any{
		"Player":      player,
		"Card":        r.Card,
		"IsCancelled": r.IsCancelled,
	}).Debug("play finish")
	player.T.Past.Add(r.Card.ID())
	g.MarkUpdate(player.ID())

	if r.IsCancelled {
		return nil
	}

	if r.Card.T.Kind == card.Being || r.Card.T.Kind == card.Item {
		ctx := game.NewTokenContext(r.Card)
		if triggered := trigger.TokenAdd(g, player, ctx); len(triggered) > 0 {
			rs = append(rs, triggered...)
		}
	}

	if powers := r.Card.T.Powers.GetTrigger("play"); len(powers) > 0 {
		for _, power := range powers {
			ctx := game.NewScriptContext(power.Script, r.Card.ID(), r.Card.Player(), r.Karma, r.Targets)
			if triggered := game.RunScript(g, ctx); len(triggered) > 0 {
				rs = append(rs, triggered...)
			}
		}
	}

	return
}

func (r *Play) JSON() map[string]any {
	json := map[string]any{
		"card":    r.Card.T.ID,
		"targets": r.Targets,
	}
	return json
}
