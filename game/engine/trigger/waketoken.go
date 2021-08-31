package trigger

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/token"
	"github.com/zachtaylor/7elements/wsout"
)

func WakeToken(game *game.T, token *token.T) (rs []game.Phaser) {
	wasAwake := token.IsAwake
	token.IsAwake = true
	if !wasAwake {
		game.Seats.Write(wsout.GameTokenJSON(token.Data()))
	}
	return
}
