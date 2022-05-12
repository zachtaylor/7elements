package apihttp

import (
	"net/http"

	"github.com/zachtaylor/7elements/db/accounts"
	"github.com/zachtaylor/7elements/server/api"
	"github.com/zachtaylor/7elements/server/internal"
)

// SignupHandler returns a http.HandlerFunc that performs internal signup
func SignupHandler(server internal.Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := server.Log().Add("Addr", r.RemoteAddr)

		if r.Method != "POST" {
			w.WriteHeader(404)
			log.Add("Method", r.Method).Warn("only POST allowed")
			return
		}

		session, _ := server.GetSessionManager().GetRequestCookie(r)
		if session != nil {
			http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
			log.Add("SessionID", session.ID).Warn("session exists")
			return
		}

		username := r.FormValue("username")
		email := r.FormValue("email")
		password1 := server.HashPassword(r.FormValue("password1"))
		password2 := server.HashPassword(r.FormValue("password2"))
		log.Add("Username", username).Add("Email", email)

		if err := api.CheckUsername(username); err != nil {
			log.Add("Error", err.Error()).Warn("invalid username")
			http.Redirect(w, r, "/signup?username", http.StatusSeeOther)
		} else if a, err := accounts.Get(server.GetDB(), username); a != nil {
			http.Redirect(w, r, "/signup?usernametaken&email="+email, http.StatusSeeOther)
			log.Add("Error", err).Error("duplicate exists")
		} else if password1 != password2 {
			http.Redirect(w, r, "/signup?passwordmatch&email="+email+"&username="+username, http.StatusSeeOther)
			log.Warn("password mismatch")
		} else if _, _, err := api.Signup(server, username, email, password1); err != nil {
			http.Redirect(w, r, "/signup", http.StatusSeeOther)
			log.Add("Error", err).Error("internal error")
		} else {
			log.Info("ok")
		}
	})
}
