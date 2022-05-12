package server

import (
	"net/http"

	"github.com/zachtaylor/7elements/server/runtime"
	"taylz.io/http/handler"
)

func preStart(runtime *runtime.T) {
	runtime.Log().Info("7tcg server starting")
	setRoutes(runtime)
}

func Start(runtime *runtime.T, port string) {
	preStart(runtime)
	http.ListenAndServe(port, runtime.Handler)
}

func StartTLS(runtime *runtime.T, cert string, key string) {
	preStart(runtime)
	go http.ListenAndServe(":80", handler.RedirectHTTPS)
	http.ListenAndServeTLS(":443", cert, key, runtime.Handler)
}
