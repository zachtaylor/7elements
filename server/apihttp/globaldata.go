package apihttp

import (
	"net/http"

	"github.com/zachtaylor/7elements/server/runtime"
)

func GlobalDataHandler(rt *runtime.T) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write(rt.GlobalData())
	})
}
