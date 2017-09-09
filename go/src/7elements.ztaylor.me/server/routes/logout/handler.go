package logout

import (
	"net/http"
	"ztaylor.me/http/sessions"
	"ztaylor.me/log"
)

var Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log := log.Add("RemoteAddr", r.RemoteAddr)

	if session, err := sessions.ReadRequestCookie(r); session != nil {
		log.Debug("/api/logout")
		sessions.Revoke(session.Username)
	} else {
		log.Add("Error", err).Warn("/api/logout: cookie missing")
	}

	http.Redirect(w, r, "/", 307)
})
