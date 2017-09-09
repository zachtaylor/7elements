package api

import (
	"7elements.ztaylor.me/accounts"
	"7elements.ztaylor.me/accountscards"
	"net/http"
	"time"
	"ztaylor.me/http/sessions"
	"ztaylor.me/json"
	"ztaylor.me/log"
)

var MyAccountJsonHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log.Add("RemoteAddr", r.RemoteAddr)

	if r.Method != "GET" {
		w.WriteHeader(404)
		log.Add("Method", r.Method).Warn("myaccount: only GET allowed")
		return
	}

	session, err := sessions.ReadRequestCookie(r)
	if session == nil {
		sessions.EraseSessionId(w)
		w.WriteHeader(400)
		w.Write([]byte("session missing"))
		log.Add("Error", err).Error("myaccount: session missing")
		return
	}

	account, err := accounts.Get(session.Username)
	if err != nil {
		sessions.EraseSessionId(w)
		w.WriteHeader(500)
		log.Add("Error", err).Error("myaccount: collection")
		return
	}

	accountcards, err := accountscards.Get(session.Username)
	if err != nil {
		sessions.EraseSessionId(w)
		w.WriteHeader(500)
		log.Add("Error", err).Error("myaccount: collection")
		return
	}

	json.Json{
		"username":     account.Username,
		"language":     account.Language,
		"email":        account.Email,
		"session-life": session.Expire.Sub(time.Now()).String(),
		"coins":        account.Coins,
		"packs":        account.Packs,
		"cards":        accountcards.Json(),
	}.Write(w)
	log.Info("myaccount")
})
