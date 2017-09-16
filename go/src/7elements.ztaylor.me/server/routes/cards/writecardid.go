package cards

import (
	"7elements.ztaylor.me/cards"
	"net/http"
	"ztaylor.me/log"
	// "strconv"
)

func WriteCardId(cardid int, w http.ResponseWriter, texts map[int]*cards.Texts) {
	log := log.Add("CardId", cardid)

	if card := cards.CardCache[cardid]; card == nil {
		w.WriteHeader(500)
		log.Error("/api/cards: card missing")
	} else if texts[cardid] == nil {
		w.WriteHeader(500)
		w.Write([]byte("api/cards: card text missing"))
		log.Add("CardId", cardid).Error("/api/cards: card text missing")
	} else {
		j := cards.Json(card, cards.BodyCache[cardid], texts[cardid])
		j.Write(w)
	}
}
