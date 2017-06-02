package myaccount

import (
	"7elements.ztaylor.me"
	"7elements.ztaylor.me/log"
	"7elements.ztaylor.me/server/json"
	"7elements.ztaylor.me/server/sessionman"
	"7elements.ztaylor.me/server/util"
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
	if session == nil {
		if err != nil {
			log.Add("Error", err)
			sessionman.EraseSessionId(w)
		}
		w.WriteHeader(400)
		w.Write([]byte("session missing"))
		log.Error("myaccount: session missing")
		return
	}

	account := SE.Accounts.Cache[session.Username]
	if account == nil {
		sessionman.EraseSessionId(w)
		w.WriteHeader(500)
		log.Add("Error", err).Error("myaccount: account missing")
		return
	}

	accountcards, err := serverutil.GetAccountsCards(account.Username)
	if err != nil {
		sessionman.EraseSessionId(w)
		w.WriteHeader(500)
		log.Add("Error", err).Error("myaccount: collection")
		return
	}

	accountpacks, err := serverutil.GetAccountsPacks(account.Username)
	if err != nil {
		sessionman.EraseSessionId(w)
		w.WriteHeader(500)
		log.Add("Error", err).Error("myaccount: packs")
		return
	}

	j := json.Json{
		"username":       account.Username,
		"language":       account.Language,
		"email":          account.Email,
		"session-expire": session.Expire.Sub(time.Now()).String(),
		"coins":          account.Coins,
		"packs":          len(accountpacks),
	}
	cardCount := 0
	for _, list := range accountcards {
		cardCount += len(list)
	}
	j["cards"] = cardCount

	j.Write(w)
})
