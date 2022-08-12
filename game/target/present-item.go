package target

import "github.com/zachtaylor/7elements/game"

func PresentItem(g *game.G, id string) (*game.Token, error) {
	if token := g.Token(id); token == nil {
		return nil, ErrNotToken
	} else if token.T.Body != nil {
		return nil, ErrNotItem
	} else if owner := g.Player(token.Player()); owner == nil {
		return nil, ErrBadPlayer
	} else if !owner.T.Present.Has(id) {
		return nil, ErrNotPresent
	} else {
		return token, nil
	}
}
