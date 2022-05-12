package target

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game/v2"
)

func MyPastBeing(g *game.G, player *game.Player, id string) (*game.Card, error) {
	if c := g.Card(id); c == nil {
		return nil, ErrNotCard
	} else if c.T.Kind != card.Being {
		return nil, ErrNotBeing
	} else if c.Player() != player.ID() {
		return nil, ErrNotMine
	} else if !player.T.Past.Has(id) {
		return nil, ErrNotPast
	} else {
		return c, nil
	}
}
