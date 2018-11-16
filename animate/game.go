package animate

import (
	"github.com/zachtaylor/7elements"
)

func GameState(game *vii.Game) {
	game.WriteJson(Build("/game/state", Json{
		"gameid": game.Key,
		"event":  game.State.EventName(),
		"timer":  int(game.State.Timer.Seconds()),
		"reacts": game.State.Reacts,
		"data":   game.State.Event.Json(game),
	}))
}

func GameSeat(game *vii.Game, seat *vii.GameSeat) {
	json := seat.Json(false)
	json["gameid"] = game.Key
	game.WriteJson(Build("/game/seat", json))
}

func GameReact(game *vii.Game, username string) {
	game.WriteJson(Build("/game/react", Json{
		"gameid":   game.Key,
		"event":    game.State.EventName(),
		"username": username,
		"react":    game.State.Reacts[username],
	}))
}

func GameCard(game *vii.Game, card *vii.GameCard) {
	game.WriteJson(Build("/game/react", Json{
		"gameid":   game.Key,
		"username": card.Username,
		"card":     card.Json(),
	}))
}

func GameHand(game *vii.Game, seat *vii.GameSeat) {
	seat.WriteJson(Build("/game/hand", Json{
		"gameid": game.Key,
		"cards":  seat.Hand.Json(),
	}))
}

func GameSpawn(game *vii.Game, card *vii.GameCard) {
	game.WriteJson(Build("/game/spawn", Json{
		"gameid":   game.Key,
		"username": card.Username,
		"card":     card.Json(),
	}))
}

func GameElement(game *vii.Game, username string, e int) {
	game.WriteJson(Build("/game/element", vii.Json{
		"gameid":   game.Key,
		"username": username,
		"element":  e,
	}))
}

func Choice(w vii.JsonWriter, game *vii.Game, prompt string, choices []Json, data Json) {
	w.WriteJson(Build("/game/choice", Json{
		"animate": "choice",
		"gameid":  game.Key,
		"prompt":  prompt,
		"choices": choices,
		"data":    data,
	}))
}

func GameError(w vii.JsonWriter, g *vii.Game, source, message string) {
	w.WriteJson(Build("/game/error", Json{
		"gameid":  g.Key,
		"source":  source,
		"message": message,
	}))
}
