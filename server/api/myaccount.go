package api

import (
	"net/http"

	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/http/sessions"
	"ztaylor.me/js"
	"ztaylor.me/log"
)

func MyAccountHandler(sessions *sessions.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if session := sessions.ReadRequestCookie(r); session == nil {
			log.WithFields(log.Fields{
				"RemoteAddr": r.RemoteAddr,
			}).Warn("api/myaccount: session required")
		} else if account, err := vii.AccountService.Get(session.Name()); err != nil {
			log.Add("Error", err).Add("Name", session.Name()).Error("api/myaccount: account missing")
		} else if accountcards, err := vii.AccountCardService.Get(session.Name()); err != nil {
			log.Add("Error", err).Add("Name", session.Name()).Error("api/myaccount: accountcards missing")
		} else if accountdecks, err := vii.AccountDeckService.Get(session.Name()); err != nil {
			log.Add("Error", err).Add("Name", session.Name()).Error("api/myaccount: accountdecks missing")
		} else {
			games := make(map[string]js.Object)
			for _, gameid := range game.Service.GetPlayerGames(session.Name()) {
				if game := game.Service.Get(gameid); game != nil {
					games[gameid] = game.Json(session.Name())
				}
			}
			w.Write([]byte(js.Object{
				"username": session.Name(),
				"email":    account.Email,
				"coins":    account.Coins,
				"cards":    accountcards.Json(),
				"decks":    accountdecks.Json(),
				"games":    games,
			}.String()))
		}
	})
}
