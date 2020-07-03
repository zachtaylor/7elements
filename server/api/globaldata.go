package api

import (
	"net/http"

	"github.com/zachtaylor/7elements/runtime"
	"ztaylor.me/cast"
)

func GlobalDataHandler(t *runtime.T) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(cast.BytesS(t.JSON().String()))
	})
}
