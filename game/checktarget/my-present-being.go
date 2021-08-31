package checktarget

import (
	"errors"

	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/game/token"
)

func MyPresentBeing(game *game.T, seat *seat.T, id string) (*token.T, error) {
	if token := game.GetToken(id); token == nil {
		return nil, errors.New("not token: " + id)
	} else if token.Body == nil {
		return nil, errors.New("not being: " + id)
	} else if token.User != seat.Username {
		return nil, errors.New("not mine: " + id)
	} else if seat.Present[token.ID] == nil {
		return nil, errors.New("not present: " + id)
	} else {
		return token, nil
	}
}
