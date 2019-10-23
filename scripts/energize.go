package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/target"
	"ztaylor.me/log"
)

const EnergizeID = "energize"

func init() {
	game.Scripts[EnergizeID] = Energize
}

func Energize(g *game.T, seat *game.Seat, arg interface{}) []game.Event {
	log := g.Log().With(log.Fields{
		"Target":   arg,
		"Username": seat.Username,
	}).Tag(logtag + EnergizeID)
	card, err := target.PresentBeingItem(g, seat, arg)
	if err != nil {
		log.Add("Error", err).Error()
		return nil
	}
	log.Info()
	card.IsAwake = true
	g.SendCardUpdate(card)
	return nil
}
