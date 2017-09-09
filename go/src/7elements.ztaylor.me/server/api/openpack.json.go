package api

import (
	"7elements.ztaylor.me/accounts"
	"7elements.ztaylor.me/accountscards"
	"net/http"
	"ztaylor.me/http/sessions"
	"ztaylor.me/json"
	"ztaylor.me/log"
)

var OpenPackJsonHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log := log.Add("RemoteAddr", r.RemoteAddr)

	if r.Method != "GET" {
		w.WriteHeader(404)
		log.Add("Method", r.Method).Warn("openpack.json: only GET allowed")
		return
	}

	session, err := sessions.ReadRequestCookie(r)
	if session == nil {
		if err != nil {
			sessions.EraseSessionId(w)
			log.Add("Error", err)
		}
		w.WriteHeader(400)
		w.Write([]byte("session missing"))
		log.Error("openpack.json: session missing")
		return
	}

	log.Add("Username", session.Username)
	account := accounts.Test(session.Username)
	if account == nil {
		sessions.EraseSessionId(w)
		w.WriteHeader(500)
		log.Error("openpack.json: account missing")
		return
	}

	if account.Packs < 1 {
		w.WriteHeader(500)
		w.Write([]byte("no packs"))
		log.Warn("openpack.json: no packs to open")
		return
	}

	if err := accounts.UpdatePackCount(account.Username, account.Packs-1); err != nil {
		w.WriteHeader(500)
		w.Write([]byte("error opening pack"))
		log.Add("Error", err).Error("openpack.json: error opening pack")
		return
	}

	accountcards, err := accountscards.Get(account.Username)
	if err != nil {
		sessions.EraseSessionId(w)
		w.WriteHeader(500)
		log.Add("Error", err).Error("openpack.json: collection")
		return
	}

	carddata := make([]int, 7)
	for i, card := range accountscards.NewPack(account.Username) {
		carddata[i] = card.CardId

		if err := accountscards.InsertCard(card); err != nil {
			sessions.EraseSessionId(w)
			w.WriteHeader(500)
			log.Add("Error", err).Error("openpack.json: insert card copy")
			return
		}

		if list := accountcards[card.CardId]; list != nil {
			accountcards[card.CardId] = append(list, card)
		} else {
			accountcards[card.CardId] = []*accountscards.AccountCard{card}
		}
	}

	json.Json{
		"cards": carddata,
	}.Write(w)
	log.Add("PacksRemaining", account.Packs).Add("Cards", carddata).Info("openpack")
})
