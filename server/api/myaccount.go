package api

import (
	"net/http"

	"github.com/zachtaylor/7elements/server/runtime"
	"ztaylor.me/cast"
)

func MyAccountHandler(rt *runtime.T) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := rt.Log()
		if s, err := rt.Sessions.Cookie(r); s == nil {
			log.Add("RemoteAddr", r.RemoteAddr).Add("Error", err).Warn("session required")
		} else if p := rt.Players.Get(s.Name()); p == nil {
			log.Add("RemoteAddr", r.RemoteAddr).Add("Error", err).Warn("player missing")
		} else {
			w.Write(cast.BytesS(p.Account.JSON().String()))
		}
	})
}
