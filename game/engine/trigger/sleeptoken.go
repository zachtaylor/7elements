package trigger

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/phase"
	"github.com/zachtaylor/7elements/game/token"
	"github.com/zachtaylor/7elements/wsout"
)

func SleepToken(game *game.T, token *token.T) (rs []game.Phaser) {
	wasAwake := token.IsAwake
	token.IsAwake = false
	if wasAwake {
		if triggered := token.Powers.GetTrigger("become-asleep"); len(triggered) > 0 {
			for _, power := range triggered {
				rs = append(rs, phase.NewTrigger(token.User, token, power, token.ID))
			}
		}
		game.Seats.Write(wsout.GameTokenJSON(token.Data()))
	}
	return
}
