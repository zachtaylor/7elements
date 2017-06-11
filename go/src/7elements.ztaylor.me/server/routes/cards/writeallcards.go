package cards

import (
	"7elements.ztaylor.me"
	"7elements.ztaylor.me/log"
	"7elements.ztaylor.me/server/json"
	// "7elements.ztaylor.me/server/sessionman"
	"net/http"
)

func WriteAllCards(w http.ResponseWriter, lang string) {
	j := json.Json{}

	texts := SE.CardTexts.Cache[lang]
	if texts == nil {
		w.Write([]byte("cards: language missing"))
		w.WriteHeader(500)
		log.Add("Language", lang).Error("cards: language missing")
		return
	}

	for cardid, card := range SE.Cards.Cache {
		j[json.UItoS(cardid)] = MakeCardJson(card, texts[cardid])
	}

	j.Write(w)
}
