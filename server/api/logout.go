package api

import (
	"fmt"
	"net/http"

	"github.com/zachtaylor/7elements/server/runtime"
	httpsessions "ztaylor.me/http/session"
)

func LogoutHandler(t *runtime.T) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := t.Log().Tag("api/logout").Add("Addr", r.RemoteAddr)

		if session, _ := t.Sessions.Cookie(r); session != nil {
			log.Debug("revoking")
			session.Close()
		} else {
			log.Warn("cookie missing")
		}

		httpsessions.EraseSessionID(w, !t.Settings.Devenv)
		w.Write([]byte(fmt.Sprintf(redirectHomeTpl, "Logout Success")))
	})
}
