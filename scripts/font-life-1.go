package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/trigger"
)

const fontoflife1ID = "font-life-1"

func init() {
	game.Scripts[fontoflife1ID] = FontOfLife1
}

func FontOfLife1(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	if token, ok := me.(*game.Token); !ok || token == nil {
		err = ErrMeToken
	} else {
		events = trigger.HealSeat(g, token.Card, s, 1)
	}
	return
}
