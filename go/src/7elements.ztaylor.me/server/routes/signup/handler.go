package signup

import (
	"7elements.ztaylor.me/accounts"
	"7elements.ztaylor.me/decks"
	"7elements.ztaylor.me/server/routes/login"
	"7elements.ztaylor.me/server/security"
	"net/http"
	"time"
	"ztaylor.me/events"
	"ztaylor.me/http/sessions"
	"ztaylor.me/log"
)

var Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log := log.Add("RemoteAddr", r.RemoteAddr)

	if r.Method != "POST" {
		w.WriteHeader(404)
		log.Add("Method", r.Method).Warn("signup: only POST allowed")
		return
	}

	session, err := sessions.ReadRequestCookie(r)
	if session != nil {
		http.Redirect(w, r, "/", 307)
		log.Add("SessionId", session.Id).Info("signup: request has valid session cookie")
		return
	} else if err != nil {
		log.Clone().Add("Error", err).Warn("signup: ignoring cookie...")
	}

	username := r.FormValue("username")
	email := r.FormValue("email")
	password1 := security.HashPassword(r.FormValue("password1"))
	password2 := security.HashPassword(r.FormValue("password2"))

	log.Add("Username", username).Add("Email", email)

	if accounts.Test(username) != nil {
		sessions.EraseSessionId(w)
		http.Redirect(w, r, "/signup/?usernametaken&email="+email, 307)
		log.Error("signup: duplicate is online right")
		return
	} else if account, _ := accounts.Load(username); account != nil {
		sessions.EraseSessionId(w)
		http.Redirect(w, r, "/signup/?usernametaken&email="+email, 307)
		log.Add("Error", err).Error("signup: duplicate exists")
		return
	}

	if password1 != password2 {
		http.Redirect(w, r, "/signup/?passwordmatch&email="+email+"&username="+username, 307)
		log.Warn("signup: password mismatch")
		return
	}

	acceptLanguage := r.Header.Get("Accept-Language")
	acceptLanguage = acceptLanguage[0:5]
	if acceptLanguage == "" {
		acceptLanguage = "en-US"
	}

	login.DoUnsafe(&accounts.Account{
		Username: username,
		Email:    email,
		Password: password1,
		Skill:    1000,
		Packs:    21,
		Language: acceptLanguage,
		Register: time.Now(),
	}, w, r, "Signup success!")

	decks.Store(username, decks.NewDecks())
	for _, deck := range decks.Test(username) {
		deck.Username = username
	}
	if err := decks.Insert(username, 0); err != nil {
		log.Add("Error", err).Error("signup: grant 3 decks")
		return
	}

	if err := accounts.Insert(username); err != nil {
		accounts.Forget(username)
		decks.Forget(username)
		log.Add("Error", err).Error("signup: account insert")
		sessions.EraseSessionId(w)
		w.WriteHeader(500)
	} else {
		log.Debug("signup: sucess")
		events.Fire("Signup", username)
	}
})
