package apihttp

import (
	"net/http"

	"github.com/zachtaylor/7elements/account"
	"github.com/zachtaylor/7elements/db/accounts"
	"github.com/zachtaylor/7elements/server/api"
	"github.com/zachtaylor/7elements/server/runtime"
)

// SignupHandler returns a http.HandlerFunc that performs internal signup
func SignupHandler(rt *runtime.T) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := rt.Log().Add("Addr", r.RemoteAddr)

		if r.Method != "POST" {
			w.WriteHeader(404)
			log.Add("Method", r.Method).Warn("only POST allowed")
			return
		}

		session := rt.Sessions.RequestSessionCookie(r)
		if session != nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			log.Add("SessionId", session.ID).Info("request has valid session cookie")
			return
		}

		username := r.FormValue("username")
		email := r.FormValue("email")
		password1 := rt.PassHash(r.FormValue("password1"))
		password2 := rt.PassHash(r.FormValue("password2"))
		log.Add("Username", username).Add("Email", email)

		if err := api.CheckUsername(username); err != nil {
			log.Add("Error", err.Error()).Warn("invalid username")
			http.Redirect(w, r, "/signup?username", http.StatusSeeOther)
		} else if a, err := accounts.Get(rt.DB, username); a != nil {
			http.Redirect(w, r, "/signup?usernametaken&email="+email, http.StatusSeeOther)
			log.Add("Error", err).Error("duplicate exists")
		} else if password1 != password2 {
			http.Redirect(w, r, "/signup?passwordmatch&email="+email+"&username="+username, http.StatusSeeOther)
			log.Warn("password mismatch")
		} else {
			session := rt.Sessions.Grant(username)
			account := account.Make(username, email, password1, session.ID())
			if err := accounts.Insert(rt.DB, account); err != nil {
				log.Add("Error", err).Error("signup failed")
				return
			}
			rt.Accounts.Set(username, account)
			rt.Sessions.WriteSessionCookie(w, session)
			w.Write([]byte(redirectHomeTpl))
			go rt.Ping()
			log.Info()
		}
	})
}
