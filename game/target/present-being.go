package target

import "github.com/zachtaylor/7elements/game"

func PresentBeing(g *game.G, id string) (*game.Token, error) {
	if token := g.Token(id); token == nil {
		return nil, ErrNotToken
	} else if token.T.Body == nil {
		return nil, ErrNotBeing
	} else if player := g.Player(token.Player()); player == nil {
		return nil, ErrBadPlayer
	} else if !player.T.Present.Has(id) {
		return nil, ErrNotPresent
	} else {
		return token, nil
	}
}
