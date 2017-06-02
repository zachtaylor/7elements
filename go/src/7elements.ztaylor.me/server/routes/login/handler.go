package login

import (
	"7elements.ztaylor.me"
	"7elements.ztaylor.me/log"
	"7elements.ztaylor.me/server/security"
	"7elements.ztaylor.me/server/sessionman"
	"net/http"
)

var Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log.Add("RemoteAddr", r.RemoteAddr)

	if r.Method != "POST" {
		w.WriteHeader(404)
		log.Add("Method", r.Method).Warn("login: only POST allowed")
		return
	}

	if session, err := sessionman.ReadRequestCookie(r); session != nil {
		http.Redirect(w, r, "/", 307)
		log.Add("SessionId", session.Id).Info("login: request has valid session cookie")
		return
	} else if err != nil {
		log.Clone().Add("Error", err).Warn("login: ignoring cookie...")
	}

	username := r.FormValue("username")
	password := security.HashPassword(r.FormValue("password"))

	log.Add("Username", username)

	if account := SE.Accounts.Cache[username]; account != nil {
		if account.SessionId > 0 {
			log.Add("SessionId", account.SessionId)
		}

		sessionman.EraseSessionId(w)
		log.Error("login: account cache hit panic")
		return
	}

	account, err := SE.Accounts.Get(username)
	if account == nil {
		if err != nil {
			log.Add("Error", err)
		}

		sessionman.EraseSessionId(w)
		http.Redirect(w, r, "/login/?account", 307)
		log.Warn("login: account not found")
		return
	}

	if account.Password != password {
		sessionman.EraseSessionId(w)
		http.Redirect(w, r, "/login/?password#"+username, 307)
		log.Warn("login: password does not match")
		return
	}

	DoUnsafe(account, w, r, "Login Success!")
	log.Debug("login: success")
})
