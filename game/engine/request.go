package engine

import (
	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/event"
	"ztaylor.me/cast"
	"ztaylor.me/log"
)

func Request(g *game.T, seat *game.Seat, uri string, json cast.JSON) []game.Event {
	switch uri {
	case "connect":
		requestConnect(g, seat)
	case "disconnect":
		requestDisconnect(g, seat)
	case "chat":
		requestChat(g, seat, json)
	case "pass":
		requestPass(g, seat, json)
	case g.State.ID():
		RequestGameState(g, seat, json)
	case "attack":
		return RequestAttack(g, seat, json)
	case "play":
		return RequestPlay(g, seat, json, g.State.EventName() != "main" || seat.Username != g.State.Event.Seat())
	case "trigger":
		return RequestTrigger(g, seat, json)
	default:
		g.Log().With(log.Fields{
			"Seat":  seat.Print(),
			"URI":   uri,
			"State": g.State.Print(),
		}).Warn("engine/request: 404")
	}
	return nil
}

func requestConnect(g *game.T, s *game.Seat) {
	g.Log().With(log.Fields{
		"Username": s.Username,
		"State":    g.State.Print(),
	}).Debug("engine/connect: seated")
	s.Send(game.BuildGameUpdate(g, s.Username))

	if connector, ok := g.State.Event.(game.ConnectEventer); ok {
		connector.OnConnect(g, s)
	}
}

func requestDisconnect(g *game.T, s *game.Seat) {
	g.Log().With(log.Fields{
		"Username": s.Username,
		"State":    g.State,
	}).Tag("engine/disconnect: left")

	if disconnector, ok := g.State.Event.(game.DisconnectEventer); ok {
		disconnector.OnDisconnect(g, s)
	}
}

func requestChat(g *game.T, seat *game.Seat, json cast.JSON) {
	text := json.GetS("text")
	g.Log().With(log.Fields{
		"Username": seat.Username,
		"Text":     text,
	}).Debug("engine/chat") // died after
	go g.GetChat().AddMessage(chat.NewMessage(seat.Username, text))
}

func requestPass(g *game.T, seat *game.Seat, json cast.JSON) {
	log := g.Log().With(log.Fields{
		"State":    g.State.Print(),
		"Username": seat.Username,
	}).Tag("engine/pass")
	if pass := json.GetS("pass"); pass == "" {
		log.Warn("target missing")
	} else if pass != g.State.ID() {
		log.Add("PassID", pass).Warn("target mismatch")
	} else {
		g.State.Reacts[seat.Username] = "pass"
		g.SendAll(game.BuildReactUpdate(g, seat.Username))
	}
}

// RequestAttack causes AttackEvent to stack on MainEvent
func RequestAttack(g *game.T, s *game.Seat, json cast.JSON) []game.Event {
	log := g.Log().With(log.Fields{
		"Seat": s.Print(),
	}).Tag("engine/attack")

	if gcid := json.GetS("gcid"); gcid == "" {
		log.Error("gcid missing")
	} else if card := g.Cards[gcid]; card == nil {
		log.Add("GCID", gcid).Error("gcid invalid")
	} else if card.Username != s.Username {
		log.Add("Owner", card.Username).Error("card belongs to a different player")
	} else if card.Card.Type != vii.CTYPbody {
		log.Add("Card", card.Print()).Error("card type must be body")
		s.Send(game.BuildErrorUpdate(card.Card.Name, `not "body" type`))
	} else if s.Present[gcid] == nil {
		log.Add("Card", card.Print()).Error("card must be in present")
		s.Send(game.BuildErrorUpdate(card.Card.Name, `not in your present`))
	} else if !card.IsAwake {
		log.Add("Card", card.Print()).Error("card must be awake")
		s.Send(game.BuildErrorUpdate(card.Card.Name, `not awake`))
	} else {
		log.Add("Card", card.Print()).Info("accept")
		card.IsAwake = false
		g.SendAll(game.BuildCardUpdate(card))
		return []game.Event{event.NewAttackEvent(s.Username, card)}
	}
	return nil
}

