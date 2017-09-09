package cards

import (
	"7elements.ztaylor.me/cards"
	"net/http"
	"ztaylor.me/log"
	// "strconv"
)

func WriteCardId(cardid int, w http.ResponseWriter, lang string) {
	log := log.Add("CardId", cardid)

	if card := cards.CardCache[cardid]; card == nil {
		w.WriteHeader(500)
		log.Error("/api/cards: card missing")
	} else if texts := cards.TextsCache[lang]; texts == nil {
		w.WriteHeader(500)
		w.Write([]byte("language missing"))
		log.Add("Language", lang).Error("/api/cards: language missing")
	} else if texts[cardid] == nil {
		w.Write([]byte("cards: language: card missing"))
		w.WriteHeader(500)
		log.Add("CardId", cardid).Add("Language", lang).Error("/api/cards: language missing")
	} else {
		j := cards.Json(card, cards.BodyCache[cardid], texts[cardid])
		j.Write(w)
		log.Add("Language", lang).Error("/api/cards: language missing")
	}
}
