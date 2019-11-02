package target

import (
	"errors"

	"github.com/zachtaylor/7elements/game"
)

func OtherPresentBeing(g *game.T, seat *game.Seat, card *game.Card, arg interface{}) (*game.Card, error) {
	if gcid, ok := arg.(string); !ok {
		return nil, errors.New("no gcid")
	} else if c := g.Cards[gcid]; card == nil {
		return nil, errors.New("no card: " + gcid)
	} else if s := g.GetSeat(c.Username); s == nil {
		return nil, errors.New("no seat")
	} else if !s.HasPresentCard(c.Id) {
		return nil, errors.New("not present")
	} else if card.Body == nil {
		return nil, errors.New("not being")
	} else if card.Id == c.Id {
		return nil, errors.New("not other")
	} else {
		return card, nil
	}
}
