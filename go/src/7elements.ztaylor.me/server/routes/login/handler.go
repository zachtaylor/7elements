package login

import (
	"7elements.ztaylor.me/accounts"
	"7elements.ztaylor.me/server/security"
	"net/http"
	"ztaylor.me/http/sessions"
	"ztaylor.me/log"
)

var Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log := log.Add("RemoteAddr", r.RemoteAddr)

	if r.Method != "POST" {
		w.WriteHeader(404)
		log.Add("Method", r.Method).Warn("login: only POST allowed")
		return
	}

	if session, err := sessions.ReadRequestCookie(r); session != nil {
		http.Redirect(w, r, "/", 307)
		log.Add("SessionId", session.Id).Info("login: request already has valid session cookie")
		return
	} else if err != nil {
		log.Clone().Add("Error", err).Warn("login: ignoring cookie...")
	}

	username := r.FormValue("username")
	password := security.HashPassword(r.FormValue("password"))

	log.Add("Username", username)

	if account := accounts.Test(username); account != nil {
		log.Add("SessionId", account.SessionId)

		if session := sessions.Cache[account.SessionId]; session == nil {
			log.Error("login: account cache hit no session")
		} else if account.Password != password {
			log.Warn("login: password does not match")
		} else {
			log.Add("SessionId", account.SessionId).Info("login: rewrite sessionid")
			DoUnsafe(account, w, r, "Login Re-Accepted!")
		}

		return
	}

	account, err := accounts.Load(username)
	if account == nil {
		if err != nil {
			log.Add("Error", err)
		}

		sessions.EraseSessionId(w)
		http.Redirect(w, r, "/login/?account", 307)
		log.Warn("login: account not found")
		return
	}

	if a := accounts.Test(username); a != nil {
		http.Redirect(w, r, "/login/?account", 307)
		log.Add("SessionId", a.SessionId).Warn("login: account already online")
		return
	}

	if account.Password != password {
		sessions.EraseSessionId(w)
		http.Redirect(w, r, "/login/?password#"+username, 307)
		log.Warn("login: password does not match")
		return
	}

	DoUnsafe(account, w, r, "Login Success!")
	log.Info("login succeeded")
})
