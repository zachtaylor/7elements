package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/target"
	"ztaylor.me/log"
)

const HandrailsID = "handrails"

func init() {
	game.Scripts[HandrailsID] = Handrails
}

func Handrails(g *game.T, seat *game.Seat, arg interface{}) []game.Event {
	log := g.Log().With(log.Fields{
		"Target":   arg,
		"Username": seat.Username,
	}).Tag(logtag + HandrailsID)
	card, err := target.PresentBeing(g, seat, arg)
	if err != nil {
		log.Add("Error", err).Error()
		return nil
	}
	log.Info()
	card.IsAwake = true
	g.SendCardUpdate(card)
	return nil
}
