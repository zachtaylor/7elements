package scripts

import (
	vii "github.com/zachtaylor/7elements"

	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/log"
)

const SummonersPortalID = "summoners-portal"

func init() {
	game.Scripts[SummonersPortalID] = SummonersPortal
}

func SummonersPortal(g *game.T, seat *game.Seat, target interface{}) []game.Event {
	card := seat.Deck.Draw()
	log := g.Log().With(log.Fields{
		"Username": seat.Username,
		"Card":     card,
	}).Tag("/scripts/" + SummonersPortalID)

	if card == nil {
		log.Error(`card is nil`)
	} else if card.Card.Type == vii.CTYPbody || card.Card.Type == vii.CTYPitem {
		seat.Present[card.Id] = card
		g.SendAll(game.BuildSpawnUpdate(g, card))
	} else {
		log.Add("BurnedCard", true)
		seat.Past[card.Id] = card
		g.SendAll(game.BuildSeatUpdate(seat))
	}
	log.Info(SummonersPortalID)
	return nil
}
