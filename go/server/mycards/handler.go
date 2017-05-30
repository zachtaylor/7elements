package mycards

import (
	"7elements.ztaylor.me"
	"7elements.ztaylor.me/log"
	"7elements.ztaylor.me/server/json"
	"7elements.ztaylor.me/server/sessionman"
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
	if err != nil {
		sessionman.EraseSessionId(w)
		w.WriteHeader(400)
		w.Write([]byte("session missing"))
		log.Add("Error", err).Error("mycards: session missing")
		return
	}

	account := SE.Accounts.Cache[session.Username]
	if account == nil {
		sessionman.EraseSessionId(w)
		w.WriteHeader(500)
		log.Add("Error", err).Error("mycards: account missing")
		return
	}

	cardcollection := SE.AccountsCards.Cache[account.Username]
	if cardcollection == nil {
		if cc, err := SE.AccountsCards.Get(account.Username); err != nil {
			sessionman.EraseSessionId(w)
			w.WriteHeader(500)
			log.Add("Error", err).Error("mycards: collection")
			return
		} else {
			cardcollection = cc
			SE.AccountsCards.Cache[account.Username] = cardcollection
		}
	}

	j := json.Json{}
	for cardId, list := range cardcollection {
		j[json.UItoS(cardId)] = len(list)
	}

	j.Write(w)
})
