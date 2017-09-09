package cards

import (
	"7elements.ztaylor.me/cards"
	"ztaylor.me/json"
	"ztaylor.me/log"
	// "ztaylor.me/http/sessions"
	"net/http"
)

func WriteAllCards(w http.ResponseWriter, lang string) {
	j := json.Json{}

	texts := cards.TextsCache[lang]
	if texts == nil {
		w.Write([]byte("language missing"))
		w.WriteHeader(500)
		log.Add("Language", lang).Error("/api/cards: language missing")
		return
	}

	for cardid, card := range cards.CardCache {
		j[json.UItoS(uint(cardid))] = cards.Json(card, cards.BodyCache[cardid], texts[cardid])
	}

	j.Write(w)
}
