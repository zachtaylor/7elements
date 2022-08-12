package apihttp

import (
	"database/sql"
	"net/http"

	"github.com/zachtaylor/7elements/db/accounts"
	"github.com/zachtaylor/7elements/server/api"
	"github.com/zachtaylor/7elements/server/internal"
)

// Username returns a http.HandlerFunc that reports whether the query param v is a legal available username
func UsernameHandler(server internal.Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log := server.Log().Add("Remote", r.RemoteAddr)
		if v := r.URL.Query().Get("v"); v == "" {
			log.Warn("missing v")
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(`{"error":"missing v"}`))
		} else if err := api.CheckUsername(v); err != nil {
			log.Add("Error", err).Warn("illegal")
			w.WriteHeader(http.StatusForbidden)
			w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		} else if _, err = accounts.Get(server.DB(), v); err == sql.ErrNoRows {
			log.Info("available")
			w.Write([]byte(`{"username":"` + v + `","unique":true}`))
		} else if err != nil {
			log.Add("Error", err).Error("accounts.Get")
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(`{"error":"` + err.Error() + `"}`))
		} else {
			log.Add("Username", v).Warn("taken")
			w.WriteHeader(http.StatusConflict)
			w.Write([]byte(`{"error":"Username already in use"}`))
		}
	})
}
