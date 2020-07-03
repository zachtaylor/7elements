package target

import (
	"errors"

	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
)

func PastBeing(g *game.T, seat *game.Seat, arg interface{}) (*card.T, error) {
	if id, ok := arg.(string); !ok {
		return nil, errors.New("no id")
	} else if c := g.GetCard(id); c == nil {
		return nil, errors.New("not card: " + id)
	} else if s := g.GetSeat(c.Username); s == nil {
		return nil, errors.New("no seat: " + id)
	} else if c.Proto.Type != card.BodyType {
		return nil, errors.New("not being: " + id)
	} else if !s.HasPastCard(c.ID) {
		return nil, errors.New("not past: " + id)
	} else {
		return c, nil
	}
}
