package api

import (
	"elemen7s.com"
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
	} else if card, _ := vii.CardService.GetCard(cardid); card == nil {
		log.Error("/cards: card missing")
	} else if text, err := vii.CardTextService.GetCardText(r.Language, int(cardid)); err != nil {
		log.Add("Error", err).Error("/api/cards: card text service")
	} else {
		r.WriteJson(card.JsonWithText(text))
	}
	return nil
}

func AllCardsJson(lang string) js.Object {
	j := js.Object{}
	for cardid, card := range vii.CardService.GetAllCards() {
		if text, err := vii.CardTextService.GetCardText(lang, cardid); text == nil {
			log.Add("Error", err).Add("CardId", cardid).Error("/api/cards: card text service")
		} else {
			j[fmt.Sprintf("%d", cardid)] = card.JsonWithText(text)
		}
	}
	return j
}
