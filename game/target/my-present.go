package target

import (
	"errors"

	"github.com/zachtaylor/7elements/game"
)

func MyPresent(g *game.T, seat *game.Seat, arg interface{}) (*game.Token, error) {
	if id, ok := arg.(string); !ok {
		return nil, errors.New("no id")
	} else if t := g.Objects[id]; t == nil {
		return nil, errors.New("no token: " + id)
	} else if token, ok := t.(*game.Token); !ok {
		return nil, errors.New("not token: " + id)
	} else if token.Username != seat.Username {
		return nil, errors.New("not mine")
	} else if seat.Present[token.ID] == nil {
		return nil, errors.New("not present")
	} else {
		return token, nil
	}
}
