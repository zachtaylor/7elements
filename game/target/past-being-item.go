package target

import (
	"errors"

	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
)

func PastBeingItem(g *game.T, seat *game.Seat, arg interface{}) (*card.T, error) {
	if id, ok := arg.(string); !ok {
		return nil, errors.New("no id")
	} else if c := g.GetCard(id); c == nil {
		return nil, errors.New("no card: " + id)
	} else if s := g.GetSeat(c.Username); s == nil {
		return nil, errors.New("no seat: " + c.Proto.Name)
	} else if c.Proto.Type != card.BodyType && c.Proto.Type != card.ItemType {
		return nil, errors.New("not being or item: " + c.Proto.Name)
	} else if !s.HasPastCard(c.ID) {
		return nil, errors.New("not past: " + c.Proto.Name)
	} else {
		return c, nil
	}
}
