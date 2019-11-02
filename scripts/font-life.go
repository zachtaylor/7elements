package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/target"
	"github.com/zachtaylor/7elements/game/trigger"
	"ztaylor.me/log"
)

const FontOfLifeID = "font-life"

func init() {
	game.Scripts[FontOfLifeID] = FontOfLife
}

func FontOfLife(g *game.T, seat *game.Seat, arg interface{}) []game.Event {
	log := g.Log().With(log.Fields{
		"Target":   arg,
		"Username": seat.Username,
	}).Tag(logtag + FontOfLifeID)

	card, err := target.PresentBeing(g, seat, arg)
	if err != nil {
		log.Add("Error", err).Error()
		return nil
	}
	log.Info()
	return trigger.Heal(g, card, 2)
}
