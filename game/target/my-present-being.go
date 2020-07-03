package target

import (
	"errors"

	"github.com/zachtaylor/7elements/game"
)

func MyPresentBeing(g *game.T, seat *game.Seat, arg interface{}) (*game.Token, error) {
	if id, ok := arg.(string); !ok {
		return nil, errors.New("no id")
	} else if token := g.GetToken(id); token == nil {
		return nil, errors.New("not token: " + id)
	} else if token.Body == nil {
		return nil, errors.New("not being: " + id)
	} else if token.Username != seat.Username {
		return nil, errors.New("not mine: " + id)
	} else if seat.Present[token.ID] == nil {
		return nil, errors.New("not present: " + id)
	} else {
		return token, nil
	}
}
