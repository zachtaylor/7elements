package checktarget

import (
	"errors"

	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/game/token"
)

func OtherPresentBeing(game *game.T, seat *seat.T, me *token.T, id string) (*token.T, error) {
	if token := game.GetToken(id); token == nil {
		return nil, errors.New("no token: " + id)
	} else if s := game.Seats.Get(token.User); s == nil {
		return nil, errors.New("no seat: " + token.Card.Proto.Name)
	} else if !s.Present.Has(token.ID) {
		return nil, errors.New("not present: " + token.Card.Proto.Name)
	} else if token.Body == nil {
		return nil, errors.New("not being: " + token.Card.Proto.Name)
	} else if me.ID == token.ID {
		return nil, errors.New("is me: " + token.Card.Proto.Name)
	} else {
		return token, nil
	}
}
