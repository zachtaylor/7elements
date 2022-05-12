package apihttp

import (
	"net/http"

	"github.com/zachtaylor/7elements/server/internal"
)

func GlobalDataHandler(server internal.Server) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(server.GetGameVersion().GetData())
	})
}
