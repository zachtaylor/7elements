package logout

import (
	"7elements.ztaylor.me/log"
	"7elements.ztaylor.me/server/sessionman"
	"net/http"
)

var Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log.Add("RemoteAddr", r.RemoteAddr)

	if session, err := sessionman.ReadRequestCookie(r); err == nil {
		sessionman.RevokeSession(session.Username)
		log.Debug("logout: success")
	} else {
		log.Warn("logout: cookie missing")
	}

	http.Redirect(w, r, "/", 307)
})
