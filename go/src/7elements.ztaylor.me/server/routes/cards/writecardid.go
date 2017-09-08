package cards

import (
	"7elements.ztaylor.me/cards"
	"net/http"
	"ztaylor.me/log"
	// "strconv"
)

func WriteCardId(cardid int, w http.ResponseWriter, lang string) {
	card := cards.CardCache[cardid]
	if card == nil {
		w.WriteHeader(500)
		log.Add("CardId", cardid).Error("cards: card missing")
		return
	}

	texts := cards.TextsCache[lang]
	if texts == nil {
		w.Write([]byte("cards: language missing"))
		w.WriteHeader(500)
		log.Add("Language", lang).Error("cards: language missing")
		return
	} else if texts[cardid] == nil {
		w.Write([]byte("cards: language: card missing"))
		w.WriteHeader(500)
		log.Add("CardId", cardid).Add("Language", lang).Error("cards: language missing")
		return
	}

	j := MakeCardJson(card, cards.BodyCache[cardid], texts[cardid])
	j.Write(w)
	log.Add("CardId", cardid).Debug("cards: serve")
}
