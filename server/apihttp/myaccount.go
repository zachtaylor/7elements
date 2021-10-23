package apihttp

import (
	"encoding/json"
	"net/http"

	"github.com/zachtaylor/7elements/server/runtime"
)

func MyAccountHandler(rt *runtime.T) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := rt.Logger.Add("Addr", r.RemoteAddr)
		if s, err := rt.Sessions.GetRequestCookie(r); s == nil {
			log.Add("Error", err).Warn("session required")
		} else if account := rt.Accounts.Get(s.Name()); account == nil {
			log.Warn("account missing")
		} else {
			bytes, _ := json.Marshal(account.Data())
			w.Write(bytes)
		}
	})
}
