package apihttp

import (
	"database/sql"
	"net/http"

	"github.com/zachtaylor/7elements/db/accounts"
	"github.com/zachtaylor/7elements/server/api"
	"github.com/zachtaylor/7elements/server/internal"
)

// Username returns a http.HandlerFunc that reports whether the first query param is a legal available username
func UsernameHandler(server internal.Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := server.Log().With(map[string]any{
			"Remote": r.RemoteAddr,
		})
		for k := range r.URL.Query() {
			log = log.Add("Username", k)

			if err := api.CheckUsername(k); err != nil {
				log.Add("Error", err).Warn("illegal")
				w.Write([]byte(`{"error":"` + err.Error() + `"}`))
			} else if _, err = accounts.Get(server.GetDB(), k); err == sql.ErrNoRows {
				log.Info("available")
				w.Write([]byte(`{"unique":true}`))
			} else if err != nil {
				log.Add("Error", err).Warn("illegal")
				w.WriteHeader(http.StatusInternalServerError)
				w.Write([]byte(`{"error":"` + err.Error() + `"}`))
			} else {
				log.Add("Username", k).Add("Error", err).Warn("taken")
				w.Write([]byte(`{"unique":false}`))
			}

			return
		}
	})
}
