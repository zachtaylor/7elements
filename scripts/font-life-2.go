package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/target"
	"github.com/zachtaylor/7elements/game/trigger"
)

const fontoflife2ID = "font-life-2"

func init() {
	game.Scripts[fontoflife2ID] = FontOfLife2
}

func FontOfLife2(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	if token, ok := me.(*game.Token); !ok || token == nil {
		err = ErrMeToken
	} else if len(args) < 1 {
		err = ErrNoTarget
	} else if target, targetErr := target.MyPresentBeing(g, s, args[0]); targetErr == nil || targetErr != nil {
		err = targetErr
	} else {
		events = trigger.Heal(g, target, 1)
	}
	return
}
