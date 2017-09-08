package games

import (
	"7elements.ztaylor.me/elements"
	"7elements.ztaylor.me/games/cards"
	"7elements.ztaylor.me/server/sessionman"
	"ztaylor.me/json"
	"ztaylor.me/log"
)

func Element(game *Game, socket *sessionman.Socket, data json.Json, log log.Log) {
	if game.GamePhase != GPHSplay {
		log.Add("GamePhase", game.GamePhase).Warn("element: rejected")
	} else if game.TurnPhase != TPHSbegin {
		log.Add("TurnPhase", game.TurnPhase).Warn("element: rejected")
	} else if cast, ok := data["elementid"].(float64); !ok {
		log.Warn("element: elementid error")
	} else if elementid := int(cast); elementid < 0 || elementid > 7 {
		log.Warn("element: elementid rejected")
		socket.Send("turn", MakeTurnJson(game, game.CurrentTurn()))
	} else if seat := game.GetSeat(socket.Username); seat == nil {
		log.Error("element: socket has no seat in game")
		socket.Send("error", nil)
	} else {
		element := elements.Elements[elementid]
		game.CurrentTurn().Element = element
		seat.Elements.Append(element)
		card := seat.Deck.Draw()
		seat.Hand = append(seat.Hand, card)

		socket.Send("animate", json.Json{
			"animate": "draw card",
			"card":    card.Json(),
			"hand":    gamecards.Stack(seat.Hand).Json(),
		})

		SendAllSeats(game, "elements", json.Json{
			"username": seat.Username,
			"elements": seat.Elements,
		})

		log.Debug("element")
		game.Context.Cancel()
	}
}
