package apihttp

import (
	"net/http"

	"github.com/zachtaylor/7elements/server/runtime"
)

func LogoutHandler(rt *runtime.T) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := rt.Log().Add("Addr", r.RemoteAddr)

		if session := rt.Sessions.RequestSessionCookie(r); session != nil {
			log.Debug("revoking")
			rt.Sessions.Remove(session.ID())
		} else {
			log.Warn("cookie missing")
		}

		// httpsessions.EraseSessionID(w, !t.IsDevEnv)
		w.Write([]byte(redirectHomeTpl))
	})
}
