package api

import (
	"elemen7s.com"
	"elemen7s.com/server/security"
	"net/http"
	zhttp "ztaylor.me/http"
	"ztaylor.me/log"
)

var LoginHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log := log.Add("Addr", r.RemoteAddr)

	if r.Method != "POST" {
		w.WriteHeader(404)
		log.Add("Method", r.Method).Warn("login: only POST allowed")
		return
	}

	if session, err := zhttp.ReadRequestCookie(r); session != nil {
		http.Redirect(w, r, "/", 307)
		log.Add("SessionId", session.Id).Info("login: request already has valid session cookie")
		return
	} else if err != nil {
		log.Clone().Add("Error", err).Warn("login: ignoring cookie...")
	}

	username := r.FormValue("username")
	password := security.HashPassword(r.FormValue("password"))

	log.Add("Username", username)

	if account := vii.AccountService.Test(username); account != nil {
		log.Add("SessionId", account.SessionId)

		if session := zhttp.SessionCache(account.SessionId); session == nil {
			log.Error("login: account cache hit no session")
		} else if account.Password != password {
			log.Warn("login: password does not match")
		} else {
			log.Add("SessionId", account.SessionId).Info("login: rewrite sessionid")
			GrantSession(account, w, r, "Login Re-Accepted!")
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

	GrantSession(account, w, r, "Login Success!")
})
