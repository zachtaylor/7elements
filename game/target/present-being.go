package target

import (
	"errors"

	"github.com/zachtaylor/7elements/game"
)

func PresentBeing(g *game.T, seat *game.Seat, arg interface{}) (*game.Token, error) {
	if id, ok := arg.(string); !ok {
		return nil, errors.New("no id")
	} else if obj := g.Objects[id]; obj == nil {
		return nil, errors.New("no token: " + id)
	} else if token, ok := obj.(*game.Token); !ok {
		return nil, errors.New("not token: " + id)
	} else if s := g.GetSeat(token.Username); s == nil {
		return nil, errors.New("no seat")
	} else if !s.HasPresent(token.ID) {
		return nil, errors.New("not present")
	} else if token.Body == nil {
		return nil, errors.New("not being")
	} else {
		return token, nil
	}
}
