package scripts

import (
	"fmt"

	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/event"
	"github.com/zachtaylor/7elements/game/target"
	"ztaylor.me/log"
)

const PixieID = "pixie"

func init() {
	game.Scripts[PixieID] = Pixie
}

func Pixie(g *game.T, seat *game.Seat, arg interface{}) []game.Event {
	log := g.Log().With(log.Fields{
		"Arg":      arg,
		"Username": seat.Username,
	}).Tag(logtag + PixieID)
	me, ok := arg.(*game.Card)
	if !ok {
		log.Error("this?")
		return nil
	}
	hp := me.Body.Health
	me.Body.Health = 0
	log.Add("Heal", hp)
	return []game.Event{event.NewTargetEvent(
		seat.Username,
		"player-being",
		fmt.Sprintf("Target Player or Being gains %d Health", hp),
		func(val string) []game.Event {
			if val == seat.Username {
				seat.Life += hp
				g.SendSeatUpdate(seat)
				return nil
			} else if opponent := g.GetOpponentSeat(seat.Username); val == opponent.Username {
				opponent.Life += hp
				g.SendSeatUpdate(opponent)
				return nil
			}
			card, err := target.PresentBeing(g, seat, val)
			if err != nil {
				log.Add("Error", err).Error()
				return nil
			}
			log.Add("Target", card).Info()
			card.Body.Health += hp
			g.SendCardUpdate(card)
			return nil
		},
	)}
}
