package api

import (
	"net/http"

	"github.com/zachtaylor/7elements/server/internal"
)

// SignupHandler returns a http.HandlerFunc that performs internal signup
func SignupHandler(rt *Runtime) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := rt.Root.Logger.New().Source().Add("Addr", r.RemoteAddr)

		if r.Method != "POST" {
			w.WriteHeader(404)
			log.Add("Method", r.Method).Warn("only POST allowed")
			return
		}

		session, _ := rt.Sessions.Cookie(r)
		if session != nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			log.Add("SessionId", session.ID).Info("request has valid session cookie")
			return
		}

		username := r.FormValue("username")
		email := r.FormValue("email")
		password1 := internal.HashPassword(r.FormValue("password1"), rt.Salt)
		password2 := internal.HashPassword(r.FormValue("password2"), rt.Salt)
		log.Add("Username", username).Add("Email", email)

		if !internal.CheckUsername(username) {
			http.Redirect(w, r, "/signup?username", http.StatusSeeOther)
			log.Warn("invalid username")
		} else if rt.Root.Accounts.Test(username) != nil {
			http.Redirect(w, r, "/signup?usernametaken&email="+email, http.StatusSeeOther)
			log.Error("duplicate is online right")
		} else if account, err := rt.Root.Accounts.Get(username); account != nil {
			http.Redirect(w, r, "/signup?usernametaken&email="+email, http.StatusSeeOther)
			log.Add("Error", err).Error("duplicate exists")
		} else if password1 != password2 {
			http.Redirect(w, r, "/signup?passwordmatch&email="+email+"&username="+username, http.StatusSeeOther)
			log.Warn("password mismatch")
		} else if account, session, err := internal.Signup(rt.Root, rt.Sessions, username, email, password1); err != nil {
			http.Redirect(w, r, "/login?account", http.StatusSeeOther)
			log.Add("Error", err).Warn("500")
		} else {
			session.WriteCookie(w)
			w.Write([]byte(redirectHomeTpl))
			log.Add("SessionID", account.SessionID).Info("accept")
		}
	})
}
