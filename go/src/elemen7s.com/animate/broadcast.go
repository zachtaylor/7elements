package animate

import (
	"elemen7s.com"
	"ztaylor.me/js"
)

func BroadcastCardUpdate(game *vii.Game, card *vii.GameCard) {
	game.Send("animate", js.Object{
		"animate": "cardupdate",
		"gameid":  game,
		"data":    card.Json(),
	})
}

func BroadcastSeatUpdate(game *vii.Game, seat *vii.GameSeat) {
	game.Send("animate", js.Object{
		"animate": "seatupdate",
		"gameid":  game,
		"data":    seat.Json(false),
	})
}

func BroadcastAddElement(game *vii.Game, username string, e int) {
	game.Send("animate", js.Object{
		"animate":  "add element",
		"gameid":   game,
		"username": username,
		"element":  e,
	})
}

func BroadcastAlertError(game *vii.Game, text string) {
	game.Send("alert", js.Object{
		"gameid":   game,
		"class":    "error",
		"username": "Game#" + game.Key,
		"message":  text,
	})
}

func BroadcastAlertTip(game *vii.Game, username string, text string) {
	game.Send("alert", js.Object{
		"gameid":   game,
		"class":    "tip",
		"username": username,
		"message":  text,
		"timer":    1000,
	})
}

func BroadcastChat(game *vii.Game, username, text string) {
	Chat(game, game, username, text)
}
