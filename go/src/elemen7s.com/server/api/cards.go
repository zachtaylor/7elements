package api

import (
	"elemen7s.com"
	"elemen7s.com/cards"
	"fmt"
	"strconv"
	"ztaylor.me/http"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func CardsHandler(r *http.Request) error {
	if r.Quest == "/cards.json" {
		r.WriteJson(AllCardsJson(r.Language))
	} else if len(r.Quest) < 13 {
		log.Error("/cards: card id unavailable")
	} else if cardid, err := strconv.Atoi(r.Quest[7 : len(r.Quest)-5]); err != nil {
		log.Add("Error", err).Error("/api/cards: parse card id")
	} else if card := cards.CardCache[cardid]; card == nil {
		log.Error("/cards: card missing")
	} else if text, err := vii.CardTextService.Get(r.Language, int(cardid)); err != nil {
		log.Add("Error", err).Error("/api/cards: card text service")
	} else {
		r.WriteJson(cards.JsonWithText(card, text))
	}
	return nil
}

func AllCardsJson(lang string) js.Object {
	j := js.Object{}
	for cardid, card := range cards.CardCache {
		if text, err := vii.CardTextService.Get(lang, cardid); text == nil {
			log.Add("Error", err).Add("CardId", cardid).Error("/api/cards: card text service")
		} else {
			j[fmt.Sprintf("%d", cardid)] = cards.JsonWithText(card, text)
		}
	}
	return j
}
