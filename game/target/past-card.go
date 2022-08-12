package target

import "github.com/zachtaylor/7elements/game"

func PastCard(g *game.G, objectID string) (*game.Card, error) {
	if card := g.Card(objectID); card == nil {
		return nil, ErrNotCard
	} else if player := g.Player(card.Player()); player == nil {
		return nil, ErrBadPlayer
	} else if !player.T.Past.Has(card.ID()) {
		return nil, ErrNotPast
	} else {
		return card, nil
	}
}
