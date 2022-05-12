package request

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/element"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/phase"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/wsout"
)

func Play(game *game.T, seat *seat.T, json map[string]any, onlySpells bool) (rs []game.Phaser) {
	log := game.Log().With(map[string]any{
		"Seat": seat.String(),
	})

	if id, _ := json["id"].(string); id == "" {
		log.Error("no id")
	} else if c := seat.Hand[id]; c == nil {
		log.Error("no card")
		seat.Writer.WriteMessageData(wsout.Error("vii", "bad card id"))
	} else if c.Proto.Type != card.SpellType && onlySpells {
		log.Add("Card", c.String()).Error("card type must be spell")
		seat.Writer.WriteMessageData(wsout.Error(c.Proto.Name, `not "spell" type`))
	} else if pay, _ := json["pay"].(string); pay == "" {
		log.Warn("elements payment missing")
		seat.Writer.WriteMessageData(wsout.Error(c.Proto.Name, "requires elements payment"))
	} else if paycount, err := element.ParseCount(pay); err != nil {
		log.Add("Error", err).Add("Pay", pay).Error("element pay parse failed")
	} else if activeKarma := seat.Karma.Active(); !activeKarma.Test(paycount) {
		log.Add("PayCount", paycount).Error("karma cannot afford payment")
		seat.Writer.WriteMessageData(wsout.Error(c.Proto.Name, "not enough elements"))
	} else if !paycount.Test(c.Proto.Costs) {
		log.With(map[string]any{
			"Pay":   paycount,
			"Cost":  c.Proto.Costs,
			"Karma": activeKarma,
		}).Out("bogus payment detected")
		seat.Writer.WriteMessageData(wsout.Error("vii", "internal error"))
	} else {
		log.With(map[string]any{
			"Pay":   paycount,
			"Cost":  c.Proto.Costs,
			"Karma": activeKarma,
			"Card":  c.String(),
		}).Info()
		seat.Karma.Deactivate(paycount)
		delete(seat.Hand, id)
		game.Seats.WriteMessageData(wsout.GameSeat(seat.JSON()))
		seat.Writer.WriteMessageData(wsout.GameHand(seat.Hand.Keys()))
		target, _ := json["target"].(string)
		rs = append(rs, phase.NewPlay(seat.Username, c, target))
	}
	return
}
