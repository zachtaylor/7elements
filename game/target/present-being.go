package target

import (
	"errors"

	"github.com/zachtaylor/7elements/game"
)

func PresentBeing(g *game.T, seat *game.Seat, arg interface{}) (*game.Card, error) {
	if gcid, ok := arg.(string); !ok {
		return nil, errors.New("no gcid")
	} else if card := g.Cards[gcid]; card == nil {
		return nil, errors.New("no card: " + gcid)
	} else if s := g.GetSeat(card.Username); s == nil {
		return nil, errors.New("no seat")
	} else if !s.HasPresentCard(card.Id) {
		return nil, errors.New("not present")
	} else if card.Body == nil {
		return nil, errors.New("not being")
	} else {
		return card, nil
	}
}
