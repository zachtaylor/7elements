package apihttp

import (
	"net/http"

	"github.com/zachtaylor/7elements/server/api"
	"github.com/zachtaylor/7elements/server/internal"
	"taylz.io/http/session"
)

// LoginHandler returns a http.HandlerFunc that performs internal login
func LoginHandler(server internal.Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := server.Log().Add("Addr", r.RemoteAddr)

		// if r.Method == "OPTIONS" {
		// 	w.Header().Add("Allow", "OPTIONS, POST")
		// 	return
		// }

		if r.Method != "POST" {
			w.WriteHeader(404)
			log.Add("Method", r.Method).Warn("only POST allowed")
			return
		}

		if s, err := server.Sessions().ReadHTTP(r); s == nil {
			if err == session.ErrNoID {
				log.Add("Remote", r.RemoteAddr).Out("first time login")
			}
		} else {
			http.Redirect(w, r, r.Header.Get("Origin"), http.StatusSeeOther)
			log.Add("SessionID", s.ID()).Warn("request cookie exists")
			return
		}

		username := r.FormValue("username")
		if err := api.CheckUsername(username); err != nil {
			http.Redirect(w, r, r.Header.Get("Origin")+"/account#login", http.StatusSeeOther)
			log.Add("Error", err.Error()).Warn("invalid username")
		}
		log = log.Add("Username", username)

		password := server.HashPassword(r.FormValue("password"))
		account, session, err := api.Login(server, username, password)
		if account == nil || session == nil {
			log.Add("Error", err).Warn("deny")
			http.Redirect(w, r, r.Header.Get("Origin")+"/account#login", http.StatusSeeOther)
		} else {
			log.Add("SessionID", account.SessionID).Info("ok")
			server.Sessions().WriterHTTP(w, session)
		}
	})
}
