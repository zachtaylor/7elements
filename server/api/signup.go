package api

import (
	"net/http"
	"time"

	"github.com/zachtaylor/7elements/account"
	"ztaylor.me/events"
	// "github.com/zachtaylor/7elements/emails"
	// "github.com/zachtaylor/7elements/options"
)

func SignupHandler(rt *Runtime) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := rt.Root.Logger.New().Tag("api/signup").Add("Addr", r.RemoteAddr)

		if r.Method != "POST" {
			w.WriteHeader(404)
			log.Add("Method", r.Method).Warn("only POST allowed")
			return
		}

		session := rt.Sessions.Cookie(r)
		if session != nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			log.Add("SessionId", session.ID).Info("request has valid session cookie")
			return
		}

		username := r.FormValue("username")
		log.Add("Username", username)

		if !CheckUsername(username) {
			http.Redirect(w, r, "/signup?username", http.StatusSeeOther)
			log.Warn("invalid username")
			return
		}

		email := r.FormValue("email")
		password1 := HashPassword(r.FormValue("password1"), rt.Salt)
		password2 := HashPassword(r.FormValue("password2"), rt.Salt)
		log.Add("Email", email)

		if rt.Root.Accounts.Test(username) != nil {
			http.Redirect(w, r, "/signup?usernametaken&email="+email, http.StatusSeeOther)
			log.Error("duplicate is online right")
			return
		} else if account, err := rt.Root.Accounts.Get(username); account != nil {
			http.Redirect(w, r, "/signup?usernametaken&email="+email, http.StatusSeeOther)
			log.Add("Error", err).Error("duplicate exists")
			return
		}

		if password1 != password2 {
			http.Redirect(w, r, "/signup?passwordmatch&email="+email+"&username="+username, http.StatusSeeOther)
			log.Warn("password mismatch")
			return
		}

		a := &account.T{
			Username: username,
			Email:    email,
			Password: password1,
			Skill:    1000,
			Coins:    7,
			Register: time.Now(),
		}

		if err := rt.Root.Accounts.Insert(a); err != nil {
			http.Redirect(w, r, "/signup/?error="+email+"&username="+username, http.StatusInternalServerError)
			log.Add("Error", err).Error("account insert")
			return
		}

		events.Fire("Signup", username)

		for i := 1; i < 4; i++ {
			deck := account.NewDeck()
			deck.ID = i
			deck.Username = username
			if err := rt.Root.AccountsDecks.Update(deck); err != nil {
				log.Add("Error", err).Error("grant decks")
				return
			}
		}

		// if err := emails.SendValidationEmail(account); err != nil {
		// 	log.Clone().Add("mail-user", options.String("mail-user")).Add("mail-pass", options.String("mail-pass")).Add("mail-host", options.String("mail-host")).Add("Error", err).Error("/api/signup: send validation email")
		// }

		if s, err := Login(rt, a); s == nil {
			log.Add("Error", err).Error("login")
		} else {
			w.Write([]byte(redirectHomeTpl))
		}
	})
}
