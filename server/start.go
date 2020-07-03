package server

import (
	"net/http"

	"github.com/zachtaylor/7elements/server/runtime"
	"ztaylor.me/http/handler"
)

func Start(runtime *runtime.T, port string) {
	server := Routes(runtime)
	rt.Root.Logger.New().Add("port", port).Info("elemen7s server started!")
	http.ListenAndServe(port, server)
}

func StartTLS(runtime *runtime.T, cert string, key string) {
	server := Routes(runtime)
	rt.Root.Logger.New().Info("elemen7s server started!")
	go http.ListenAndServe(":80", handler.RedirectHTTPS)
	http.ListenAndServeTLS(":443", cert, key, server)
}
