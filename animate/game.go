package animate

import (
	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/game"
)

func GameState(game *game.T) {
	game.WriteJson(Build("/game/state", game.State.Json(game)))
}

func GameReconnect(game *game.T, seat *game.Seat) {
	seat.WriteJson(Build("/game", game.Json(seat.Username)))
}

func GameSeat(game *game.T, seat *game.Seat) {
	json := seat.Json(false)
	json["gameid"] = game.ID()
	game.WriteJson(Build("/game/seat", json))
}

func GameReact(game *game.T, username string) {
	game.WriteJson(Build("/game/react", Json{
		"gameid":   game.ID(),
		"stateid":  game.State.ID(),
		"event":    game.State.EventName(),
		"username": username,
		"react":    game.State.Reacts[username],
	}))
}

func GameCard(game *game.T, card *game.Card) {
	game.WriteJson(Build("/game/card", Json{
		"gameid":   game.ID(),
		"username": card.Username,
		"card":     card.Json(),
	}))
}

func GameHand(game *game.T, seat *game.Seat) {
	seat.WriteJson(Build("/game/hand", Json{
		"gameid": game.ID(),
		"cards":  seat.Hand.Json(),
	}))
}

func GameSpawn(game *game.T, card *game.Card) {
	game.WriteJson(Build("/game/spawn", Json{
		"gameid":   game.ID(),
		"username": card.Username,
		"card":     card.Json(),
	}))
}

func GameElement(game *game.T, username string, e int) {
	game.WriteJson(Build("/game/element", vii.Json{
		"gameid":   game.ID(),
		"username": username,
		"element":  e,
	}))
}

func Choice(w vii.JsonWriter, game *game.T, prompt string, choices []Json, data Json) {
	w.WriteJson(Build("/game/choice", Json{
		"animate": "choice",
		"gameid":  game.ID(),
		"prompt":  prompt,
		"choices": choices,
		"data":    data,
	}))
}

func GameError(w vii.JsonWriter, g *game.T, source, message string) {
	w.WriteJson(Build("/game/error", Json{
		"gameid":  g.ID(),
		"source":  source,
		"message": message,
	}))
}
