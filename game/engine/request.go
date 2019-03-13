package engine

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/animate"
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/log"
)

func Request(game *game.T, seat *game.Seat, uri string, json vii.Json) *game.State {
	switch uri {
	case "connect":
		Connect(game, seat)
	case "disconnect":
		RequestDisconnect(game, seat)
	case "chat":
		RequestChat(game, seat, json)
	case "pass":
		RequestPass(game, seat, json)
	case game.State.ID():
		RequestGameState(game, seat, json)
	case "play":
		return RequestPlay(game, seat, json, game.State.EventName() != "main" || seat.Username != game.State.Seat)
	case "trigger":
		return RequestTrigger(game, seat, json)
	default:
		game.Log().WithFields(log.Fields{
			"Seat":       seat,
			"RequestURI": uri,
			"State":      game.State.EventName(),
		}).Warn("engine/request: 404")
	}
	return nil
}

func Connect(g *game.T, seat *game.Seat) {
	log := g.Log().Add("Seat", seat).Add("State", g.State)
	animate.GameReconnect(g, seat)
	log.Debug("engine/connect")

	if connector, _ := g.State.Event.(game.ConnectEventer); connector == nil {
		log.Debug("engine/connect: state: no data")
	} else {
		log.Debug("engine/connect: state")
		connector.OnConnect(g, seat)
	}
}

func RequestDisconnect(game *game.T, seat *game.Seat) {
	game.Log().Add("Seat", seat).Info("player disconnect")
	// animate.GameDisconnect(game, seat) // todo
}

func RequestChat(game *game.T, seat *game.Seat, json vii.Json) {
	animate.Chat(game, seat.Username, "game#"+game.ID(), json.Sval("message"))
}

func RequestPass(game *game.T, seat *game.Seat, json vii.Json) {
	if pass := json.Sval("pass"); pass == "" {
		game.Log().WithFields(log.Fields{
			"data": json,
			"Seat": seat.Username,
		}).Warn("engine/request: pass: no target")
	} else if pass != game.State.ID() {
		game.Log().WithFields(log.Fields{
			"pass":  json,
			"State": game.State.ID(),
			"Seat":  seat.Username,
		}).Warn("engine/request: pass: target mismatch")
	} else {
		game.Log().WithFields(log.Fields{
			"Seat":  seat.Username,
			"State": game.State.ID(),
		}).Debug("engine/request: pass")
		game.State.Reacts[seat.Username] = "pass"
		animate.GameReact(game, seat.Username)
	}
}

func RequestPlay(game *game.T, seat *game.Seat, json vii.Json, onlySpells bool) *game.State {
	log := game.Log().WithFields(log.Fields{
		"Username": seat.Username,
		"Elements": seat.Elements,
	})

	gcid := json.Sval("gcid")
	if gcid == "" {
		log.Error("engine/request: play: gcid missing")
		return nil
	}
	log.Add("GCID", gcid)

	card := game.Cards[gcid]
	if card == nil {
		log.Error("engine/request: play: gcid not found")
	} else if card.Username != seat.Username {
		log.Add("Owner", card.Username).Error("engine/request: play: card belongs to a different player")
	} else if card.Card.Type != vii.CTYPspell && onlySpells {
		animate.GameError(seat, game, card.Card.Name, `not "spell" type`)
		log.Error("engine/request: play: not spell type")
	} else if !seat.HasCardInHand(gcid) {
		animate.GameError(seat, game, card.Card.Name, `not in your hand`)
		log.Error("engine/request: play: card not in hand")
	} else if !seat.Elements.GetActive().Test(card.Card.Costs) {
		animate.GameError(seat, game, card.Card.Name, `not enough elements`)
		log.Error("engine/request: play: cannot afford")
	} else {
		log.Info("engine/request: play")
		seat.Elements.Deactivate(card.Card.Costs)
		return game.NewState(seat.Username, Play(game, seat, card, json["target"]))
	}
	return nil
}

func RequestTrigger(game *game.T, seat *game.Seat, json vii.Json) *game.State {
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

	card := seat.Present[gcid]
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
			return game.NewState(seat.Username, Trigger(game, seat, card, power, card))
		} else {
			return game.NewState(seat.Username, Trigger(game, seat, card, power, json["target"]))
		}
	}
	return nil
}

func RequestGameState(g *game.T, seat *game.Seat, json vii.Json) {
	log := g.Logger.WithFields(log.Fields{
		"State":    g.State.EventName(),
		"Username": seat.Username,
	})
	if requester, _ := g.State.Event.(game.RequestEventer); requester == nil {
		log.Warn("request state failed")
	} else {
		log.Debug("request state")
		requester.Request(g, seat, json)
	}
}
