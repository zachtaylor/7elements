package api

import (
	"7elements.ztaylor.me/decks"
	"net/http"
	"ztaylor.me/http/sessions"
	"ztaylor.me/log"
)

func DecksJsonHandler(w http.ResponseWriter, r *http.Request) {
	log := log.Add("RemoteAddr", r.RemoteAddr)
	session, err := sessions.ReadRequestCookie(r)
	if session == nil {
		if err != nil {
			sessions.EraseSessionId(w)
			log.Add("Error", err)
		}
		w.WriteHeader(400)
		w.Write([]byte("session missing"))
		log.Error("decks.json: session missing")
		return
	}

	log.Add("Username", session.Username)
	decks, err := decks.Get(session.Username)
	if err != nil {
		sessions.EraseSessionId(w)
		w.WriteHeader(500)
		log.Add("Error", err).Error("decks.json")
		return
	}

	decks.Json().Write(w)
	log.Debug("decks.json")
}
