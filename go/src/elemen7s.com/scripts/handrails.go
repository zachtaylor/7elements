package scripts

import (
	"elemen7s.com"
	"elemen7s.com/games"
	"strconv"
)

const HandrailsID = "handrails"

func init() {
	games.Scripts[HandrailsID] = Handrails
}

func Handrails(game *games.Game, seat *games.Seat, target interface{}) {
	log := game.Log().Add("Target", target).Add("Username", seat.Username)

	var gcid int
	switch v := target.(type) {
	case string:
		if t, err := strconv.ParseInt(v, 10, 0); err != nil {
			log.Add("Error", err).Error(HandrailsID + ": parse target gcid")
			return
		} else {
			gcid = int(t)
		}
	case int:
		gcid = v
	case float64:
		gcid = int(v)
	default:
		log.Error(HandrailsID + ": parse unknown type target gcid")
		return
	}

	card := game.Cards[gcid]
	if card == nil {
		log.Add("Error", "gcid not found").Error(HandrailsID)
		return
	} else if ownerSeat := game.GetSeat(card.Username); ownerSeat == nil {
		log.Add("Error", "card owner not found").Error(HandrailsID)
		return
	} else if !ownerSeat.HasAliveCard(gcid) {
		log.Add("Error", "card not in play").Error(HandrailsID)
		return
	} else if card.Card.CardType != vii.CTYPbody {
		log.Add("CardType", card.Card.CardType).Add("Error", "card not type body").Error(HandrailsID)
		return
	}

	card.Awake = true
	games.BroadcastAnimateCardUpdate(game, card)
	games.BroadcastAnimateSeatUpdate(game, seat)
	game.Active.OnActivate(game.Active, game)
	log.Info(HandrailsID)
}
