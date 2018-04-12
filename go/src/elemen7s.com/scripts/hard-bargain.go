package scripts

import (
	"elemen7s.com"
	"elemen7s.com/games"
	"strconv"
)

const HardBargainID = "hard-bargain"

func init() {
	games.Scripts[HardBargainID] = HardBargain
}

func HardBargain(game *games.Game, seat *games.Seat, target interface{}) {
	log := game.Log().Add("Target", target).Add("Username", seat.Username)

	var gcid int
	switch v := target.(type) {
	case string:
		if t, err := strconv.ParseInt(v, 10, 0); err != nil {
			log.Add("Error", err).Error(HardBargainID + ": parse target gcid")
			return
		} else {
			gcid = int(t)
		}
	case int:
		gcid = v
	case float64:
		gcid = int(v)
	default:
		log.Error(HardBargainID + ": parse unknown type target gcid")
		return
	}

	card := game.Cards[gcid]
	if card == nil {
		log.Add("Error", "gcid not found").Error(HardBargainID)
	} else if ownerSeat := game.GetSeat(card.Username); ownerSeat == nil {
		log.Add("Error", "card owner not found").Error(HardBargainID)
	} else if !ownerSeat.HasAliveCard(gcid) {
		log.Add("Error", "card not in play").Error(HardBargainID)
	} else if card.Card.CardType != vii.CTYPitem {
		log.Add("CardType", card.Card.CardType).Add("Error", "card not type item").Error(HardBargainID)
	} else {
		delete(ownerSeat.Alive, gcid)
		ownerSeat.Graveyard[gcid] = card
		games.BroadcastAnimateSeatUpdate(game, seat)
		games.BroadcastAnimateSeatUpdate(game, ownerSeat)
		game.Active.OnActivate(game.Active, game)
		log.Info(HardBargainID)
	}
}
