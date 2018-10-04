package animate

import (
	"github.com/zachtaylor/7elements"
	"ztaylor.me/js"
)

func Pass(w vii.JsonWriter, game *vii.Game, username string, target string) {
	w.WriteJson(js.Object{
		"uri": "pass",
		"data": js.Object{
			"gameid":   game,
			"username": username,
			"target":   target,
		},
	})
}

func Choice(w vii.JsonWriter, game *vii.Game, prompt string, choices []js.Object, data js.Object) {
	w.WriteJson(js.Object{
		"uri": "animate",
		"data": js.Object{
			"animate": "choice",
			"gameid":  game,
			"prompt":  prompt,
			"choices": choices,
			"data":    data,
		},
	})
}

func Hand(game *vii.Game, seat *vii.GameSeat) {
	seat.WriteJson(Build("/game/hand", js.Object{
		"gameid": game,
		"cards":  seat.Hand.Json(),
	}))
}

func Spawn(game *vii.Game, card *vii.GameCard) {
	game.WriteJson(Build("/game/spawn", js.Object{
		"gameid":   game,
		"username": card.Username,
		"card":     card.Json(),
	}))
}

func GameError(w vii.JsonWriter, g *vii.Game, source, message string) {
	w.WriteJson(js.Object{
		"uri": "/game/error",
		"data": js.Object{
			"gameid":  g.Key,
			"source":  source,
			"message": message,
		},
	})
}
