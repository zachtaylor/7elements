package checktarget

import (
	"errors"

	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/seat"
)

func PastCard(game *game.T, seat *seat.T, id string) (*card.T, error) {
	if c := game.GetCard(id); c == nil {
		return nil, errors.New("not card: " + id)
	} else if s := game.Seats.Get(c.User); s == nil {
		return nil, errors.New("no seat: " + id)
	} else if !s.Past.Has(c.ID) {
		return nil, errors.New("not past: " + id)
	} else {
		return c, nil
	}
}
