package trigger

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/token"
	"github.com/zachtaylor/7elements/wsout"
)

func HealToken(game *game.T, token *token.T, n int) []game.Phaser {
	token.Body.Health += n
	game.Seats.Write(wsout.GameTokenJSON(token.Data()))
	return nil // todo
}
