package api

import (
	"net/http"

	"ztaylor.me/cast"
)

func MyAccountHandler(rt *Runtime) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := rt.Root.Logger.New().Source()
		if s, err := rt.Sessions.Cookie(r); s == nil {
			log.Add("RemoteAddr", r.RemoteAddr).Add("Error", err).Warn("session required")
		} else {
			w.Write(cast.BytesS(rt.Root.FindAccountJSON(s.Name()).String()))
		}
	})
}
