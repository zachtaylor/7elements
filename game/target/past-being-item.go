package target

import (
	"errors"

	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
)

func PastBeingItem(g *game.T, seat *game.Seat, arg interface{}) (*card.T, error) {
	if id, ok := arg.(string); !ok {
		return nil, errors.New("no id")
	} else if obj := g.Objects[id]; obj == nil {
		return nil, errors.New("no object: " + id)
	} else if c, ok := obj.(*card.T); !ok {
		return nil, errors.New("not card: " + id)
	} else if s := g.GetSeat(c.Username); s == nil {
		return nil, errors.New("no seat")
	} else if c.Proto.Type != card.BodyType && c.Proto.Type != card.ItemType {
		return nil, errors.New("not being or item")
	} else if !s.HasPastCard(c.ID) {
		return nil, errors.New("not past")
	} else {
		return c, nil
	}
}
