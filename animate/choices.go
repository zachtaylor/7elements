package animate

import (
	"fmt"

	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/game"
)

var newElementChoices = []Json{
	Json{"choice": "1", "display": `<img src="/img/icon/element-1.png">`},
	Json{"choice": "2", "display": `<img src="/img/icon/element-2.png">`},
	Json{"choice": "3", "display": `<img src="/img/icon/element-3.png">`},
	Json{"choice": "4", "display": `<img src="/img/icon/element-4.png">`},
	Json{"choice": "5", "display": `<img src="/img/icon/element-5.png">`},
	Json{"choice": "6", "display": `<img src="/img/icon/element-6.png">`},
	Json{"choice": "7", "display": `<img src="/img/icon/element-7.png">`},
}

func NewElementChoice(w JsonWriter, game *game.T) {
	Choice(w, game, "Create an Element", newElementChoices, Json{})
}

func NoviceSeerChoice(w JsonWriter, game *game.T, card *game.Card) {
	prompt := fmt.Sprintf("Destroy %s?", card.Card.Name)
	choices := []Json{
		Json{
			"choice":  "no",
			"display": "no",
		},
		Json{
			"choice":  "yes",
			"display": "yes",
		},
	}
	json := Json{
		"card": card.Json(),
	}
	Choice(w, game, prompt, choices, json)
}

func GraveBirth(w JsonWriter, game *game.T) {
	prompt := "Create a <b>Body</b> from any player's <b>Past</b>"
	cards := []Json{}
	choices := []Json{}
	for _, seat := range game.Seats {
		for _, card := range seat.Past {
			if card.Card.Type != vii.CTYPbody {
				continue
			}

			cards = append(cards, card.Json())
			choices = append(choices, Json{
				"choice":  card.Id,
				"display": card.Card.Name,
			})
		}
	}

	json := Json{
		"cards": cards,
	}
	Choice(w, game, prompt, choices, json)
}
