package signup

import (
	"7elements.ztaylor.me"
	"7elements.ztaylor.me/log"
	"7elements.ztaylor.me/server/login"
	"7elements.ztaylor.me/server/security"
	"7elements.ztaylor.me/server/sessionman"
	"net/http"
	"time"
)

var Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log.Add("RemoteAddr", r.RemoteAddr)

	if r.Method != "POST" {
		w.WriteHeader(404)
		log.Add("Method", r.Method).Warn("signup: only POST allowed")
		return
	}

	if session, err := sessionman.ReadRequestCookie(r); err == nil {
		http.Redirect(w, r, "/", 307)
		log.Add("SessionId", session.Id).Info("signup: request has valid session cookie")
		return
	} else {
		log.Clone().Add("Error", err).Warn("signup: ignoring cookie...")
	}

	username := r.FormValue("username")
	email := r.FormValue("email")
	password1 := security.HashPassword(r.FormValue("password1"))
	password2 := security.HashPassword(r.FormValue("password2"))

	log.Add("Username", username).Add("Email", email)

	if account := SE.Accounts.Cache[username]; account != nil {
		sessionman.EraseSessionId(w)
		http.Redirect(w, r, "/signup/?usernametaken&email="+email, 307)
		log.Error("signup: duplicate is online right")
		return
	} else if account, _ := SE.Accounts.Get(username); account != nil {
		sessionman.EraseSessionId(w)
		http.Redirect(w, r, "/signup/?usernametaken&email="+email, 307)
		log.Error("signup: duplicate exists")
		return
	}

	if password1 != password2 {
		http.Redirect(w, r, "/signup/?passwordmatch&email="+email+"&username="+username, 307)
		log.Warn("signup: password mismatch")
	}

	acceptLanguage := r.Header.Get("Accept-Language")
	acceptLanguage = acceptLanguage[0:5]
	if acceptLanguage == "" {
		acceptLanguage = "en-US"
	}

	login.DoUnsafe(&SE.Account{
		Username: username,
		Email:    email,
		Password: password1,
		Language: acceptLanguage,
		Register: time.Now(),
	}, w, r, "Signup success!")

	if err := SE.Accounts.Insert(username); err != nil {
		delete(SE.Accounts.Cache, username)
		log.Add("Error", err).Error("signup: account insert")
		sessionman.EraseSessionId(w)
		w.WriteHeader(500)
	} else {
		log.Debug("signup: sucess")
	}
})
