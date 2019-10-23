package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/target"
	"ztaylor.me/log"
)

const InspireGrowthID = "inspire-growth"

func init() {
	game.Scripts[InspireGrowthID] = InspireGrowth
}

func InspireGrowth(g *game.T, seat *game.Seat, arg interface{}) []game.Event {
	log := g.Log().With(log.Fields{
		"Target":   arg,
		"Username": seat.Username,
	}).Tag(logtag + InspireGrowthID)
	card, err := target.PresentBeing(g, seat, arg)
	if err != nil {
		log.Add("Error", err).Error()
		return nil
	}
	log.Info()
	card.Body.Attack++
	g.SendCardUpdate(card)
	return nil
}
