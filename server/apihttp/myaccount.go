package apihttp

import (
	"encoding/json"
	"net/http"

	"github.com/zachtaylor/7elements/server/internal"
)

func MyAccountHandler(server internal.Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := server.Log().Add("Addr", r.RemoteAddr)
		if s, err := server.Sessions().ReadHTTP(r); s == nil {
			log.Add("Error", err).Warn("session required")
		} else if account := server.Accounts().Get(s.Name()); account == nil {
			log.Warn("account missing")
		} else {
			bytes, _ := json.Marshal(account.Data())
			w.Write(bytes)
		}
	})
}
