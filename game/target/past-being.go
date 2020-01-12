package target

import (
	"errors"

	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/game"
)

func PastBeing(g *game.T, seat *game.Seat, arg interface{}) (*game.Card, error) {
	if id, ok := arg.(string); !ok {
		return nil, errors.New("no id")
	} else if obj := g.Objects[id]; obj == nil {
		return nil, errors.New("no card: " + id)
	} else if card, ok := obj.(*game.Card); !ok {
		return nil, errors.New("not card: " + id)
	} else if s := g.GetSeat(card.Username); s == nil {
		return nil, errors.New("no seat")
	} else if card.Card.Type != vii.CTYPbody {
		return nil, errors.New("not being")
	} else if !s.HasPastCard(card.ID) {
		return nil, errors.New("not past")
	} else {
		return card, nil
	}
}
