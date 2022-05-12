package phase

import (
	"github.com/zachtaylor/7elements/deck"
	"github.com/zachtaylor/7elements/game/trigger"
	"github.com/zachtaylor/7elements/game/v2"
	"taylz.io/yas"
)

func NewStart(priority game.Priority) game.Phaser {
	return &Start{
		PriorityContext: game.PriorityContext(priority),
		Ans:             make(map[string]string),
	}
}

type Start struct {
	game.PriorityContext
	Ans map[string]string
}

func (r *Start) Type() string { return "start" }

// OnActivate implements game.OnActivatePhaser
func (r *Start) OnActivate(g *game.G) []game.Phaser {
	g.Log().Trace("activate")
	priority := r.Priority()
	for _, playerID := range priority {
		player := g.Player(playerID)
		player.T.Life = g.Rules().PlayerLife
		player.T.Future = deck.Shuffle(player.T.Future)
		_ = trigger.DrawCard(g, player, g.Rules().PlayerHand)
	}
	return nil
}

// OnConnect implements game.OnConnectPhaser
func (r *Start) OnConnect(g *game.G, player *game.Player) {
	// if seat == nil {
	// g.Log().Trace("announce")
	// go game.Chat("sunrise", r.Seat())
	// }
}

// func (r *Start) GetNext(game *game.G) game.Phaser { return NewSunrise(r.Seat()) }

func (r *Start) JSON() map[string]any { return nil }

// Request implements Requester
func (r *Start) OnRequest(g *game.G, state *game.State, player *game.Player, json map[string]any) {
	choice, _ := json["choice"].(string)
	log := g.Log().Add("Player", player).Add("Choice", choice)

	if ans := r.Ans[player.ID()]; ans != "" {
		log.Add("Answer", ans).Warn("already recorded")
		return
	} else if choice == "keep" {
		r.Ans[player.ID()] = "keep"
	} else if choice == "mulligan" {
		r.Ans[player.ID()] = "mulligan"
		for cardID := range player.T.Hand {
			player.T.Past.Set(cardID)
		}
		player.T.Hand = make(yas.Set[string])
		_ = trigger.DrawCard(g, player, g.Rules().PlayerHand)
	} else {
		log.Warn("unrecognized")
		return
	}

	state.T.React.Set(player.ID())
	g.MarkUpdate(state.ID())
	log.Info("confirm")
}
