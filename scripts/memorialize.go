package scripts

import (
	vii "github.com/zachtaylor/7elements"

	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/cast"
	"ztaylor.me/log"
)

const MemorializeID = "memorialize"

func init() {
	game.Scripts[MemorializeID] = Memorialize
}

func Memorialize(g *game.T, seat *game.Seat, target interface{}) []game.Event {
	log := g.Log().With(log.Fields{
		"Target":   target,
		"Username": seat.Username,
	}).Tag("scripts/memorialize")

	gcid := cast.String(target)
	card := seat.Past[gcid]
	if card == nil {
		log.Error("gcid not found")
	} else if card.Card.Type != vii.CTYPbody {
		log.Add("CardType", card.Card.Type).Error("card not type body")
	} else {
		log.Info()
		c := game.NewCard(card.Card)
		c.Username = seat.Username
		g.RegisterCard(c)
		seat.Hand[c.Id] = c
		g.SendAll(game.BuildSeatUpdate(seat))
		seat.Send(game.BuildHandUpdate(seat))
		return nil
	}

	return nil
}
