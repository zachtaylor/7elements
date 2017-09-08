package logout

import (
	"7elements.ztaylor.me/server/sessionman"
	"net/http"
	"ztaylor.me/log"
)

var Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log := log.Add("RemoteAddr", r.RemoteAddr)

	if session, err := sessionman.ReadRequestCookie(r); session != nil {
		sessionman.RevokeSession(session.Username)
		log.Debug("logout: success")
	} else {
		log.Add("Error", err).Warn("logout: cookie missing")
	}

	http.Redirect(w, r, "/", 307)
})
