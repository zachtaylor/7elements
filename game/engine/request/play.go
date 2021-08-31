package request

import (
	"github.com/zachtaylor/7elements/card"
	"github.com/zachtaylor/7elements/element"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/phase"
	"github.com/zachtaylor/7elements/game/seat"
	"github.com/zachtaylor/7elements/wsout"
	"taylz.io/http/websocket"
)

func Play(game *game.T, seat *seat.T, json websocket.MsgData, onlySpells bool) (rs []game.Phaser) {
	log := game.Log().With(websocket.MsgData{
		"Seat": seat.String(),
	})

	if id, _ := json["id"].(string); id == "" {
		log.Error("no id")
	} else if c := seat.Hand[id]; c == nil {
		log.Error("no card")
		seat.Writer.Write(wsout.ErrorJSON("vii", "bad card id"))
	} else if c.Proto.Type != card.SpellType && onlySpells {
		log.Add("Card", c.String()).Error("card type must be spell")
		seat.Writer.Write(wsout.ErrorJSON(c.Proto.Name, `not "spell" type`))
	} else if pay, _ := json["pay"].(string); pay == "" {
		log.Warn("elements payment missing")
		seat.Writer.Write(wsout.ErrorJSON(c.Proto.Name, "requires elements payment"))
	} else if paycount, err := element.ParseCount(pay); err != nil {
		log.Add("Error", err).Add("Pay", pay).Error("element pay parse failed")
	} else if activeKarma := seat.Karma.Active(); !activeKarma.Test(paycount) {
		log.Add("PayCount", paycount).Error("karma cannot afford payment")
		seat.Writer.Write(wsout.ErrorJSON(c.Proto.Name, "not enough elements"))
	} else if !paycount.Test(c.Proto.Costs) {
		log.With(websocket.MsgData{
			"Pay":   paycount,
			"Cost":  c.Proto.Costs,
			"Karma": activeKarma,
		}).Out("bogus payment detected")
		seat.Writer.Write(wsout.ErrorJSON("vii", "internal error"))
	} else {
		log.With(websocket.MsgData{
			"Pay":   paycount,
			"Cost":  c.Proto.Costs,
			"Karma": activeKarma,
			"Card":  c.String(),
		}).Info()
		seat.Karma.Deactivate(paycount)
		delete(seat.Hand, id)
		game.Seats.Write(wsout.GameSeatJSON(seat.Data()))
		seat.Writer.Write(wsout.GameHand(seat.Hand.Keys()).EncodeToJSON())
		target, _ := json["target"].(string)
		rs = append(rs, phase.NewPlay(seat.Username, c, target))
	}
	return
}
