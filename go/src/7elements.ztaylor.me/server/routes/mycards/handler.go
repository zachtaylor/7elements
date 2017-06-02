package mycards

import (
	"7elements.ztaylor.me"
	"7elements.ztaylor.me/log"
	"7elements.ztaylor.me/server/json"
	"7elements.ztaylor.me/server/sessionman"
	"7elements.ztaylor.me/server/util"
	"net/http"
)

var Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log.Add("RemoteAddr", r.RemoteAddr)

	if r.Method != "GET" {
		w.WriteHeader(404)
		log.Add("Method", r.Method).Warn("mycards: only GET allowed")
		return
	}

	session, err := sessionman.ReadRequestCookie(r)
	if session == nil {
		if err != nil {
			sessionman.EraseSessionId(w)
			log.Add("Error", err)
		}
		w.WriteHeader(400)
		w.Write([]byte("session missing"))
		log.Error("mycards: session missing")
		return
	}

	account := SE.Accounts.Cache[session.Username]
	if account == nil {
		sessionman.EraseSessionId(w)
		w.WriteHeader(500)
		log.Add("Error", err).Error("mycards: account missing")
		return
	}

	accountcards, err := serverutil.GetAccountsCards(account.Username)
	if err != nil {
		sessionman.EraseSessionId(w)
		w.WriteHeader(500)
		log.Add("Error", err).Error("mycards: accountcards error")
		return
	}

	j := json.Json{}
	for cardId, list := range accountcards {
		j[json.UItoS(cardId)] = len(list)
	}

	j.Write(w)
})
