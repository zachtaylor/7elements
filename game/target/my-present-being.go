package target

import (
	"errors"

	"github.com/zachtaylor/7elements/game"
)

func MyPresentBeing(g *game.G, player *game.Player, id string) (*game.Token, error) {
	if token := g.Token(id); token == nil {
		return nil, errors.New("not token: " + id)
	} else if token.T.Body == nil {
		return nil, errors.New("not being: " + id)
	} else if token.Player() != player.ID() {
		return nil, errors.New("not mine: " + id)
	} else if !player.T.Present.Has(id) {
		return nil, errors.New("not present: " + id)
	} else {
		return token, nil
	}
}
