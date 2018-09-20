package server

import (
	"net/http"

	"ztaylor.me/http/handler"
	"ztaylor.me/log"
)

func Start(fs http.FileSystem, dbsalt string, port string) {
	server := New(fs, dbsalt)
	log.Add("port", port).Info("elemen7s server started!")
	server.ListenAndServe(port)
}

func StartTLS(fs http.FileSystem, dbsalt string, cert string, key string) {
	server := New(fs, dbsalt)
	log.Info("elemen7s server started!")
	go http.ListenAndServe(":80", handler.RedirectHTTPS)
	server.ListenAndServeTLS(cert, key)
}
