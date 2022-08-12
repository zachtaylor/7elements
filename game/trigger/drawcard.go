package trigger

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/out"
	"taylz.io/yas"
)

func DrawCard(g *game.G, player *game.Player, n int) []game.Phaser {
	future := player.T.Future
	if count := len(future); count < n {
		n = count
	}
	newcards := make([]string, n)
	for i := 0; i < n; i++ {
		newcards[i], future = yas.Shift(future)
	}
	if len(newcards) > 0 {
		for _, cardID := range newcards {
			player.T.Hand.Add(cardID)
			out.PrivateCard(player, g.Card(cardID))
		}

		g.MarkUpdate(player.ID())
		out.PrivateHand(player)
	}
	return nil // todo
}
