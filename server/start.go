package server

import (
	"net/http"

	"ztaylor.me/http/handler"
	"ztaylor.me/log"
)

func Start(fs http.FileSystem, port string) {
	server := New(fs)
	log.Add("port", port).Info("elemen7s server started!")
	server.ListenAndServe(port)
}

func StartTLS(fs http.FileSystem, cert string, key string) {
	server := New(fs)
	log.Info("elemen7s server started!")
	go http.ListenAndServe(":80", handler.RedirectHTTPS)
	server.ListenAndServeTLS(cert, key)
}
