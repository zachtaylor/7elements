package target

import (
	"errors"

	"github.com/zachtaylor/7elements/game"
)

func OtherPresentBeing(g *game.T, seat *game.Seat, me_tok *game.Token, arg interface{}) (*game.Token, error) {
	if id, ok := arg.(string); !ok {
		return nil, errors.New("no id")
	} else if target := g.GetToken(id); target == nil {
		return nil, errors.New("no token: " + id)
	} else if s := g.GetSeat(target.Username); s == nil {
		return nil, errors.New("no seat: " + target.Card.Proto.Name)
	} else if !s.HasPresent(target.ID) {
		return nil, errors.New("not present: " + target.Card.Proto.Name)
	} else if target.Body == nil {
		return nil, errors.New("not being: " + target.Card.Proto.Name)
	} else if me_tok.ID == target.ID {
		return nil, errors.New("is me: " + target.Card.Proto.Name)
	} else {
		return target, nil
	}
}
