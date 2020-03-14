package request

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/game"
	pkg_state "github.com/zachtaylor/7elements/game/state"
	"github.com/zachtaylor/7elements/game/update"
	"ztaylor.me/cast"
)

func play(g *game.T, seat *game.Seat, json cast.JSON, onlySpells bool) []game.Stater {
	log := g.Log().With(cast.JSON{
		"Seat": seat.String(),
	}).Source()

	if id := json.GetS("id"); id == "" {
		log.Error("no id")
	} else if c := seat.Hand[id]; c == nil {
		log.Error("no card")
		update.ErrorW(seat, `vii`, `bad card id`)
	} else if c.Proto.Type != card.SpellType && onlySpells {
		log.Add("Card", c.String()).Error("card type must be spell")
		update.ErrorW(seat, c.Proto.Name, `not "spell" type`)
	} else if !seat.Karma.Active().Test(c.Proto.Costs) {
		log.Add("Card", c.String()).Error("not enough elements")
		update.ErrorW(seat, c.Proto.Name, `not enough elements`)
	} else {
		log.Add("Card", c.String()).Info("accept")
		seat.Karma.Deactivate(c.Proto.Costs)
		delete(seat.Hand, id)
		update.Seat(g, seat)
		update.Hand(seat)
		return []game.Stater{pkg_state.NewPlay(seat.Username, c, json["target"])}
	}
	return nil
}
