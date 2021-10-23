package trigger

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/token"
	"github.com/zachtaylor/7elements/wsout"
)

func DamageToken(game *game.T, token *token.T, n int) []game.Phaser {
	token.Body.Health -= n
	// game.Chat("vii", strconv.FormatInt(int64(n), 10)+" damage to "+token.Card.Proto.Name)
	if token.Body.Health > 0 {
		game.Seats.Write(wsout.GameTokenJSON(token.Data()))
	} else {
		return RemoveToken(game, token)
	}
	return nil
}
