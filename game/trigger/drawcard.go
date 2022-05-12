package trigger

import (
	"github.com/zachtaylor/7elements/game/out"
	"github.com/zachtaylor/7elements/game/v2"
	"github.com/zachtaylor/7elements/yas/slices"
)

func DrawCard(g *game.G, player *game.Player, n int) []game.Phaser {
	future := player.T.Future
	if count := len(future); count < n {
		n = count
	}
	newcards := make([]string, n)
	for i := 0; i < n; i++ {
		newcards[i], future = slices.Shift(future)
	}
	if len(newcards) > 0 {
		for _, cardID := range newcards {
			player.T.Hand.Set(cardID)
			out.PrivateCard(g, player, cardID)
		}

		g.MarkUpdate(player.ID())
		out.PrivateHand(g, player)
	}
	return nil // todo
}
