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

func BroadcastAnimateSleep(game *Game, gcid int) {
	game.Broadcast("animate", js.Object{
		"animate": "sleep",
		"gameid":  game.Id,
		"gcid":    gcid,
	})
}

func BroadcastAnimateAddElement(game *Game, username string, e int) {
	game.Broadcast("animate", js.Object{
		"animate":  "add element",
		"gameid":   game.Id,
		"username": username,
		"element":  e,
	})
}

func BroadcastAnimateSpawn(game *Game, card *Card) {
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

func AnimateChoice(player Player, game *Game, prompt string, choices []js.Object, data js.Object) {
	player.Send("animate", js.Object{
		"animate": "choice",
		"gameid":  game.Id,
		"prompt":  prompt,
		"choices": choices,
		"data":    data,
	})
}

func AnimateNoviceSeerChoice(player Player, game *Game, card *Card) {
	prompt := fmt.Sprintf("Destroy %s?", card.CardText.Name)
	choices := []js.Object{
		js.Object{
			"choice":  "no",
			"display": "no",
		},
		js.Object{
			"choice":  "yes",
			"display": "yes",
		},
	}
	json := js.Object{
		"card": card.Json(),
	}
	AnimateChoice(player, game, prompt, choices, json)
}

var timeWalkerChoices = []js.Object{
	js.Object{"choice": "1", "display": `<img src="/img/icon/element-1.png">`},
	js.Object{"choice": "2", "display": `<img src="/img/icon/element-2.png">`},
	js.Object{"choice": "3", "display": `<img src="/img/icon/element-3.png">`},
	js.Object{"choice": "4", "display": `<img src="/img/icon/element-4.png">`},
	js.Object{"choice": "5", "display": `<img src="/img/icon/element-5.png">`},
	js.Object{"choice": "6", "display": `<img src="/img/icon/element-6.png">`},
	js.Object{"choice": "7", "display": `<img src="/img/icon/element-7.png">`},
}

func AnimateTimeWalkerChoice(player Player, game *Game) {
	AnimateChoice(player, game, "Create an Element", timeWalkerChoices, js.Object{})
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

func BroadcastAnimateMulligan(game *Game, username string) {
	game.Broadcast("animate", js.Object{
		"animate":  "mulligan",
		"gameid":   game.Id,
		"username": username,
	})
}
