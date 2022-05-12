package apihttp

import (
	"net/http"

	"github.com/zachtaylor/7elements/db/accounts"
	"github.com/zachtaylor/7elements/server/api"
	"github.com/zachtaylor/7elements/server/internal"
	"taylz.io/http/session"
)

// LoginHandler returns a http.HandlerFunc that performs internal login
func LoginHandler(server internal.Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := server.Log().Add("Addr", r.RemoteAddr)

		if r.Method != "POST" {
			w.WriteHeader(404)
			log.Add("Method", r.Method).Warn("only POST allowed")
			return
		}

		if s, err := server.GetSessionManager().GetRequestCookie(r); s == nil {
			if err == session.ErrNoID {
				log.Add("Remote", r.RemoteAddr).Out("first time login")
			}
		} else {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			log.Add("SessionID", s.ID()).Warn("request cookie exists")
			return
		}

		username := r.FormValue("username")
		if err := api.CheckUsername(username); err != nil {
			http.Redirect(w, r, "/login?username", http.StatusSeeOther)
			log.Add("Error", err.Error()).Warn("invalid username")
		}
		log = log.Add("Username", username)

		if account, err := accounts.Get(server.GetDB(), username); account == nil {
			http.Redirect(w, r, "/login?account", http.StatusSeeOther)
			log.Add("Error", err).Warn("invalid account")
		} else if password := server.HashPassword(r.FormValue("password")); password != account.Password {
			http.Redirect(w, r, "/login?password", http.StatusSeeOther)
			log.Warn("wrong password")
		} else {
			session := server.GetSessionManager().Must(username)
			account.SessionID = session.ID()
			server.GetAccounts().Set(username, account)
			server.GetSessionManager().WriteSetCookie(w, session)
			w.Write([]byte(redirectHomeTpl))
			log.Add("SessionID", account.SessionID).Info("ok")
		}
	})
}
