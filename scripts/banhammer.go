package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/target"
	"github.com/zachtaylor/7elements/game/trigger"
	"ztaylor.me/log"
)

const BanhammerID = "banhammer"

func init() {
	game.Scripts[BanhammerID] = Banhammer
}

func Banhammer(g *game.T, seat *game.Seat, arg interface{}) []game.Event {
	log := g.Log().With(log.Fields{
		"Target":   arg,
		"Username": seat.Username,
	}).Tag(logtag + BanhammerID)

	card, err := target.PresentBeingItem(g, seat, arg)
	if err != nil {
		log.Add("Error", err).Error()
		return nil
	}
	log.Info()

	return trigger.Death(g, card)
}
