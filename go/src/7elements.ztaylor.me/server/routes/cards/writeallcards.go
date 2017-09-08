package cards

import (
	"7elements.ztaylor.me/cards"
	"ztaylor.me/json"
	"ztaylor.me/log"
	// "7elements.ztaylor.me/server/sessionman"
	"net/http"
)

func WriteAllCards(w http.ResponseWriter, lang string) {
	j := json.Json{}

	texts := cards.TextsCache[lang]
	if texts == nil {
		w.Write([]byte("cards: language missing"))
		w.WriteHeader(500)
		log.Add("Language", lang).Error("cards: language missing")
		return
	}

	for cardid, card := range cards.CardCache {
		j[json.UItoS(uint(cardid))] = MakeCardJson(card, cards.BodyCache[cardid], texts[cardid])
	}

	j.Write(w)
}
