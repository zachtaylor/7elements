package api

import (
	"fmt"
	"net/http"

	httpsessions "ztaylor.me/http/session"
)

func LogoutHandler(rt *Runtime) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := rt.Root.Logger.New().Tag("api/logout").Add("Addr", r.RemoteAddr)

		if session := rt.Sessions.Cookie(r); session != nil {
			log.Debug("revoking")
			rt.Sessions.Remove(session)
		} else {
			log.Warn("cookie missing")
		}

		httpsessions.EraseSessionID(w)
		w.Write([]byte(fmt.Sprintf(redirectHomeTpl, "Logout Success")))
	})
}
