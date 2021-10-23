package trigger

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/token"
	"github.com/zachtaylor/7elements/wsout"
)

func RemoveToken(game *game.T, token *token.T) (rs []game.Phaser) {
	if token.Body != nil {
		token.Body.Health = 0
	}
	seat := game.Seats.Get(token.User)
	// game.Chat("vii", token.Card.Proto.Name+" died")
	delete(seat.Present, token.ID)
	game.Seats.Write(wsout.GamePresentJSON(seat.Username, seat.Present.Keys()))

	if powers := token.Powers.GetTrigger("death"); len(powers) > 0 {
		for _, power := range powers {
			rs = append(rs, game.Engine().NewTrigger(game, seat, token, power))
		}
	}

	return
}
