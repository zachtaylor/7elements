// Package animate provides macros to send ui-control messages
package animate // import "github.com/zachtaylor/7tcg/animate"

import (
	"github.com/zachtaylor/7tcg"
	"ztaylor.me/js"
)

func Pass(w vii.Receiver, game *vii.Game, username string, target string) {
	w.Send("pass", js.Object{
		"gameid":   game,
		"username": username,
		"target":   target,
	})
}

func Chat(w vii.Receiver, game *vii.Game, username string, message string) {
	w.Send("chat", js.Object{
		"gameid":   game,
		"username": username,
		"message":  message,
	})
}

func Choice(w vii.Receiver, game *vii.Game, prompt string, choices []js.Object, data js.Object) {
	w.Send("animate", js.Object{
		"animate": "choice",
		"gameid":  game,
		"prompt":  prompt,
		"choices": choices,
		"data":    data,
	})
}

func Hand(game *vii.Game, seat *vii.GameSeat) {
	seat.Send("hand", js.Object{
		"gameid": game,
		"cards":  seat.Hand.Json(),
	})
}

func Spawn(game *vii.Game, card *vii.GameCard) {
	game.Send("spawn", js.Object{
		"gameid":   game,
		"username": card.Username,
		"card":     card.Json(),
	})
}

func Error(w vii.Receiver, game *vii.Game, source string, message string) {
	w.Send("alert", js.Object{
		"class":    "error",
		"gameid":   game,
		"username": source,
		"message":  message,
	})
}
