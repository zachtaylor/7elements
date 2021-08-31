package apihttp

import (
	"encoding/json"
	"net/http"

	"github.com/zachtaylor/7elements/server/runtime"
)

func MyAccountHandler(rt *runtime.T) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := rt.Log()
		if s := rt.Sessions.RequestSessionCookie(r); s == nil {
			log.Add("RemoteAddr", r.RemoteAddr).Warn("session required")
		} else if account := rt.Accounts.Get(s.Name()); account == nil {
			log.Add("RemoteAddr", r.RemoteAddr).Warn("account missing")
		} else {
			bytes, _ := json.Marshal(account.Data())
			w.Write(bytes)
		}
	})
}
