package apihttp

import (
	"net/http"

	"github.com/zachtaylor/7elements/db/accounts"
	"github.com/zachtaylor/7elements/server/api"
	"github.com/zachtaylor/7elements/server/runtime"
)

// LoginHandler returns a http.HandlerFunc that performs internal login
func LoginHandler(rt *runtime.T) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := rt.Log().Add("Addr", r.RemoteAddr)

		if r.Method != "POST" {
			w.WriteHeader(404)
			log.Add("Method", r.Method).Warn("only POST allowed")
			return
		}

		if session := rt.Sessions.RequestSessionCookie(r); session != nil {
			http.Redirect(w, r, "/", http.StatusSeeOther)
			log.Add("SessionID", session.ID()).Info("request cookie exists")
			return
		}

		username := r.FormValue("username")
		if err := api.CheckUsername(username); err != nil {
			http.Redirect(w, r, "/login?username", http.StatusSeeOther)
			log.Add("Error", err.Error()).Warn("invalid username")
		}
		log.Add("Username", username)

		if account, err := accounts.Get(rt.DB, username); account == nil {
			http.Redirect(w, r, "/login?account", http.StatusSeeOther)
			log.Add("Error", err).Warn("invalid account")
		} else if password := rt.PassHash(r.FormValue("password")); password != account.Password {
			http.Redirect(w, r, "/login?password", http.StatusSeeOther)
			log.Warn("wrong password")
		} else {
			session := rt.Sessions.Grant(username)
			account.SessionID = session.ID()
			rt.Accounts.Set(username, account)
			rt.Sessions.WriteSessionCookie(w, session)
			w.Write([]byte(redirectHomeTpl))
			log.Add("SessionID", account.SessionID).Info("accept")
		}
	})
}
