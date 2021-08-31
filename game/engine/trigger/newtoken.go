package trigger

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/game/token"
	"github.com/zachtaylor/7elements/wsout"
)

func NewToken(game *game.T, seat *seat.T, token *token.T) (rs []game.Phaser) {
	game.RegisterToken(token)
	seat.Present[token.ID] = token
	game.Seats.Write(wsout.GameToken(token.Data()).EncodeToJSON())
	game.Seats.Write(wsout.GamePresentJSON(seat.Username, seat.Present.Keys()))
	return
}
