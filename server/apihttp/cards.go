package apihttp

// 	"strconv"

// 	"ztaylor.me/http"
// 	"ztaylor.me/js"
// 	"ztaylor.me/log"

// func CardsHandler(r *http.Quest) error {
// 	if r.Quest == "/api/cards.json" {
// 		r.Write([]byte(js.String(AllCardsJson())))
// 	} else if len(r.Quest) < 17 {
// 		log.Error("/cards: card id unavailable")
// 	} else if cardid, err := strconv.Atoi(r.Quest[11 : len(r.Quest)-5]); err != nil {
// 		log.Add("Error", err).Error("/api/cards: parse card id")
// 	} else if card, err := vii.CardService.Get(cardid); card == nil {
// 		log.Add("Error", err).Error("/api/cards: card missing")
// 	} else {
// 		log.Add("Card", card.Id).Info("/api/cards")
// 		r.WriteJson(card.Json())
// 	}
// 	return nil
// }
