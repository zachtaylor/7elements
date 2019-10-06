package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/log"
)

const GraceID = "grace"

func init() {
	game.Scripts[GraceID] = Grace
}

func Grace(g *game.T, seat *game.Seat, target interface{}) []game.Event {
	log := g.Log().With(log.Fields{
		"Target":   target,
		"Username": seat.Username,
	}).Tag("scripts/grace")

	if card := game.TargetBeing(g, target); card == nil {
		log.Warn("target failed")
	} else {
		card.Body.Health += 2
		g.SendAll(game.BuildCardUpdate(card))
	}
	return nil
}
