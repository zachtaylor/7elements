package server

import (
	"net/http"
	"strings"

	"ztaylor.me/env"
	zhttp "ztaylor.me/http"
	"ztaylor.me/log"
)

var fileserver = http.FileServer(http.Dir(env.Default("WWW_PATH", "www/")))

var PageHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log := log.Add("Addr", r.RemoteAddr)
	if host := env.Default("HOST", ""); r.URL.Path == "/" && host != "" && host != r.Host {
		log.Add("Hostname", host).Add("RequestHost", r.Host).Warn("page: redirect hostname")
		http.Redirect(w, r, "https://"+host, 307)
		return
	}

	path := r.URL.Path
	if i := strings.LastIndex(path, "/"); i > 1 {
		path = path[i:]
	}
	if !strings.Contains(path, ".") {
		indexRequest, _ := http.NewRequest(r.Method, "/", nil)
		fileserver.ServeHTTP(w, indexRequest)
		log.Info("page: single page app response")

		if addr := r.RemoteAddr[0:strings.LastIndex(r.RemoteAddr, ":")]; addr != `[::1]` {
			zhttp.TrackAddr(addr)
		}

		return
	}

	log.Debug("page: ", r.Method, " ", r.Host, r.RequestURI)

	fileserver.ServeHTTP(w, r)
})
