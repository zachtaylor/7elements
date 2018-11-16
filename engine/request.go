package engine

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func Request(game *vii.Game, seat *vii.GameSeat, json js.Object) vii.GameEvent {
	switch event := json["event"]; event {
	case "reconnect":
		return RequestReconnect(game, seat)
	case "play":
		return RequestPlay(game, seat, json, game.State.EventName() != "main" || seat.Username != game.State.Seat)
	case "trigger":
		return RequestTrigger(game, seat, json)
	case "chat":
		animate.Chat(game, seat.Username, "game#"+game.Key, json.Sval("message"))
		break
	case game.State.EventName():
		return RequestGameState(game, seat, json)
	default:
		game.Log().WithFields(log.Fields{
			"Seat":         seat,
			"RequestEvent": event,
			"GameState":    game.State.EventName(),
		}).Warn("receive: out of sync")
	}
	return nil
}

func RequestReconnect(game *vii.Game, seat *vii.GameSeat) vii.GameEvent {
	seat.WriteJson(animate.Build("/game", game.Json(seat.Username)))
	animate.GameState(game)
	game.Log().Add("Seat", seat).Info("player reconnect")
	game.State.Event.OnReconnect(game, seat)
	return nil
}

func RequestPass(game *vii.Game, seat *vii.GameSeat, json vii.Json) {
	if json.Sval("mode") != game.State.EventName() {
		game.Log().WithFields(log.Fields{
			"Username":    seat.Username,
			"Timeline":    game.State.EventName(),
			"RequestMode": json.Val("mode"),
		}).Warn("try pass out of sync")
	}

	if _, ok := game.State.Reacts[seat.Username]; ok {
		game.Log().Add("Player", seat).Warn("pass: already recorded")
	} else {
		game.State.Reacts[seat.Username] = "pass"
	}
}

func RequestPlay(game *vii.Game, seat *vii.GameSeat, json js.Object, onlySpells bool) vii.GameEvent {
	log := game.Log().WithFields(log.Fields{
		"Username": seat.Username,
		"Elements": seat.Elements,
	})

	gcid := json.Sval("gcid")
	if gcid == "" {
		log.Error("RequestPlay: gcid missing")
		return nil
	}
	log.Add("GCID", gcid)

	card := game.Cards[gcid]
	if card == nil {
		log.Error("RequestPlay: gcid not found")
	} else if card.Username != seat.Username {
		log.Add("Owner", card.Username).Error("RequestPlay: card belongs to a different player")
	} else if card.Card.Type != vii.CTYPspell && onlySpells {
		animate.GameError(seat, game, card.Card.Name, `not "spell" type`)
		log.Error("RequestPlay: not spell type")
	} else if !seat.HasCardInHand(gcid) {
		animate.GameError(seat, game, card.Card.Name, `not in your hand`)
		log.Error("RequestPlay: not in your hand")
	} else if !seat.Elements.GetActive().Test(card.Card.Costs) {
		animate.GameError(seat, game, card.Card.Name, `not enough elements`)
		log.Error("RequestPlay: cannot afford")
	} else {
		seat.Elements.Deactivate(card.Card.Costs)
		return Play(game, seat, card, json["target"])
	}
	return nil
}

func RequestTrigger(game *vii.Game, seat *vii.GameSeat, json js.Object) vii.GameEvent {
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
		animate.GameError(seat, game, card.Card.Name, `not awake`)
		log.Error("try-trigger: card is asleep")
	} else if !seat.Elements.GetActive().Test(power.Costs) {
		animate.GameError(seat, game, card.Card.Name, `not enough elements`)
		log.Add("Costs", power.Costs).Error("try-trigger: cannot afford")
	} else {
		seat.Elements.Deactivate(power.Costs)
		card.IsAwake = card.IsAwake && !power.UsesTurn
		if power.Target == "self" {
			return Trigger(game, seat, card, power, card)
		} else {
			return Trigger(game, seat, card, power, json["target"])
		}
	}
	return nil
}

func RequestGameState(game *vii.Game, seat *vii.GameSeat, json js.Object) vii.GameEvent {
	game.Log().Protect(func() {
		game.State.Event.Receive(game, seat, json)
	})

	if !(len(game.State.Reacts) < len(game.Seats)) {
		game.State.Timer = 1
	}
	return nil
}
