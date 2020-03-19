package apiws

import (
	"time"

	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/account"
	"github.com/zachtaylor/7elements/server/api"
	"ztaylor.me/events"
	"ztaylor.me/http/websocket"
	"ztaylor.me/log"
)

// TODO write error messages to socket

func Signup(rt *Runtime) websocket.Handler {
	return websocket.HandlerFunc(func(socket *websocket.T, m *websocket.Message) {
		log := rt.Runtime.Root.Logger.New().Tag("apiws/signup").Add("Socket", socket.String())
		if socket.Session != nil {
			log.Warn("session exists")
		} else if username := m.Data.GetS("username"); username == "" {
			log.Warn("username missing")
		} else if !api.CheckUsername(username) {
			log.Warn("username banned")
		} else if email := m.Data.GetS("email"); false {
		} else if account, err := rt.Runtime.Root.Accounts.Get(username); account != nil {
			log.Add("Error", err).Warn("account exists")
		} else if password1, password2 := api.HashPassword(m.Data.GetS("password1"), rt.Runtime.Salt), api.HashPassword(m.Data.GetS("password2"), rt.Runtime.Salt); password1 != password2 {
			log.Add("Error", err).Error("passwords don't match")
		} else if account = signup(rt.Runtime, log, username, email, password1); account == nil {
			log.Add("Error", err).Error("signup failed")
		} else {
			go ping(rt)
			redirect(socket, "/")

			if s, err := api.Login(rt.Runtime, account); s == nil {
				log.Add("Error", err).Error("login")
			} else {
				socket.Session = s
				socket.Message("/myaccount", rt.Runtime.Root.AccountJSON(s.Name()))
				log.Info()
			}
		}
	})
}
func signup(rt *api.Runtime, log *log.Entry, username, email, password string) *account.T {
	account := &account.T{
		Username: username,
		Email:    email,
		Password: password,
		Skill:    1000,
		Register: time.Now(),
	}
	rt.Root.Accounts.Cache(account)

	if err := rt.Root.Accounts.Insert(account); err != nil {
		// http.Redirect(w, r, "/signup/?error="+email+"&username="+username, http.StatusInternalServerError)
		log.Add("Error", err).Error("account insert")
		return nil
	}

	events.Fire("Signup", username)

	for i := 1; i < 4; i++ {
		deck := vii.NewAccountDeck()
		deck.ID = i
		deck.Username = username
		if err := rt.Root.AccountsDecks.Update(deck); err != nil {
			log.Add("Error", err).Error("grant decks")
			return nil
		}
	}

	// if err := emails.SendValidationEmail(account); err != nil {
	// 	log.Clone().Add("mail-user", options.String("mail-user")).Add("mail-pass", options.String("mail-pass")).Add("mail-host", options.String("mail-host")).Add("Error", err).Error("/api/signup: send validation email")
	// }

	if s, err := api.Login(rt, account); s == nil {
		log.Add("Error", err).Error("login")
		return nil
	} else {
		// w.Write([]byte(redirectHomeTpl))
	}
	return account
}
