package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/log"
)

const VineSpiritID = "vine-spirit"

func init() {
	game.Scripts[VineSpiritID] = VineSpirit
}

func VineSpirit(g *game.T, seat *game.Seat, arg interface{}) []game.Event {
	log := g.Log().With(log.Fields{
		"Target":   arg,
		"Username": seat.Username,
	}).Tag(logtag + VineSpiritID)
	me, ok := arg.(*game.Card)
	if !ok {
		log.Error("this?")
		return nil
	}
	log.Info()
	me.Body.Attack++
	g.SendCardUpdate(me)
	return nil
}
