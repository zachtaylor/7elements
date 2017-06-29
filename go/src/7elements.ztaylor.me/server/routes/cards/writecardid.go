package cards

import (
	"7elements.ztaylor.me"
	"7elements.ztaylor.me/log"
	"net/http"
	// "strconv"
)

func WriteCardId(id int, w http.ResponseWriter, lang string) {
	card := SE.Cards.Cache[id]
	if card == nil {
		w.WriteHeader(500)
		log.Add("CardId", id).Error("cards: card missing")
		return
	}

	texts := SE.CardTexts.Cache[lang]
	if texts == nil {
		w.Write([]byte("cards: language missing"))
		w.WriteHeader(500)
		log.Add("Language", lang).Error("cards: language missing")
		return
	} else if texts[id] == nil {
		w.Write([]byte("cards: language: card missing"))
		w.WriteHeader(500)
		log.Add("CardId", id).Add("Language", lang).Error("cards: language missing")
		return
	}

	j := MakeCardJson(card, texts[id])
	j.Write(w)
	log.Add("CardId", id).Debug("cards: serve")
}
