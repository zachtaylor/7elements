package server

import (
	"net/http"

	zhttp "ztaylor.me/http"
	"ztaylor.me/log"
)

func Start(port string) {
	log.Add("port", port).Info("elemen7s server started!")
	zhttp.StartServer(port, Server)
}

func StartTLS(cert string, key string) {
	go zhttp.StartServer(":80", http.HandlerFunc(redirectHttps))
	zhttp.StartTLS(cert, key, Server)
}

func redirectHttps(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req,
		"https://"+req.Host+req.URL.String(),
		http.StatusMovedPermanently)
}
