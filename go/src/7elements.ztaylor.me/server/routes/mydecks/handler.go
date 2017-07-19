package mydecks

import (
	"7elements.ztaylor.me"
	"7elements.ztaylor.me/log"
	"7elements.ztaylor.me/server/sessionman"
	"net/http"
	"strconv"
)

var Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log.Add("RemoteAddr", r.RemoteAddr)

	session, err := sessionman.ReadRequestCookie(r)
	if session == nil {
		if err != nil {
			sessionman.EraseSessionId(w)
			log.Add("Error", err)
		}
		w.WriteHeader(400)
		w.Write([]byte("session missing"))
		log.Error("mydecks: session missing")
		return
	}

	account := SE.Accounts.Cache[session.Username]
	if account == nil {
		sessionman.EraseSessionId(w)
		w.WriteHeader(500)
		log.Add("Error", err).Error("mydecks: account missing")
		return
	}

	var deckid = 0

	if r.RequestURI == "/api/mydecks.json" {
	} else if len(r.RequestURI) < 19 {
		w.WriteHeader(500)
		log.Error("mydecks: deck id unavailable")
		return
	} else if deckidParse, err := strconv.Atoi(r.RequestURI[13 : len(r.RequestURI)-5]); err == nil {
		deckid = deckidParse
	} else {
		w.WriteHeader(500)
		log.Add("Error", err).Error("mydecks: parse deck id")
		return
	}

	if r.Method == "GET" {
		if deckid == 0 {
			GetAll(w, r, account)
		} else {
			GetId(w, r, deckid, account)
		}
	} else if r.Method == "POST" {
		Post(w, r, deckid, account)
	}
})
