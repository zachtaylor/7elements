package server

import (
	"net/http"

	"github.com/zachtaylor/7elements/server/internal"
	"taylz.io/http/handler"
)

func Start(server internal.Server, port string) {
	http.ListenAndServe(port, server)
}

func StartTLS(server internal.Server, cert string, key string) {
	go http.ListenAndServe(":80", handler.RedirectHTTPS)
	http.ListenAndServeTLS(":443", cert, key, server)
}
