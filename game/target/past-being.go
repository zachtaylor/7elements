package target

import (
	"errors"

	"github.com/zachtaylor/7elements/game"
)

func PastBeing(g *game.T, seat *game.Seat, arg interface{}) (*game.Card, error) {
	if gcid, ok := arg.(string); !ok {
		return nil, errors.New("no gcid")
	} else if card := g.Cards[gcid]; card == nil {
		return nil, errors.New("no card")
	} else if s := g.GetSeat(card.Username); s == nil {
		return nil, errors.New("no seat")
	} else if card.Body == nil {
		return nil, errors.New("not being")
	} else if !s.HasPastCard(card.Id) {
		return nil, errors.New("not past")
	} else {
		return card, nil
	}
}
