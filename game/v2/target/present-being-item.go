package target

import "github.com/zachtaylor/7elements/game/v2"

func PresentBeingItem(g *game.G, id string) (*game.Token, error) {
	if token := g.Token(id); token == nil {
		return nil, ErrNotToken
	} else if player := g.Player(token.Player()); player == nil {
		return nil, ErrBadPlayer
	} else if !player.T.Present.Has(token.ID()) {
		return nil, ErrNotPresent
	} else {
		return token, nil
	}
}
