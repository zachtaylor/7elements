package target

import (
	"errors"

	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
)

func MyPastBeing(g *game.T, seat *game.Seat, arg interface{}) (*card.T, error) {
	if id, ok := arg.(string); !ok {
		return nil, errors.New("no id")
	} else if c := g.GetCard(id); c == nil {
		return nil, errors.New("not card: " + id)
	} else if c.Proto.Type != card.BodyType {
		return nil, errors.New("not being: " + id)
	} else if !seat.HasPastCard(c.ID) {
		return nil, errors.New("not my past: " + id)
	} else {
		return c, nil
	}
}
