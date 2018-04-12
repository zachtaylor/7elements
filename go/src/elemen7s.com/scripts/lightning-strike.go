package scripts

import (
	"elemen7s.com"
	"elemen7s.com/games"
	"strconv"
)

const LightningStrikeID = "lightning-strike"

func init() {
	games.Scripts[LightningStrikeID] = LightningStrike
}

func LightningStrike(game *games.Game, seat *games.Seat, target interface{}) {
	log := game.Log().Add("Target", target).Add("Username", seat.Username)

	var gcid int
	switch v := target.(type) {
	case string:
		if t, err := strconv.ParseInt(v, 10, 0); err != nil {
			log.Add("Error", err).Error(LightningStrikeID + ": parse target gcid")
			return
		} else {
			gcid = int(t)
		}
	case int:
		gcid = v
	case float64:
		gcid = int(v)
	default:
		log.Error(LightningStrikeID + ": parse unknown type target gcid")
		return
	}

	card := game.Cards[gcid]
	if card == nil {
		log.Add("Error", "gcid not found").Error(LightningStrikeID)
	} else if ownerSeat := game.GetSeat(card.Username); ownerSeat == nil {
		log.Add("Error", "card owner not found").Error(LightningStrikeID)
	} else if !ownerSeat.HasAliveCard(gcid) {
		log.Add("Error", "card not in play").Error(LightningStrikeID)
	} else if card.Card.CardType != vii.CTYPbody {
		log.Add("CardType", card.Card.CardType).Add("Error", "card not type body").Error(LightningStrikeID)
	} else {
		games.Damage(game, card, 3)
		games.BroadcastAnimateSeatUpdate(game, seat)
		games.BroadcastAnimateSeatUpdate(game, ownerSeat)
		game.Active.Activate(game)
		log.Info(LightningStrikeID)
	}
}