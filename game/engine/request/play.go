package request

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
	pkg_state "github.com/zachtaylor/7elements/game/state"
	"github.com/zachtaylor/7elements/out"
	"ztaylor.me/cast"
)

func play(g *game.T, seat *game.Seat, json cast.JSON, onlySpells bool) []game.Stater {
	log := g.Log().With(cast.JSON{
		"Seat": seat.String(),
	})

	if id := json.GetS("id"); id == "" {
		log.Error("no id")
	} else if c := seat.Hand[id]; c == nil {
		log.Error("no card")
		out.Error(seat.Player, `vii`, `bad card id`)
	} else if c.Proto.Type != card.SpellType && onlySpells {
		log.Add("Card", c.String()).Error("card type must be spell")
		out.Error(seat.Player, c.Proto.Name, `not "spell" type`)
	} else if !seat.Karma.Active().Test(c.Proto.Costs) {
		log.Add("Card", c.String()).Error("not enough elements")
		out.Error(seat.Player, c.Proto.Name, `not enough elements`)
	} else {
		log.Add("Card", c.String()).Info("accept")
		seat.Karma.Deactivate(c.Proto.Costs)
		delete(seat.Hand, id)
		out.GameSeat(g, seat.JSON())
		out.GameHand(seat.Player, seat.Hand.JSON())
		return []game.Stater{pkg_state.NewPlay(seat.Username, c, json["target"])}
	}
	return nil
}
