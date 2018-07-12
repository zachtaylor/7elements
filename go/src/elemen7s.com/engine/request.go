package engine

import (
	"elemen7s.com"
	"elemen7s.com/animate"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func Request(game *vii.Game, t *Timeline, username string, json js.Object) *Timeline {
	seat := game.GetSeat(username)
	if seat == nil {
		return nil
	}

	switch event := json["event"]; event {
	case "play":
		return RequestPlay(game, t, seat, json, t.Name() != "main" || username != t.HotSeat)
	case "trigger":
		return RequestTrigger(game, t, seat, json)
	case "reconnect":
		RequestReconnect(game, t, seat)
		break
	case "chat":
		animate.BroadcastChat(game, username, json.Sval("message"))
		break
	case t.Name():
		RequestTimeline(game, t, seat, json)
		break
	default:
		game.Log().WithFields(log.Fields{
			"Seat":     seat,
			"Timeline": t,
		}).Warn("receive: out of sync")
	}
	return nil
}

func RequestReconnect(game *vii.Game, t *Timeline, seat *vii.GameSeat) {
	json := seat.Json(true)
	opponentsdata := make([]string, 0)
	for _, seat2 := range game.Seats {
		if seat2.Username != seat.Username {
			opponentsdata = append(opponentsdata, seat2.Username)
		}
	}
	json["opponents"] = opponentsdata

	json["gameid"] = game.Key
	json["timer"] = t.Lifetime.Seconds()
	json["mode"] = t.Name()

	seatdata := js.Object{}
	for _, seat := range game.Seats {
		seatdata[seat.Username] = seat.Json(false)
	}
	json["seats"] = seatdata

	seat.Send("game", json)
	t.OnReconnect(game, seat)
	game.Log().Add("Seat", seat).Info("/api/join: socket join")
}

func RequestPass(game *vii.Game, t *Timeline, seat *vii.GameSeat, json js.Object) {
	if json.Sval("mode") != t.Name() {
		game.Log().WithFields(log.Fields{
			"Username":    seat.Username,
			"Timeline":    t,
			"RequestMode": json.Val("mode"),
		}).Warn("try pass out of sync")
	}

	if _, ok := t.Reacts[seat.Username]; ok {
		game.Log().Add("Player", seat).Warn("pass: already recorded")
	} else {
		t.Reacts[seat.Username] = "pass"
	}
}

func RequestPlay(game *vii.Game, t *Timeline, seat *vii.GameSeat, json js.Object, onlySpells bool) *Timeline {
	log := game.Log().WithFields(log.Fields{
		"Username": seat.Username,
		"Elements": seat.Elements,
	})

	gcid := json.Sval("gcid")
	if gcid == "" {
		log.Error("games.RequestPlay: gcid missing")
		return nil
	}
	log.Add("GCID", gcid)

	card := game.Cards[gcid]
	if card == nil {
		log.Error("games.RequestPlay: gcid not found")
	} else if card.Username != seat.Username {
		log.Add("Owner", card.Username).Error("games.RequestPlay: card belongs to a different player")
	} else if card.Card.CardType != vii.CTYPspell && onlySpells {
		animate.Error(seat, game, card.CardText.Name, `not "spell" type`)
		log.Error("games.RequestPlay: not spell type")
	} else if !seat.HasCardInHand(gcid) {
		animate.Error(seat, game, card.CardText.Name, `not in your hand`)
		log.Error("games.RequestPlay: not in your hand")
	} else if !seat.Elements.GetActive().Test(card.Card.Costs) {
		animate.Error(seat, game, card.CardText.Name, `not enough elements`)
		log.Error("games.RequestPlay: cannot afford")
	} else {
		seat.Elements.Deactivate(card.Card.Costs)
		return t.Fork(game, Play(game, t, seat, card, json["target"]))
	}
	return nil
}

func RequestTrigger(game *vii.Game, t *Timeline, seat *vii.GameSeat, json js.Object) *Timeline {
	log := game.Log().WithFields(log.Fields{
		"Username": seat.Username,
		"Elements": seat.Elements,
	})

	gcid := json.Sval("gcid")
	if gcid == "" {
		log.Error("games.Trigger: gcid missing")
		return nil
	}

	powerid := json.Ival("powerid")
	if powerid < 1 {
		log.Error("games.Trigger: powerid missing")
		return nil
	}
	log.Add("GCID", gcid).Add("PowerId", powerid)

	card := seat.Alive[gcid]
	if card == nil {
		log.Error("try-trigger: gcid not found")
	} else if power := card.Powers[powerid]; power == nil {
		log.Error("try-trigger: powerid not found")
	} else if !card.IsAwake && power.UsesTurn {
		animate.Error(seat, game, card.CardText.Name, `not awake`)
		log.Error("try-trigger: card is asleep")
	} else if !seat.Elements.GetActive().Test(power.Costs) {
		animate.Error(seat, game, card.CardText.Name, `not enough elements`)
		log.Add("Costs", power.Costs).Error("try-trigger: cannot afford")
	} else {
		seat.Elements.Deactivate(power.Costs)
		card.IsAwake = card.IsAwake && !power.UsesTurn
		if power.Target == "self" {
			return t.Fork(game, Trigger(game, t, seat, card, power, card))
		} else {
			return t.Fork(game, Trigger(game, t, seat, card, power, json["target"]))
		}
	}
	return nil
}

func RequestTimeline(game *vii.Game, t *Timeline, seat *vii.GameSeat, json js.Object) {
	defer func() {
		if err := recover(); err != nil {
			game.Log().Add("Error", err).Error("")
		}
	}()
	t.Receive(game, seat, json)
}
