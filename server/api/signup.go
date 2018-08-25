package api

import (
	"net/http"
	"time"

	"github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/server/security"
	"ztaylor.me/events"
	zhttp "ztaylor.me/http"
	"ztaylor.me/log"
	// "github.com/zachtaylor/7elements/emails"
	// "github.com/zachtaylor/7elements/options"
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
		log.Add("SessionId", session.ID).Info("signup: request has valid session cookie")
		return
	} else if err != nil {
		log.Clone().Add("Error", err).Warn("signup: ignoring cookie...")
	}

	username := r.FormValue("username")
	email := r.FormValue("email")
	password1 := security.HashPassword(r.FormValue("password1"))
	password2 := security.HashPassword(r.FormValue("password2"))

	log.Add("Username", username).Add("Email", email)

	if vii.AccountService.Test(username) != nil {
		http.Redirect(w, r, "/signup/?usernametaken&email="+email, 307)
		log.Error("signup: duplicate is online right")
		return
	} else if account, err := vii.AccountService.Load(username); account != nil {
		http.Redirect(w, r, "/signup/?usernametaken&email="+email, 307)
		log.Add("Error", err).Error("signup: duplicate exists")
		return
	}

	if password1 != password2 {
		http.Redirect(w, r, "/signup/?passwordmatch&email="+email+"&username="+username, 307)
		log.Warn("signup: password mismatch")
		return
	}

	account := &vii.Account{
		Username: username,
		Email:    email,
		Password: password1,
		Skill:    1000,
		Coins:    7,
		Register: time.Now(),
	}

	if err = vii.AccountService.Insert(account); err != nil {
		http.Redirect(w, r, "/signup/?error="+email+"&username="+username, 500)
		log.Add("Error", err).Error("/api/signup: account insert")
		w.WriteHeader(500)
		return
	} else {
		events.Fire("Signup", username)
	}

	for i := 1; i < 4; i++ {
		deck := vii.NewAccountDeck()
		deck.ID = i
		deck.Username = username
		if err := vii.AccountDeckService.Update(deck); err != nil {
			log.Add("Error", err).Error("/api/signup: grant decks")
			return
		}
	}

	// if err := emails.SendValidationEmail(account); err != nil {
	// 	log.Clone().Add("mail-user", options.String("mail-user")).Add("mail-pass", options.String("mail-pass")).Add("mail-host", options.String("mail-host")).Add("Error", err).Error("/api/signup: send validation email")
	// }

	GrantSession(account, w, r, "Signup success!")
	log.Info("/api/signup")
})
