package scripts

import (
	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/event/end"
	"github.com/zachtaylor/7elements/game/trigger"
	"ztaylor.me/cast"
	"ztaylor.me/log"
)

const IfritID = "ifrit"

func init() {
	game.Scripts[IfritID] = Ifrit
}

func Ifrit(g *game.T, seat *game.Seat, target interface{}) []game.Event {
	log := g.Log().With(log.Fields{
		"Target":   target,
		"Username": seat.Username,
	}).Tag("scripts/ifrit")

	opponent := g.GetOpponentSeat(seat.Username)
	if target == opponent.Username {
		// TODO trigger.DamageSeat
		opponent.Life--
		g.SendSeatUpdate(opponent)
		if opponent.Life < 0 {
			return []game.Event{
				end.New(seat.Username, opponent.Username),
			}
		}
		return nil
	} else if target == seat.Username {
		seat.Life--
		g.SendSeatUpdate(seat)
		if opponent.Life < 0 {
			return []game.Event{
				end.New(opponent.Username, seat.Username),
			}
		}
		return nil
	}

	gcid := cast.String(target)
	card := g.Cards[gcid]
	if card == nil {
		log.Error("gcid not found")
	} else if ownerSeat := g.GetSeat(card.Username); ownerSeat == nil {
		log.Error("card owner not found")
	} else if !ownerSeat.HasPresentCard(gcid) {
		log.Error("card not in present")
	} else if card.Card.Type != vii.CTYPbody {
		log.Add("CardType", card.Card.Type).Error("card not type body")
	} else {
		log.Info("confirm")
		return trigger.Damage(g, card, 1)
	}

	return nil
}
