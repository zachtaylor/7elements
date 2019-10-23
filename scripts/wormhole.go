package scripts

import (
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/target"
	"ztaylor.me/log"
)

const WormholeID = "wormhole"

func init() {
	game.Scripts[WormholeID] = Wormhole
}

func Wormhole(g *game.T, seat *game.Seat, arg interface{}) []game.Event {
	log := g.Log().With(log.Fields{
		"Username": seat.Username,
	}).Tag(logtag + MemorializeID)
	card, err := target.PresentBeing(g, seat, arg)
	if err != nil {
		log.Add("Error", err).Error()
		return nil
	}
	log.Add("Target", card).Info("confirm")
	card.IsAwake = false
	seat.Hand[card.Id] = card
	delete(seat.Present, card.Id)
	g.SendSeatUpdate(seat)
	return nil
}
