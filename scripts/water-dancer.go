package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/target"
	"ztaylor.me/log"
)

const WaterDancerID = "water-dancer"

func init() {
	game.Scripts[WaterDancerID] = WaterDancer
}

func WaterDancer(g *game.T, seat *game.Seat, arg interface{}) []game.Event {
	log := g.Log().With(log.Fields{
		"Target":   arg,
		"Username": seat.Username,
	}).Tag(logtag + WaterDancerID)
	card, err := target.PresentBeing(g, seat, arg)
	if err != nil {
		log.Add("Error", err).Error()
		return nil
	}
	log.Info()
	card.IsAwake = false
	g.SendCardUpdate(card)
	return nil
}
