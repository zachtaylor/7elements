package api

import (
	"net/http"

	"ztaylor.me/cast"
)

func MyAccountHandler(rt *Runtime) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := rt.Root.Logger.New().Tag("api/myaccount")
		if session := rt.Sessions.Cookie(r); session == nil {
			log.Add("RemoteAddr", r.RemoteAddr).Warn("session required")
		} else {
			w.Write(cast.BytesS(rt.Root.AccountJSON(session.Name()).String()))
		}
	})
}
