package api

import (
	"net/http"

	zhttp "ztaylor.me/http"
	"ztaylor.me/log"
)

var LogoutHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log := log.Add("Addr", r.RemoteAddr)

	if session, err := zhttp.ReadRequestCookie(r); session != nil {
		log.Debug("/api/logout")
		zhttp.SessionService.Revoke(session.ID)
	} else {
		log.Add("Error", err).Warn("/api/logout: cookie missing")
	}

	http.Redirect(w, r, "/", 307)
})