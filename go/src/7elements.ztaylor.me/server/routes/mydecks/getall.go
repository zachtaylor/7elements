package mydecks

import (
	"7elements.ztaylor.me"
	"7elements.ztaylor.me/log"
	"7elements.ztaylor.me/server/json"
	"7elements.ztaylor.me/server/sessionman"
	"7elements.ztaylor.me/server/util"
	"net/http"
)

func GetAll(w http.ResponseWriter, r *http.Request, account *SE.Account) {
	log.Add("Username", account.Username)

	accountdecks, err := serverutil.GetAccountsDecks(account.Username)
	if err != nil {
		sessionman.EraseSessionId(w)
		w.WriteHeader(500)
		log.Add("Error", err).Error("mydecks: accountdecks error")
		return
	}

	j := json.Json{}
	for deckid, deck := range accountdecks {
		deckidstring := json.I64toS(int64(deckid))
		j[deckidstring] = MakeDeckJson(deck)
	}

	j.Write(w)
	log.Debug("mydecks: success")
}
