package target

import (
	"errors"

	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
)

func MyPastBeing(g *game.T, seat *game.Seat, arg interface{}) (*game.Card, error) {
	if cid, ok := arg.(string); !ok {
		return nil, errors.New("no cid")
	} else if obj := g.Objects[cid]; obj == nil {
		return nil, errors.New("no object: " + cid)
	} else if c, ok := obj.(*game.Card); !ok {
		return nil, errors.New("not card: " + c.String())
	} else if c.Card.Type != card.BodyType {
		return nil, errors.New("not being")
	} else if !seat.HasPastCard(c.ID) {
		return nil, errors.New("not my past")
	} else {
		return c, nil
	}
}
