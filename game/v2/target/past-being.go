package target

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game/v2"
)

func PastBeing(g *game.G, id string) (*game.Card, error) {
	if c := g.Card(id); c == nil {
		return nil, ErrNotCard
	} else if player := g.Player(c.Player()); player == nil {
		return nil, ErrBadPlayer
	} else if c.T.Kind != card.Being {
		return nil, ErrNotBeing
	} else if !player.T.Past.Has(c.ID()) {
		return nil, ErrNotPast
	} else {
		return c, nil
	}
}
