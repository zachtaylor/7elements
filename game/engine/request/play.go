package request

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/element"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/out"
	"github.com/zachtaylor/7elements/game/phase"
)

func Play(game *game.G, state *game.State, player *game.Player, json map[string]any, onlySpells bool) (rs []game.Phaser) {
	log := game.Log().Add("Player", player.ID())

	if id, _ := json["id"].(string); id == "" {
		log.Error("no id")
	} else if !player.T.Hand.Has(id) {
		log.Error("not in hand")
		player.T.Writer.Write(out.ErrorMessage("vii", "bad card id"))
	} else if c := game.Card(id); c == nil {
		log.Error("not card")
		player.T.Writer.Write(out.ErrorMessage("vii", "bad card id"))
	} else if c.T.Kind != card.Spell && onlySpells {
		log.Add("Card", c.ID()).Error("card type must be spell")
		player.T.Writer.Write(out.ErrorMessage(c.T.Name, `not "spell" type`))
	} else if pay, _ := json["pay"].(string); pay == "" {
		log.Warn("elements payment missing")
		player.T.Writer.Write(out.ErrorMessage(c.T.Name, "requires elements payment"))
	} else if paycount, err := element.ParseCount(pay); err != nil {
		log.Add("Error", err).Add("Pay", pay).Error("element pay parse failed")
	} else if activeKarma := player.T.Karma.Active(); !activeKarma.Test(paycount) {
		log.Add("PayCount", paycount).Error("karma cannot afford payment")
		player.T.Writer.Write(out.ErrorMessage(c.T.Name, "not enough elements"))
	} else if !paycount.Test(c.T.Costs) {
		log.With(map[string]any{
			"Pay":   paycount,
			"Cost":  c.T.Costs,
			"Karma": activeKarma,
		}).Out("bogus payment detected")
		player.T.Writer.Write(out.ErrorMessage("vii", "internal error"))
	} else {
		log.With(map[string]any{
			"Pay":   paycount,
			"Cost":  c.T.Costs,
			"Karma": activeKarma,
			"Card":  c.ID(),
		}).Info()
		player.T.Karma.Deactivate(paycount)
		delete(player.T.Hand, id)
		game.MarkUpdate(player.ID())
		out.PrivateHand(player)
		target, _ := json["target"].(string)
		rs = append(rs, phase.NewPlay(game, player.T.Writer.Name(), c, paycount, []string{target}))
	}
	return
}
