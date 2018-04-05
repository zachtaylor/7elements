package games

import (
	"elemen7s.com"
	"ztaylor.me/js"
)

func TryPass(game *Game, seat *Seat, json js.Object) {
	if json.Sval("mode") == game.Active.ModeName() {
		game.Active.RespPass(game, seat)
	} else {
		game.Log().Add("Username", seat.Username).Add("RequestMode", json.Val("mode")).Add("CurrentMode", game.Active.ModeName()).Warn("try pass out of sync")
	}
}

func TryPlay(game *Game, seat *Seat, j js.Object, onlySpells bool) {
	log := game.Log().Add("Username", seat.Username).Add("Elements", seat.Elements.String())

	gcid := j.Ival("gcid")
	if gcid < 1 {
		log.Error("games.Play: gcid missing")
		return
	}

	log.Add("GCID", gcid)
	card := game.Cards[gcid]
	if card == nil {
		log.Error("games.Play: gcid not found")
	} else if card.Username != seat.Username {
		log.Add("Owner", card.Username).Error("games.Play: card belongs to a different player")
	} else if card.Card.CardType != vii.CTYPspell && onlySpells {
		AnimateAlertError(seat, game, card.CardText.Name, `not "spell" type`)
		log.Error("games.Trigger: not spell type")
	} else if !seat.HasCardInHand(gcid) {
		AnimateAlertError(seat, game, card.CardText.Name, `not in your hand`)
		log.Error("games.Trigger: play cannot afford")
	} else if !seat.Elements.GetActive().Test(card.Card.Costs) {
		AnimateAlertError(seat, game, card.CardText.Name, `not enough elements`)
		log.Error("games.Trigger: play cannot afford")
	} else {
		Play(game, seat, card, nil)
	}
}

func TryTrigger(game *Game, seat *Seat, json js.Object) {
	log := game.Log().Add("Username", seat.Username).Add("Elements", seat.Elements.String())

	gcid := json.Ival("gcid")
	if gcid < 1 {
		log.Error("games.Trigger: gcid missing")
		return
	}

	powerid := json.Ival("powerid")
	if gcid < 1 {
		log.Error("games.Trigger: powerid missing")
		return
	}

	log.Add("GCID", gcid).Add("PowerId", powerid)
	card := seat.Alive[gcid]
	if card == nil {
		log.Error("games.Trigger: gcid not found")
	} else if power := card.Powers[powerid]; power == nil {
		log.Error("games.Trigger: powerid not found")
	} else if !card.Awake {
		log.Error("games.Trigger: card is asleep")
	} else if !seat.Elements.GetActive().Test(power.Costs) {
		AnimateAlertError(seat, game, card.CardText.Name, `not enough elements`)
		log.Error("games.Trigger: cannot afford")
	} else {
		Trigger(game, seat, card, power, nil)
	}
}
