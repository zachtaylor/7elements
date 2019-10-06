package server

import (
	"net/http"

	"github.com/zachtaylor/7elements/server/api"
	"ztaylor.me/http/handler"
)

func Start(rt *api.Runtime, port string) {
	server := Routes(rt)
	rt.Root.Logger.New().Add("port", port).Info("elemen7s server started!")
	http.ListenAndServe(port, server)
}

func StartTLS(rt *api.Runtime, cert string, key string) {
	server := Routes(rt)
	rt.Root.Logger.New().Info("elemen7s server started!")
	go http.ListenAndServe(":80", handler.RedirectHTTPS)
	http.ListenAndServeTLS(":443", cert, key, server)
}
