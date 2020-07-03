package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/target"
	"github.com/zachtaylor/7elements/game/trigger"
)

const HandrailsID = "handrails"

func init() {
	game.Scripts[HandrailsID] = Handrails
}

func Handrails(g *game.T, seat *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	if token, ok := me.(*game.Token); !ok || token == nil {
		err = ErrMeToken
	} else if len(args) < 1 {
		err = ErrNoTarget
	} else if target, targetErr := target.MyPresentBeing(g, seat, args[0]); targetErr != nil {
		err = targetErr
	} else {
		events = trigger.Wake(g, target)

	}
	return
}
