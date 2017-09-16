package cards

import (
	"7elements.ztaylor.me/accounts"
	"7elements.ztaylor.me/cards"
	"7elements.ztaylor.me/server/util"
	"net/http"
	"strconv"
	"ztaylor.me/http/sessions"
	"ztaylor.me/log"
)

var Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log := log.Add("RemoteAddr", r.RemoteAddr)
	acceptlanguage := serverutil.ReadAcceptLanguage(r)

	if acceptlanguage != "en-US" {
		log.Clone().Add("AcceptLanguage", acceptlanguage).Add("UserAgent", r.UserAgent()).Add("RemoteAddr", r.RemoteAddr).Warn("/api/cards: corrected en-us language")
		acceptlanguage = "en-US"
	}

	if session, _ := sessions.ReadRequestCookie(r); session != nil {
		if account := accounts.Test(session.Username); account != nil {
			acceptlanguage = account.Language
		}
	}

	texts := cards.TextsCache[acceptlanguage]

	if r.RequestURI == "/api/cards.json" {
		WriteAllCards(w, texts)
		log.Debug("cards.json")
	} else if len(r.RequestURI) < 17 {
		w.WriteHeader(500)
		log.Error("cards.json: card id unavailable")
	} else if cardidI, err := strconv.Atoi(r.RequestURI[11 : len(r.RequestURI)-5]); err == nil {
		WriteCardId(cardidI, w, texts)
		log.Add("CardId", cardidI).Debug("cards.json")
	} else {
		w.WriteHeader(500)
		log.Add("Error", err).Error("cards.json: parse card id")
	}
})
