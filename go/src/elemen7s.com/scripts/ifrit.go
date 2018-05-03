package scripts

import (
	"elemen7s.com"
	"elemen7s.com/games"
)

const IfritID = "ifrit"

func init() {
	games.Scripts[IfritID] = Ifrit
}

func Ifrit(game *games.Game, seat *games.Seat, target interface{}) {
	log := game.Log().Add("Target", target).Add("Username", seat.Username)

	if target == "player" {
		for _, s := range game.Seats {
			if s.Username != seat.Username {
				s.Life--
				games.BroadcastAnimateSeatUpdate(game, s)
				games.BroadcastAnimateSeatUpdate(game, seat)
				game.Active.Activate(game)
				log.Add("Seat", s).Info(IfritID)
				return
			}
		}
	}

	gcid := CastString(target)
	card := game.Cards[gcid]
	if card == nil {
		log.Add("Error", "gcid not found").Error(IfritID)
		return
	} else if ownerSeat := game.GetSeat(card.Username); ownerSeat == nil {
		log.Add("Error", "card owner not found").Error(IfritID)
		return
	} else if !ownerSeat.HasAliveCard(gcid) {
		log.Add("Error", "card not in play").Error(IfritID)
		return
	} else if card.Card.CardType != vii.CTYPbody {
		log.Add("CardType", card.Card.CardType).Add("Error", "card not type body").Error(IfritID)
		return
	}

	games.Damage(game, card, 1)
	games.BroadcastAnimateSeatUpdate(game, game.GetSeat(card.Username))
	game.Active.Activate(game)
	log.Info(IfritID)
}
