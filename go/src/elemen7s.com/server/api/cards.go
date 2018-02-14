package api

import (
	"elemen7s.com/cards"
	"elemen7s.com/cards/texts"
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
	} else {
		text := texts.Get(r.Language, int(cardid))
		r.WriteJson(cards.JsonWithText(card, text))
	}
	return nil
}

func AllCardsJson(lang string) js.Object {
	j := js.Object{}
	texts := texts.GetAll(lang)
	for cardid, card := range cards.CardCache {
		if texts[cardid] == nil {
			log.Add("CardId", cardid).Error("/cards: card text missing")
		} else {
			j[fmt.Sprintf("%d", cardid)] = cards.JsonWithText(card, texts[cardid])
		}
	}
	return j
}
