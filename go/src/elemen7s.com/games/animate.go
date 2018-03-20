package games

import (
	"fmt"
	"ztaylor.me/js"
)

func AnimateAddCard(player Player, card *Card) {
	player.Send("animate", js.Object{
		"animate": "add card",
		"gcid":    card.Id,
		"cardid":  card.Card.Id,
	})
}

func AnimateHand(player Player, game *Game, cards Cards) {
	player.Send("hand", js.Object{
		"gameid": game.Id,
		"cards":  cards.Json(),
	})
}

func AnimatePass(player Player, game *Game, username string) {
	player.Send("pass", js.Object{
		"gameid":   game.Id,
		"target":   game.Active.Target,
		"username": username,
	})
}

func AnimateAttack(player Player, a AttackOptions) {
	player.Send("animate", js.Object{
		"animate":       "attack options",
		"attackoptions": a,
	})
}

func AnimateAddElement(game *Game, username string, e int) {
	game.Broadcast("animate", js.Object{
		"animate":  "add element",
		"username": username,
		"element":  e,
	})
}

func AnimateSpawn(game *Game, card *Card) {
	game.Broadcast("spawn", js.Object{
		"gameid":   game.Id,
		"username": card.Username,
		"card":     card.Json(),
	})
}

func AnimateAlertError(player Player, game *Game, source string, text string) {
	player.Send("alert", js.Object{
		"class":    "error",
		"gameid":   game.Id,
		"username": source,
		"message":  text,
	})
}

func BroadcastAnimateAlertError(game *Game, text string) {
	for _, s := range game.Seats {
		AnimateAlertError(s, game, fmt.Sprintf("Game#%d", game.Id), text)
	}
}

func BroadcastAnimateAlertChat(game *Game, username, message string) {
	game.Broadcast("alert", js.Object{
		"class":    "tip",
		"gameid":   game.Id,
		"username": username,
		"message":  message,
	})
}
