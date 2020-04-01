package api

import (
	"net/http"

	"github.com/zachtaylor/7elements/server/internal"
	"ztaylor.me/cast"
	"ztaylor.me/cast/charset"
)

// LoginHandler returns a http.HandlerFunc that performs internal login
func LoginHandler(rt *Runtime) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := rt.Root.Logger.New().Tag("api/login").Add("Addr", r.RemoteAddr)

		if r.Method != "POST" {
			w.WriteHeader(404)
			log.Add("Method", r.Method).Warn("only POST allowed")
			return
		}

		if session, _ := rt.Sessions.Cookie(r); session != nil {
			session.WriteCookie(w)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			log.Add("SessionID", session.ID()).Info("request cookie exists")
			return
		}

		username := r.FormValue("username")
		log.Add("Username", username)

		if !cast.InCharset(username, charset.AlphaCapitalNumeric) {
			http.Redirect(w, r, "/login?username", http.StatusSeeOther)
			log.Warn("invalid username")
		} else if account, err := rt.Root.Accounts.Get(username); account == nil {
			http.Redirect(w, r, "/login?account", http.StatusSeeOther)
			log.Add("Error", err).Warn("invalid account")
		} else if password := internal.HashPassword(r.FormValue("password"), rt.Salt); password != account.Password {
			http.Redirect(w, r, "/login?password", http.StatusSeeOther)
			log.Warn("wrong password")
		} else if s, err := internal.Login(rt.Root, rt.Sessions, account); s == nil {
			http.Redirect(w, r, "/login", http.StatusInternalServerError)
			log.Add("Error", err).Error("update account")
		} else {
			s.WriteCookie(w)
			w.Write([]byte(redirectHomeTpl))
			log.Add("SessionID", account.SessionID).Info("accept")
		}
	})
}
