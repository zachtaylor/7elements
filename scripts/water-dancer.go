package scripts

import (

	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
	"ztaylor.me/log"
)

func init() {
	game.Scripts["water-dancer"] = WaterDancer
}

func WaterDancer(g *game.T, s *game.Seat, target interface{}) []game.Event {
	log := g.Log().With(log.Fields{
		"Target":   target,
		"Username": s.Username,
	}).Tag("scripts/water-dancer")

	if card := game.TargetBeing(g, cast.String(target)); card == nil {
		log.Warn("target failed")
	} else {
		card.IsAwake = false
		g.SendAll(game.BuildCardUpdate(card))
	}
	return nil
}
