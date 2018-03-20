package games

import (
	"elemen7s.com/cards/types"
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
	} else if card.Card.CardType != ctypes.Spell && onlySpells {
		seat.Send("alert", js.Object{
			"class":    "error",
			"gameid":   game.Id,
			"username": card.Text.Name,
			"message":  "not \"spell\" type",
		})
		log.Error("games.Play: not spell type")
		return
	} else if !seat.Elements.GetActive().Test(card.Card.Costs) {
		seat.Send("alert", js.Object{
			"class":    "error",
			"gameid":   game.Id,
			"username": card.Text.Name,
			"message":  "not enough elements",
		})
		log.Error("games.Event: play cannot afford")
		return
	} else if !seat.HasCardInHand(gcid) {
		seat.Send("alert", js.Object{
			"class":    "error",
			"gameid":   game.Id,
			"username": card.Text.Name,
			"message":  card.Text.Name + " is not in your hand",
		})
		log.Error("games.Event: play cannot afford")
		return
	}

	log.Add("CardId", card.Card.Id).Debug("play test")
	Play(e, game, card, seat)
}
