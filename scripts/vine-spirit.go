package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/log"
)

func init() {
	game.Scripts["vine-spirit"] = VineSpirit
}

func VineSpirit(g *game.T, s *game.Seat, target interface{}) []game.Event {
	log := g.Log().With(log.Fields{
		"Target":   target,
		"Username": s.Username,
	}).Tag("scripts/vine-spirit")

	me, ok := target.(*game.Card)
	if !ok {
		log.Add("Target", target).Error("target?")
		return nil
	}
	me.Body.Attack++
	g.SendAll(game.BuildCardUpdate(me))
	log.Info("confirm")
	return nil
}
