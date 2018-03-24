package games

import (
	"elemen7s.com"
	"ztaylor.me/js"
)

func TryPlay(e *Event, game *Game, seat *Seat, j js.Object, onlySpells bool) {
	log := game.Log().Add("Username", seat.Username).Add("Elements", game.GetSeat(e.Username).Elements.String())

	gcid := j.Ival("gcid")
	if gcid < 1 {
		log.Error("games.Play: gcid missing")
		return
	}

	log.Add("gcid", gcid)
	card := game.Cards[gcid]
	if card == nil {
		log.Error("games.Play: gcid not found")
		return
	} else if card.Username != seat.Username {
		log.Add("Owner", card.Username).Error("games.Play: card belongs to a different player")
		return
	} else if card.Card.CardType != vii.CTYPspell && onlySpells {
		AnimateAlertError(seat, game, card.CardText.Name, `not "spell" type`)
		log.Error("games.Play: not spell type")
		return
	} else if !seat.Elements.GetActive().Test(card.Card.Costs) {
		AnimateAlertError(seat, game, card.CardText.Name, `not enough elements`)
		log.Error("games.Event: play cannot afford")
		return
	} else if !seat.HasCardInHand(gcid) {
		AnimateAlertError(seat, game, card.CardText.Name, `not in your hand`)
		log.Error("games.Event: play cannot afford")
		return
	}

	log.Add("CardId", card.Card.Id).Debug("play test")
	Play(e, game, card, seat)
}
