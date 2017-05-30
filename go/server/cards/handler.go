package cards

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
	if err != nil {
		w.WriteHeader(400)
		w.Write([]byte(err.Error()))
		log.Add("Error", err).Error("cards: no session")
		return
	}

	account := SE.Accounts.Cache[session.Username]
	if account == nil {
		w.WriteHeader(500)
		w.Write([]byte(err.Error()))
		log.Add("Error", err).Error("cards: no account")
		return
	}

	if r.RequestURI == "/api/cards.json" {
		WriteAllCards(w, account.Language)
		return
	} else if len(r.RequestURI) < 13 {
		w.WriteHeader(500)
		log.Add("RequestURI", r.RequestURI).Error("cards: card id unavailable")
		return
	} else if cardidI, err := strconv.Atoi(r.RequestURI[7 : len(r.RequestURI)-5]); err == nil {
		WriteCardId(uint(cardidI), w, account.Language)
		return
	} else {
		w.WriteHeader(500)
		log.Add("Error", err).Error("cards: parse card id")
	}
})
