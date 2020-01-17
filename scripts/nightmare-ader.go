package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/target"
	"github.com/zachtaylor/7elements/game/trigger"
)

const nightmareaderID = "nightmare-ader"

func init() {
	game.Scripts[nightmareaderID] = NightmareAder
}

func NightmareAder(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	if len(args) < 1 {
		err = ErrNoTarget
	} else if token, e := target.PresentBeing(g, s, args[0]); err != nil {
		err = e
	} else {
		events = trigger.Damage(g, token, token.Body.Attack)
	}
	return
}
