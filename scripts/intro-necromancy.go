package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/target"
	"github.com/zachtaylor/7elements/game/trigger"
	"github.com/zachtaylor/7elements/game/update"
)

const IntroToNecromancyID = "intro-necromancy"

func init() {
	game.Scripts[IntroToNecromancyID] = IntroToNecromancy
}

func IntroToNecromancy(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	if len(args) < 1 {
		err = game.ErrNoTarget
	} else if card, _err := target.MyPastBeing(g, s, args[0]); _err != nil {
		err = _err
	} else if card == nil {
		err = game.ErrNoTarget
	} else {
		token, _events := trigger.Spawn(g, s, card)
		if len(_events) > 0 {
			events = append(events, _events...)
		}
		token.Body.Health = 1
		update.Token(g, token)
		events = append(events, trigger.DamageSeat(g, card, s, 1)...)
	}
	return
}
