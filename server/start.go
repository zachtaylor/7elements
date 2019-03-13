package server

import (
	"net/http"

	"ztaylor.me/http/handlers"
	"ztaylor.me/http/sessions"
	"ztaylor.me/log"
)

func Start(fs http.FileSystem, sessions *sessions.Service, dbsalt string, port string) {
	server := Server(fs, sessions, dbsalt)
	log.Add("port", port).Info("elemen7s server started!")
	http.ListenAndServe(port, server)
}

func StartTLS(fs http.FileSystem, sessions *sessions.Service, dbsalt string, cert string, key string) {
	server := Server(fs, sessions, dbsalt)
	log.Info("elemen7s server started!")
	go http.ListenAndServe(":80", handlers.RedirectHTTPS)
	http.ListenAndServeTLS(":443", cert, key, server)
}
