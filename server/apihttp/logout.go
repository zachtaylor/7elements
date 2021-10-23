package apihttp

import (
	"net/http"

	"github.com/zachtaylor/7elements/server/runtime"
	"taylz.io/http/session"
)

func LogoutHandler(rt *runtime.T) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := rt.Log().Add("Addr", r.RemoteAddr)

		if s, err := rt.Sessions.GetRequestCookie(r); err == nil {
			log.Add("SessionID", s.ID()).Add("Username", s.Name()).Info("ok")
			rt.Sessions.Remove(s.ID())
		} else if err == session.ErrNoCookie {
			log.Warn("cookie missing")
		} else if err == session.ErrExpired {
			log.Warn("cookie expired")
		}

		// httpsessions.EraseSessionID(w, !t.IsDevEnv)
		w.Write([]byte(redirectHomeTpl))
	})
}
