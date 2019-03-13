package api

import (
	"net/http"

	"ztaylor.me/http/sessions"
	"ztaylor.me/log"
)

const tagLogout = "/api/logout"

func LogoutHandler(sessions *sessions.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := log.Add("Addr", r.RemoteAddr)

		if session := sessions.ReadRequestCookie(r); session != nil {
			log.Debug(tagLogout)
			session.Revoke()
		} else {
			log.Warn(tagLogout, ": cookie missing")
		}

		http.Redirect(w, r, "/", 307)
	})
}
