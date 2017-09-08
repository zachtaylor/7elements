package api

import (
	"7elements.ztaylor.me/accountscards"
	"7elements.ztaylor.me/server/sessionman"
	"net/http"
	"ztaylor.me/json"
	"ztaylor.me/log"
)

var MyCardsJsonHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log := log.Add("RemoteAddr", r.RemoteAddr)

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

	log = log.Add("Username", session.Username)
	accountcards, err := accountscards.Get(session.Username)
	if err != nil {
		sessionman.EraseSessionId(w)
		w.WriteHeader(500)
		log.Add("Error", err).Error("mycards: accountcards error")
		return
	}

	j := json.Json{}
	for cardId, list := range accountcards {
		j[json.UItoS(uint(cardId))] = len(list)
	}

	j.Write(w)
	log.Debug("mycards: success")
})
