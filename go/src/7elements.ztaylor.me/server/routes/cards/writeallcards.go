package cards

import (
	"7elements.ztaylor.me/cards"
	"net/http"
	"ztaylor.me/json"
)

func WriteAllCards(w http.ResponseWriter, texts map[int]*cards.Texts) {
	j := json.Json{}

	for cardid, card := range cards.CardCache {
		j[json.UItoS(uint(cardid))] = cards.Json(card, cards.BodyCache[cardid], texts[cardid])
	}

	j.Write(w)
}
