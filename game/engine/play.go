package engine

import (
	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/state"
	"github.com/zachtaylor/7elements/game/update"
	"ztaylor.me/cast"
)

func Play(g *game.T, seat *game.Seat, json cast.JSON, onlySpells bool) []game.Stater {
	log := g.Log().With(cast.JSON{
		"Seat": seat.String(),
	}).Source()

	if id := json.GetS("id"); id == "" {
		log.Error("no id")
	} else if card := seat.Hand[id]; card == nil {
		log.Error("no card")
		update.ErrorW(seat, `vii`, `bad card id`)
	} else if card.Card.Type != vii.CTYPspell && onlySpells {
		log.Add("Card", card.String()).Error("card type must be spell")
		update.ErrorW(seat, card.Card.Name, `not "spell" type`)
	} else if !seat.Karma.Active().Test(card.Card.Costs) {
		log.Add("Card", card.String()).Error("not enough elements")
		update.ErrorW(seat, card.Card.Name, `not enough elements`)
	} else {
		log.Add("Card", card.String()).Info("accept")
		seat.Karma.Deactivate(card.Card.Costs)
		delete(seat.Hand, id)
		update.Seat(g, seat)
		update.Hand(seat)
		return []game.Stater{state.NewPlay(seat.Username, card, json["target"])}
	}
	return nil
}
