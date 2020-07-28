package api

import (
	"net/http"

	"github.com/zachtaylor/7elements/runtime"
	"github.com/zachtaylor/7elements/server/internal"
	"ztaylor.me/cast"
	"ztaylor.me/cast/charset"
)

// LoginHandler returns a http.HandlerFunc that performs internal login
func LoginHandler(rt *runtime.T) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := rt.Log().Tag("api/login").Add("Addr", r.RemoteAddr)

		if r.Method != "POST" {
			w.WriteHeader(404)
			log.Add("Method", r.Method).Warn("only POST allowed")
			return
		}

		if session, _ := rt.Sessions.Cookie(r); session != nil {
			session.WriteCookie(w, !rt.IsDevEnv)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			log.Add("SessionID", session.ID()).Info("request cookie exists")
			return
		}

		username := r.FormValue("username")
		log.Add("Username", username)

		if !cast.InCharset(username, charset.AlphaCapitalNumeric) {
			http.Redirect(w, r, "/login?username", http.StatusSeeOther)
			log.Warn("invalid username")
		} else if account, err := rt.Accounts.Get(username); account == nil {
			http.Redirect(w, r, "/login?account", http.StatusSeeOther)
			log.Add("Error", err).Warn("invalid account")
		} else if password := internal.HashPassword(r.FormValue("password"), rt.PassSalt); password != account.Password {
			http.Redirect(w, r, "/login?password", http.StatusSeeOther)
			log.Warn("wrong password")
		} else if player, err := rt.Players.Login(account); player == nil {
			http.Redirect(w, r, "/login", http.StatusInternalServerError)
			log.Add("Error", err).Error("update account")
		} else {
			player.Session.WriteCookie(w, !rt.IsDevEnv)
			w.Write([]byte(redirectHomeTpl))
			log.Add("SessionID", account.SessionID).Info("accept")
		}
	})
}
