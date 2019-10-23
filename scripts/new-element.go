package scripts

import (
	vii "github.com/zachtaylor/7elements"

	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/event"
	"ztaylor.me/cast"
	"ztaylor.me/log"
)

func init() {
	game.Scripts["new-element"] = NewElement
}

func NewElement(g *game.T, seat *game.Seat, target interface{}) []game.Event {
	log := g.Log().With(log.Fields{
		"Seat": seat.Username,
	}).Tag("scripts/new-element")
	me, ok := target.(*game.Card)
	if !ok {
		log.Add("Target", target).Error("target failed")
		return nil
	}
	return []game.Event{event.NewChoiceEvent(
		seat.Username,
		"Create a New Element",
		cast.JSON{
			"card": me.JSON(),
		},
		game.GameChoiceNewElementChoices,
		func(val interface{}) {
			log.Tag("finish")
			if i := cast.Int(val); i < 1 || i > 7 {
				log.Warn("no element")
			} else {
				e := vii.Element(i)
				log.Add("Element", e).Debug()
				seat.Elements.Append(e)
				g.SendSeatUpdate(seat)
			}
		},
	)}
}
