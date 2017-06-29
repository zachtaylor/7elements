package cards

import (
	"7elements.ztaylor.me"
	"7elements.ztaylor.me/log"
	"7elements.ztaylor.me/server/sessionman"
	"7elements.ztaylor.me/server/util"
	"net/http"
	"strconv"
)

var Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	acceptlanguage := serverutil.ReadAcceptLanguage(r)

	if session, _ := sessionman.ReadRequestCookie(r); session != nil {
		if account := SE.Accounts.Cache[session.Username]; account != nil {
			acceptlanguage = account.Language
		}
	}

	if r.RequestURI == "/api/cards.json" {
		WriteAllCards(w, acceptlanguage)
	} else if len(r.RequestURI) < 13 {
		w.WriteHeader(500)
		log.Add("RemoteAddr", r.RemoteAddr).Add("RequestURI", r.RequestURI).Error("cards: card id unavailable")
	} else if cardidI, err := strconv.Atoi(r.RequestURI[7 : len(r.RequestURI)-5]); err == nil {
		WriteCardId(cardidI, w, acceptlanguage)
	} else {
		w.WriteHeader(500)
		log.Add("RemoteAddr", r.RemoteAddr).Add("Error", err).Error("cards: parse card id")
	}
})
