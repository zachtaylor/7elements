package animate

import (
	"fmt"

	"github.com/zachtaylor/7elements"
	"ztaylor.me/js"
)

var newElementChoices = []js.Object{
	js.Object{"choice": "1", "display": `<img src="/img/icon/element-1.png">`},
	js.Object{"choice": "2", "display": `<img src="/img/icon/element-2.png">`},
	js.Object{"choice": "3", "display": `<img src="/img/icon/element-3.png">`},
	js.Object{"choice": "4", "display": `<img src="/img/icon/element-4.png">`},
	js.Object{"choice": "5", "display": `<img src="/img/icon/element-5.png">`},
	js.Object{"choice": "6", "display": `<img src="/img/icon/element-6.png">`},
	js.Object{"choice": "7", "display": `<img src="/img/icon/element-7.png">`},
}

func NewElementChoice(w vii.Receiver, game *vii.Game) {
	Choice(w, game, "Create an Element", newElementChoices, js.Object{})
}

func NoviceSeerChoice(w vii.Receiver, game *vii.Game, card *vii.GameCard) {
	prompt := fmt.Sprintf("Destroy %s?", card.Name)
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
	Choice(w, game, prompt, choices, json)
}

func GraveBirth(w vii.Receiver, game *vii.Game) {
	prompt := "Create a <b>Body</b> from any player's <b>Past</b>"
	cards := []js.Object{}
	choices := []js.Object{}
	for _, seat := range game.Seats {
		for _, card := range seat.Graveyard {
			if card.Card.Type != vii.CTYPbody {
				continue
			}

			cards = append(cards, card.Json())
			choices = append(choices, js.Object{
				"choice":  card.Id,
				"display": card.Name,
			})
		}
	}

	json := js.Object{
		"cards": cards,
	}
	Choice(w, game, prompt, choices, json)
}
