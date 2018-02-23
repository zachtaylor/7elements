package api

import (
	"elemen7s.com/accounts"
	"elemen7s.com/decks"
	// "elemen7s.com/emails"
	// "elemen7s.com/options"
	"elemen7s.com/server/security"
	"net/http"
	"time"
	"ztaylor.me/events"
	zhttp "ztaylor.me/http"
	"ztaylor.me/log"
)

var SignupHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log := log.Add("Addr", r.RemoteAddr)

	if r.Method != "POST" {
		w.WriteHeader(404)
		log.Add("Method", r.Method).Warn("signup: only POST allowed")
		return
	}

	session, err := zhttp.ReadRequestCookie(r)
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
		http.Redirect(w, r, "/signup/?usernametaken&email="+email, 307)
		log.Error("signup: duplicate is online right")
		return
	} else if account, _ := accounts.Load(username); account != nil {
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

	account := &accounts.Account{
		Username: username,
		Email:    email,
		Password: password1,
		Skill:    1000,
		Packs:    21,
		Language: acceptLanguage,
		Register: time.Now(),
	}

	decks.Store(username, decks.NewDecks())
	for _, deck := range decks.Test(username) {
		deck.Username = username
	}
	if err := decks.Insert(username, 0); err != nil {
		log.Add("Error", err).Error("/api/signup: grant decks")
		return
	}

	// if err := emails.SendValidationEmail(account); err != nil {
	// 	log.Clone().Add("mail-user", options.String("mail-user")).Add("mail-pass", options.String("mail-pass")).Add("mail-host", options.String("mail-host")).Add("Error", err).Error("/api/signup: send validation email")
	// }

	GrantSession(account, w, r, "Signup success!")

	if err := accounts.Insert(username); err != nil {
		accounts.Forget(username)
		decks.Forget(username)
		log.Add("Error", err).Error("/api/signup: account insert")
		w.WriteHeader(500)
	} else {
		log.Info("/api/signup")
		events.Fire("Signup", username)
	}
})
