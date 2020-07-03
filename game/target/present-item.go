package target

import (
	"errors"

	"github.com/zachtaylor/7elements/game"
)

func PresentItem(g *game.T, seat *game.Seat, arg interface{}) (*game.Token, error) {
	if id, ok := arg.(string); !ok {
		return nil, errors.New("no id")
	} else if token := g.GetToken(id); token == nil {
		return nil, errors.New("not token: " + id)
	} else if s := g.GetSeat(token.Username); s == nil {
		return nil, errors.New("no seat: " + id)
	} else if !s.HasPresent(token.ID) {
		return nil, errors.New("not present: " + id)
	} else if token.Body != nil {
		return nil, errors.New("not item: " + id)
	} else {
		return token, nil
	}
}
