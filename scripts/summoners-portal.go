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

func SummonersPortal(g *game.T, seat *game.Seat, arg interface{}) []game.Event {
	log := g.Log().With(log.Fields{
		"Username": seat.Username,
	}).Tag(logtag + SummonersPortalID)
	card := seat.Deck.Draw()
	if card == nil {
		log.Error(`card is nil`)
		return nil
	} else if card.Card.Type == vii.CTYPbody || card.Card.Type == vii.CTYPitem {
		seat.Present[card.Id] = card
	} else {
		log.Add("BurnedCard", true)
		g.SendSeatUpdate(seat)
	}
	seat.Past[card.Id] = card
	log.Info()
	return nil
}
