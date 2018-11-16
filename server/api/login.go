package api

import (
	"net/http"

	"github.com/zachtaylor/7elements"
	"ztaylor.me/http/sessions"
	"ztaylor.me/log"
)

func LoginHandler(dbsalt string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := log.Add("Addr", r.RemoteAddr)

		if r.Method != "POST" {
			w.WriteHeader(404)
			log.Add("Method", r.Method).Warn("login: only POST allowed")
			return
		}

		if session := sessions.ReadCookie(r); session != nil {
			http.Redirect(w, r, "/", 307)
			log.Add("SessionID", session.ID()).Info("login: request already has valid session cookie")
			return
		}

		username := r.FormValue("username")
		password := HashPassword(r.FormValue("password"), dbsalt)

		log.Add("Username", username)

		if account := vii.AccountService.Test(username); account != nil {
			log.Add("SessionID", account.SessionID)

			if account.Password != password {
				log.Warn("login: password does not match")
			} else {
				log.Add("SessionID", account.SessionID).Info("login: account is hot")
				GrantSession(w, r, account, "Login Re-Accepted!")
			}

			return
		}

		account, err := vii.AccountService.Load(username)
		if account == nil {
			if err != nil {
				log.Add("Error", err)
			}
			http.Redirect(w, r, "/login/?account", 307)
			log.Warn("login: account not found")
			return
		}

		if account.Password != password {
			http.Redirect(w, r, "/login/?password#"+username, 307)
			log.Warn("login: password does not match")
			return
		}

		GrantSession(w, r, account, "Login Success!")
	})
}
