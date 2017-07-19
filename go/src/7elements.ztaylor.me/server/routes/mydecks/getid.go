package mydecks

import (
	"7elements.ztaylor.me"
	"7elements.ztaylor.me/log"
	"7elements.ztaylor.me/server/json"
	"7elements.ztaylor.me/server/util"
	"net/http"
)

func GetId(w http.ResponseWriter, r *http.Request, deckid int, account *SE.Account) {
	log.Add("DeckId", deckid)

	accountdecks, err := serverutil.GetAccountsDecks(account.Username)
	if err != nil {
		w.WriteHeader(500)
		log.Add("Error", err).Error("mydecks: accountdecks error")
		w.Write([]byte("session missing"))
		return
	}

	if deck := accountdecks[deckid]; deck == nil {
		w.WriteHeader(500)
		log.Add("Error", err).Error("mydecks: deck missing: " + account.Username)
		w.Write([]byte("deck id not found"))
	} else {
		MakeDeckJson(deck).Write(w)
		log.Debug("mydecks: success")
	}
}

func MakeDeckJson(deck *SE.AccountDeck) json.Json {
	return json.Json{
		"id":    deck.Id,
		"name":  deck.Name,
		"cards": deck.Cards,
		"wins":  deck.Wins,
	}
}
