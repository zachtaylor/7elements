package api

import (
	"strconv"

	"github.com/zachtaylor/7elements"
	"ztaylor.me/http"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func CardsHandler(r *http.Request) error {
	if r.Quest == "/api/cards.json" {
		r.Write([]byte(js.String(AllCardsJson())))
	} else if len(r.Quest) < 17 {
		log.Error("/cards: card id unavailable")
	} else if cardid, err := strconv.Atoi(r.Quest[11 : len(r.Quest)-5]); err != nil {
		log.Add("Error", err).Error("/api/cards: parse card id")
	} else if card, err := vii.CardService.Get(cardid); card == nil {
		log.Add("Error", err).Error("/api/cards: card missing")
	} else {
		log.Add("Card", card.Id).Info("/api/cards")
		r.WriteJson(card.Json())
	}
	return nil
}

func AllCardsJson() []js.Object {
	j := make([]js.Object, 0)
	for _, card := range vii.CardService.GetAll() {
		j = append(j, card.Json())
	}
	return j
}
