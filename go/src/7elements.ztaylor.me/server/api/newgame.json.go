package api

import (
	"7elements.ztaylor.me/accounts"
	"7elements.ztaylor.me/decks"
	"7elements.ztaylor.me/queue"
	"net/http"
	"strconv"
	"ztaylor.me/http/sessions"
	"ztaylor.me/json"
	"ztaylor.me/log"
)

var NewGameJsonHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log := log.Add("RemoteAddr", r.RemoteAddr)

	if r.Method != "GET" {
		w.WriteHeader(404)
		log.Add("Method", r.Method).Warn("newgame.json: only GET allowed")
		return
	}

	session, err := sessions.ReadRequestCookie(r)
	if session == nil {
		if err != nil {
			sessions.EraseSessionId(w)
			log.Add("Error", err)
		}
		w.WriteHeader(401)
		w.Write([]byte("session missing"))
		log.Error("newgame.json: session missing")
		return
	}

	log.Add("Username", session.Username)
	account := accounts.Test(session.Username)
	if account == nil {
		sessions.EraseSessionId(w)
		w.WriteHeader(500)
		w.Write([]byte("account error"))
		log.Error("newgame.json: account missing")
		return
	}

	mydecks, err := decks.Get(session.Username)
	if mydecks == nil {
		w.WriteHeader(500)
		w.Write([]byte("decks error"))
		log.Add("Error", err).Error("newgame.json: decks missing")
		return
	}

	var deck *decks.Deck

	if parse, err := strconv.ParseInt(r.FormValue("deckid"), 10, 0); err != nil {
		w.WriteHeader(500)
		w.Write([]byte("deckid missing"))
		log.Add("Error", err).Error("newgame.json: deckid missing")
		return
	} else {
		log.Add("DeckId", int(parse))
		deck = mydecks[int(parse)]
	}

	if deck == nil {
		w.WriteHeader(500)
		w.Write([]byte("decks error"))
		log.Warn("queue: start failed, deck missing")
		return
	}

	if deck.Count() < 21 {
		w.WriteHeader(500)
		w.Write([]byte("decks error"))
		log.Add("Count", deck.Count()).Warn("queue: start failed, deck too small")
		return
	}

	gameid := <-queue.Start(session, deck).Done
	json.Json{
		"gameid": gameid,
	}.Write(w)
	log.Add("GameId", gameid).Info("newgame.json")
})
