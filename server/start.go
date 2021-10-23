package server

import (
	"net/http"

	"github.com/zachtaylor/7elements/server/runtime"
	"taylz.io/http/handler"
)

func Start(runtime *runtime.T, fs http.FileSystem, port string) {
	Routes(runtime, fs)
	runtime.Log().Add("port", port).Info("elemen7s server started!")
	http.ListenAndServe(port, runtime.Handler)
}

func StartTLS(runtime *runtime.T, fs http.FileSystem, cert string, key string) {
	Routes(runtime, fs)
	runtime.Log().Info("elemen7s server started!")
	go http.ListenAndServe(":80", handler.RedirectHTTPS)
	http.ListenAndServeTLS(":443", cert, key, runtime.Handler)
}
