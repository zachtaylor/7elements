package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/target"
	"github.com/zachtaylor/7elements/game/trigger"
	"ztaylor.me/log"
)

const GraceID = "grace"

func init() {
	game.Scripts[GraceID] = Grace
}

func Grace(g *game.T, seat *game.Seat, arg interface{}) []game.Event {
	log := g.Log().With(log.Fields{
		"Target":   arg,
		"Username": seat.Username,
	}).Tag(logtag + GraceID)

	card, err := target.PresentBeing(g, seat, arg)
	if err != nil {
		log.Add("Error", err).Error()
		return nil
	}
	log.Info()
	return trigger.Heal(g, card, 2)
}
