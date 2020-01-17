package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/target"
	"github.com/zachtaylor/7elements/game/trigger"
	"ztaylor.me/cast"
)

const IfritID = "ifrit"

func init() {
	game.Scripts[IfritID] = Ifrit
}

func Ifrit(g *game.T, s *game.Seat, me interface{}, args []interface{}) (events []game.Stater, err error) {
	token, ok := me.(*game.Token)
	if !ok {
		g.Log().Add("me", me).Error("i must be a game token")
		return
	}

	if len(args) < 1 {
		err = ErrNoTarget
	} else if arg := args[0]; arg == s.Username {
		events = trigger.DamageSeat(g, token.Card, s, 1)
	} else if seat := g.Seats[cast.String(arg)]; seat != nil {
		events = trigger.DamageSeat(g, token.Card, seat, 1)
	} else if token, e := target.PresentBeing(g, s, arg); e != nil {
		err = e
	} else if token == nil {
		err = ErrNoTarget
	} else {
		events = trigger.Damage(g, token, 1)
	}
	return
}
