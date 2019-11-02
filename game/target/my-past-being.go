package target

import (
	"errors"

	"github.com/zachtaylor/7elements/game"
)

func MyPastBeing(g *game.T, seat *game.Seat, arg interface{}) (*game.Card, error) {
	if gcid, ok := arg.(string); !ok {
		return nil, errors.New("no gcid")
	} else if card := g.Cards[gcid]; card == nil {
		return nil, errors.New("no card: " + gcid)
	} else if card.Body == nil {
		return nil, errors.New("not being")
	} else if !seat.HasPastCard(card.Id) {
		return nil, errors.New("not my past")
	} else {
		return card, nil
	}
}
