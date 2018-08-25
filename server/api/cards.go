package api

import (
	"fmt"
	"strconv"

	"github.com/zachtaylor/7elements"
	"ztaylor.me/http"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func CardsHandler(r *http.Request) error {
	if r.Quest == "/api/cards.json" {
		r.WriteJson(AllCardsJson(r.Language))
	} else if len(r.Quest) < 17 {
		log.Error("/cards: card id unavailable")
	} else if cardid, err := strconv.Atoi(r.Quest[11 : len(r.Quest)-5]); err != nil {
		log.Add("Error", err).Error("/api/cards: parse card id")
	} else if card, err := vii.CardService.GetCard(cardid); card == nil {
		log.Add("Error", err).Error("/api/cards: card missing")
	} else if text, err := vii.CardTextService.GetCardText(r.Language, int(cardid)); err != nil {
		log.Add("Error", err).Error("/api/cards: card text service")
	} else {
		log.Add("Card", card.Id).Info("/api/cards")
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
