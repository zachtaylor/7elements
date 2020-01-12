package engine

import (
	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/game"
	"github.com/zachtaylor/7elements/game/state"
	"github.com/zachtaylor/7elements/game/update"
	"ztaylor.me/cast"
)

func Request(g *game.T, seat *game.Seat, uri string, json cast.JSON) []game.Stater {
	if g.State.R.Name() == "main" && g.State.R.Seat() == seat.Username {
		return RequestAny(g, seat, uri, json)
	} else {
		return RequestResponse(g, seat, uri, json)
	}
}

func requestConnect(g *game.T, s *game.Seat) {
	g.Log().With(cast.JSON{
		"Username": s.Username,
		"State":    g.State.Print(),
	}).Debug("engine/connect: seated")
	update.Connect(g, s)
	if connector, ok := g.State.R.(game.ConnectStater); ok {
		connector.OnConnect(g, s)
	}
}

func requestDisconnect(g *game.T, s *game.Seat) {
	g.Log().With(cast.JSON{
		"Username": s.Username,
		"State":    g.State,
	}).Tag("engine/disconnect: left")

	if disconnector, ok := g.State.R.(game.DisconnectStater); ok {
		disconnector.OnDisconnect(g, s)
	}
}

func requestChat(g *game.T, seat *game.Seat, json cast.JSON) {
	text := json.GetS("text")
	g.Log().With(cast.JSON{
		"Username": seat.Username,
		"Text":     text,
	}).Debug("engine/chat") // died after
	go g.GetChat().AddMessage(chat.NewMessage(seat.Username, text))
}

func requestPass(g *game.T, seat *game.Seat, json cast.JSON) {
	log := g.Log().With(cast.JSON{
		"State":    g.State.Print(),
		"Username": seat.Username,
	}).Tag("engine/pass")
	if pass := json.GetS("pass"); pass == "" {
		log.Warn("target missing")
	} else if pass != g.State.ID() {
		log.Add("PassID", pass).Warn("target mismatch")
	} else if len(g.State.Reacts[seat.Username]) > 0 {
		update.ErrorW(seat, "pass", "response already recorded")
	} else {
		g.State.Reacts[seat.Username] = "pass"
		update.React(g, seat.Username)
	}
}

// RequestAttack causes AttackEvent to stack on MainEvent
func RequestAttack(g *game.T, seat *game.Seat, json cast.JSON) []game.Stater {
	log := g.Log().With(cast.JSON{
		"Seat": seat.Print(),
	}).Tag("engine/attack")

	if id := json.GetS("id"); id == "" {
		log.Error("id missing")
	} else if token := seat.Present[id]; token == nil {
		log.Add("ID", id).Error("id invalid")
		update.ErrorW(seat, id, `not in your present`)
	} else if token.Body == nil {
		log.Add("Token", token.String()).Error("card type must be body")
		update.ErrorW(seat, token.Card.Card.Name, `not "body" type`)
	} else if !token.IsAwake {
		log.Add("Token", token.String()).Error("card must be awake")
		update.ErrorW(seat, token.Card.Card.Name, `not awake`)
	} else {
		log.Add("Token", token.String()).Info("accept")
		token.IsAwake = false
		update.Token(g, token)
		return []game.Stater{state.NewAttack(seat.Username, token)}
	}
	return nil
}

func RequestPlay(g *game.T, seat *game.Seat, json cast.JSON, onlySpells bool) []game.Stater {
	log := g.Log().With(cast.JSON{
		"Seat": seat.Print(),
	}).Tag("engine/play")

	if id := json.GetS("id"); id == "" {
		log.Error("no id")
	} else if card := seat.Hand[id]; card == nil {
		log.Error("no card")
		update.ErrorW(seat, `vii`, `bad card id`)
	} else if card.Card.Type != vii.CTYPspell && onlySpells {
		log.Add("Card", card.Print()).Error("card type must be spell")
		update.ErrorW(seat, card.Card.Name, `not "spell" type`)
	} else if !seat.Elements.GetActive().Test(card.Card.Costs) {
		log.Add("Card", card.Print()).Error("not enough elements")
		update.ErrorW(seat, card.Card.Name, `not enough elements`)
	} else {
		log.Add("Card", card.Print()).Info("accept")
		seat.Elements.Deactivate(card.Card.Costs)
		delete(seat.Hand, id)
		update.Seat(g, seat)
		update.Hand(seat)
		return []game.Stater{state.NewPlay(seat.Username, card, json["target"])}
	}
	return nil
}

func RequestGameState(g *game.T, seat *game.Seat, json cast.JSON) {
	if requester, ok := g.State.R.(game.RequestStater); ok {
		requester.Request(g, seat, json)
	} else {
		g.Log().With(cast.JSON{
			"State":    g.State.Name(),
			"Username": seat.Username,
		}).Warn("engine/state: request failed")
	}
}