func RequestPlay(g *game.T, seat *game.Seat, json cast.JSON, onlySpells bool) []game.Event {
	log := g.Log().With(log.Fields{
		"Seat": seat.Print(),
	}).Tag("engine/play")

	gcid := json.GetS("gcid")
	if gcid == "" {
		log.Error("gcid missing")
		return nil
	}

	card := g.Cards[gcid]
	if card == nil {
		log.Add("GCID", gcid).Error("gcid invalid")
	} else if card.Username != seat.Username {
		log.Add("Owner", card.Username).Error("card belongs to a different player")
	} else if card.Card.Type != vii.CTYPspell && onlySpells {
		log.Add("Card", card.Print()).Error("card type must be spell")
		seat.Send(game.BuildErrorUpdate(card.Card.Name, `not "spell" type`))
	} else if seat.Hand[gcid] == nil {
		log.Add("Card", card.Print()).Error("card must be in hand")
		seat.Send(game.BuildErrorUpdate(card.Card.Name, `not in your hand`))
	} else if !seat.Elements.GetActive().Test(card.Card.Costs) {
		log.Add("Card", card.Print()).Error("not enough elements")
		seat.Send(game.BuildErrorUpdate(card.Card.Name, `not enough elements`))
	} else {
		log.Add("Card", card.Print()).Info("accept")
		seat.Elements.Deactivate(card.Card.Costs)
		delete(seat.Hand, gcid)
		g.SendAll(game.BuildSeatUpdate(seat))
		seat.Send(game.BuildHandUpdate(seat))
		return []game.Event{event.Play(seat.Username, card, json["target"])}
	}
	return nil
}

func RequestTrigger(g *game.T, seat *game.Seat, json cast.JSON) []game.Event {
	log := g.Log().With(log.Fields{
		"Seat":     seat.Username,
		"Elements": seat.Elements,
	}).Tag("engine/trigger")

	gcid := json.GetS("gcid")
	if gcid == "" {
		log.Error("gcid missing")
		return nil
	}
	log.Add("GCID", gcid)

	card := seat.Present[gcid]
	if card == nil {
		log.Error("gcid not found")
		return nil
	} else if card.Username != seat.Username {
		log.Add("Owner", card.Id).Error("card doesn't belong to you")
		return nil
	}

	powerid := json.GetI("powerid")
	if powerid < 1 {
		log.Error("powerid missing")
		return nil
	}
	log.Add("PowerId", powerid)

	power := card.Powers[powerid]
	if power == nil {
		log.Error("powerid not found")
		return nil
	} else if !card.IsAwake && power.UsesTurn {
		seat.Send(game.BuildErrorUpdate(card.Card.Name, `not awake`))
		log.Error("card is asleep")
		return nil
	} else if !seat.Elements.GetActive().Test(power.Costs) {
		seat.Send(game.BuildErrorUpdate(card.Card.Name, `not enough elements`))
		log.Add("Costs", power.Costs).Error("cannot afford")
		return nil
	}

	if power.Costs.Count() > 0 {
		seat.Elements.Deactivate(power.Costs)
	}
	card.IsAwake = card.IsAwake && !power.UsesTurn // if power.UsesTurn {awake=false}
	if power.UsesKill {
		delete(seat.Present, card.Id)
	}
	g.SendAll(game.BuildSeatUpdate(seat))

	if power.Target == "self" {
		return []game.Event{event.NewTriggerEvent(seat.Username, card, power, card)}
	}
	return []game.Event{event.NewTriggerEvent(seat.Username, card, power, json["target"])}
}

func RequestGameState(g *game.T, seat *game.Seat, json cast.JSON) {
	if requester, ok := g.State.Event.(game.RequestEventer); ok {
		requester.Request(g, seat, json)
	} else {
		g.Log().With(log.Fields{
			"State":    g.State.EventName(),
			"Username": seat.Username,
		}).Warn("engine/state: request failed")
	}
}
