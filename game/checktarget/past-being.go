package checktarget

import (
	"errors"

	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/seat"
)

func PastBeing(game *game.T, seat *seat.T, id string) (*card.T, error) {
	if c := game.GetCard(id); c == nil {
		return nil, errors.New("not card: " + id)
	} else if s := game.Seats.Get(c.User); s == nil {
		return nil, errors.New("no seat: " + id)
	} else if c.Proto.Type != card.BodyType {
		return nil, errors.New("not being: " + id)
	} else if !s.Past.Has(c.ID) {
		return nil, errors.New("not past: " + id)
	} else {
		return c, nil
	}
}
