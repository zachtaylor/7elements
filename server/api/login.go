package api

import (
	"net/http"
	"time"

	vii "github.com/zachtaylor/7elements"
	"ztaylor.me/cast"
	"ztaylor.me/http/session"
)

func LoginHandler(rt *Runtime) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := rt.Root.Logger.New().Tag("api/login").Add("Addr", r.RemoteAddr)

		if r.Method != "POST" {
			w.WriteHeader(404)
			log.Add("Method", r.Method).Warn("only POST allowed")
			return
		}

		if session := rt.Sessions.Cookie(r); session != nil {
			session.WriteCookie(w)
			http.Redirect(w, r, "/", http.StatusSeeOther)
			log.Add("SessionID", session.ID()).Info("request cookie exists")
			return
		}

		username := r.FormValue("username")
		log.Add("Username", username)

		if !CheckUsername(username) {
			http.Redirect(w, r, "/login?username", http.StatusSeeOther)
			log.Warn("invalid username")
		} else if account, err := rt.Root.Accounts.Get(username); account == nil {
			http.Redirect(w, r, "/login?account", http.StatusSeeOther)
			log.Add("Error", err).Warn("invalid account")
		} else if password := HashPassword(r.FormValue("password"), rt.Salt); password != account.Password {
			http.Redirect(w, r, "/login?password", http.StatusSeeOther)
			log.Warn("wrong password")
		} else if s, err := Login(rt, account); s == nil {
			http.Redirect(w, r, "/login", http.StatusInternalServerError)
			log.Add("Error", err).Error("update account")
		} else {
			s.WriteCookie(w)
			w.Write([]byte(redirectHomeTpl))
			log.Add("SessionID", account.SessionID).Info("accept")
		}
	})
}

func Login(rt *Runtime, a *vii.Account) (*session.T, error) {
	rt.Root.Accounts.Cache(a)
	log := rt.Root.Logger.New().Add("Username", a.Username).Tag("/api/do_login")
	a.LastLogin = time.Now()
	if err := rt.Root.Accounts.UpdateLogin(a); err != nil {
		return nil, err
	}
	if a.SessionID != "" {
		if s := rt.Sessions.Get(a.SessionID); s != nil {
			log.Debug("found")
			return s, nil
		}
	}
	s := rt.Sessions.Grant(a.Username)
	a.SessionID = s.ID()
	go _loginWaiter(rt, s)
	return s, nil
}

func _loginWaiter(rt *Runtime, s *session.T) {
	for {
		if _, ok := <-s.Done(); !ok {
			break
		}
	}
	rt.Root.Logger.New().With(cast.JSON{
		"SessionID": s.ID(),
		"Username":  s.Name(),
	}).Source().Info("done")
	rt.Root.Accounts.Forget(s.Name())
	rt.Root.AccountsCards.Forget(s.Name())
	rt.Root.AccountsDecks.Forget(s.Name())
}
