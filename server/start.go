package server

import (
	"net/http"
	zhttp "ztaylor.me/http"
	"ztaylor.me/log"
)

func Start(port string) {
	log.Add("port", port).Info("elemen7s server started!")
	zhttp.Start(port)
}

func StartTLS(cert string, key string) {
	go http.ListenAndServe(":80", http.HandlerFunc(redirectHttps))
	zhttp.StartTLS(cert, key)
}

func redirectHttps(w http.ResponseWriter, req *http.Request) {
	http.Redirect(w, req,
		"https://"+req.Host+req.URL.String(),
		http.StatusMovedPermanently)
}
