package target

import (
	"errors"

	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/game"
)

func MyPastBeing(g *game.T, seat *game.Seat, arg interface{}) (*game.Card, error) {
	if cid, ok := arg.(string); !ok {
		return nil, errors.New("no cid")
	} else if c := g.Objects[cid]; c == nil {
		return nil, errors.New("no card: " + cid)
	} else if card, ok := c.(*game.Card); !ok {
		return nil, errors.New("not card: " + card.String())
	} else if card.Card.Type != vii.CTYPbody {
		return nil, errors.New("not being")
	} else if !seat.HasPastCard(card.ID) {
		return nil, errors.New("not my past")
	} else {
		return card, nil
	}
}
