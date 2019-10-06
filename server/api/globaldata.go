package api

import (
	"net/http"

	"ztaylor.me/cast"
)

func GlobalDataHandler(rt *Runtime) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(cast.BytesS(rt.Root.JSON().String()))
	})
}
