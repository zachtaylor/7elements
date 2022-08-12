package apihttp

import (
	"net/http"

	"github.com/zachtaylor/7elements/server/internal"
	"taylz.io/http/session"
)

func LogoutHandler(server internal.Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := server.Log().Add("Addr", r.RemoteAddr)

		if s, err := server.Sessions().ReadHTTP(r); err == nil {
			log.Add("SessionID", s.ID()).Add("Username", s.Name()).Info("ok")
			server.Sessions().Remove(s.ID())
		} else if err == session.ErrNoID {
			log.Warn("cookie missing")
		} else if err == session.ErrExpired {
			log.Warn("cookie expired")
		} else {
			log.Info("ok")
			server.Sessions().Remove(s.ID())
		}

		// httpsessions.EraseSessionID(w, !t.IsDevEnv)
		w.Write([]byte(redirectHomeTpl))
	})
}
