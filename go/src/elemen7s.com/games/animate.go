package games

import (
	"elemen7s.com"
	"fmt"
	"ztaylor.me/js"
)

func AnimateHand(game *Game, seat *Seat) {
	seat.Send("hand", js.Object{
		"gameid": game.Id,
		"cards":  seat.Hand.Json(),
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

func BroadcastAnimateCardUpdate(game *Game, card *vii.GameCard) {
	game.Broadcast("animate", js.Object{
		"animate": "cardupdate",
		"gameid":  game.Id,
		"data":    card.Json(),
	})
}

func BroadcastAnimateSeatUpdate(game *Game, seat *Seat) {
	game.Broadcast("animate", js.Object{
		"animate": "seatupdate",
		"gameid":  game.Id,
		"data":    seat.Json(false),
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

func BroadcastAnimateSpawn(game *Game, card *vii.GameCard) {
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

func AnimateNoviceSeerChoice(player Player, game *Game, card *vii.GameCard) {
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

func AnimateGraveBirthChoice(player Player, game *Game) {
	prompt := "Create a <b>Body</b> from any player's <b>Past</b>"
	cards := []js.Object{}
	choices := []js.Object{}
	for _, seat := range game.Seats {
		for _, card := range seat.Graveyard {
			if card.Card.CardType != vii.CTYPbody {
				continue
			}

			cards = append(cards, card.Json())
			choices = append(choices, js.Object{
				"choice":  card.Id,
				"display": card.CardText.Name,
			})
		}
	}

	json := js.Object{
		"cards": cards,
	}
	AnimateChoice(player, game, prompt, choices, json)
}

var newElementChoices = []js.Object{
	js.Object{"choice": "1", "display": `<img src="/img/icon/element-1.png">`},
	js.Object{"choice": "2", "display": `<img src="/img/icon/element-2.png">`},
	js.Object{"choice": "3", "display": `<img src="/img/icon/element-3.png">`},
	js.Object{"choice": "4", "display": `<img src="/img/icon/element-4.png">`},
	js.Object{"choice": "5", "display": `<img src="/img/icon/element-5.png">`},
	js.Object{"choice": "6", "display": `<img src="/img/icon/element-6.png">`},
	js.Object{"choice": "7", "display": `<img src="/img/icon/element-7.png">`},
}

func AnimateNewElementChoice(player Player, game *Game) {
	AnimateChoice(player, game, "Create an Element", newElementChoices, js.Object{})
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
