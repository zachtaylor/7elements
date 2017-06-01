package myaccount

import (
	"7elements.ztaylor.me"
	"7elements.ztaylor.me/log"
	"7elements.ztaylor.me/server/json"
	"7elements.ztaylor.me/server/sessionman"
	"net/http"
	"time"
)

var Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log.Add("RemoteAddr", r.RemoteAddr)

	if r.Method != "GET" {
		w.WriteHeader(404)
		log.Add("Method", r.Method).Warn("myaccount: only GET allowed")
		return
	}

	session, err := sessionman.ReadRequestCookie(r)
	if err != nil {
		sessionman.EraseSessionId(w)
		w.WriteHeader(400)
		w.Write([]byte("session missing"))
		log.Add("Error", err).Error("myaccount: session missing")
		return
	}

	account := SE.Accounts.Cache[session.Username]
	if account == nil {
		sessionman.EraseSessionId(w)
		w.WriteHeader(500)
		log.Add("Error", err).Error("myaccount: account missing")
		return
	}

	cardcollection := SE.AccountsCards.Cache[account.Username]
	if cardcollection == nil {
		if cc, err := SE.AccountsCards.Get(account.Username); err != nil {
			sessionman.EraseSessionId(w)
			w.WriteHeader(500)
			log.Add("Error", err).Error("myaccount: collection")
			return
		} else {
			cardcollection = cc
			SE.AccountsCards.Cache[account.Username] = cardcollection
		}
	}

	j := json.Json{
		"username":       account.Username,
		"language":       account.Language,
		"email":          account.Email,
		"session-expire": session.Expire.Sub(time.Now()).String(),
	}
	cardCount := 0
	for _, list := range cardcollection {
		cardCount += len(list)
	}
	j["cards"] = cardCount

	j.Write(w)
})
