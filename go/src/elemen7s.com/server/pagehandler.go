package server

import (
	"net/http"
	"strings"
	"ztaylor.me/env"
	zhttp "ztaylor.me/http"
	"ztaylor.me/log"
)

var fileserver = http.FileServer(http.Dir("www"))

var PageHandler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	log := log.Add("Addr", r.RemoteAddr)
	if host := env.Default("HOST", ""); r.RequestURI == "/" && host != "" && host != r.Host && strings.Index(r.Host, "o") > 0 {
		log.Add("Hostname", host).Add("RequestHost", r.Host).Warn("page: redirect hostname")
		http.Redirect(w, r, "https://"+host, 307)
		return
	}

	if i := strings.Index(r.RequestURI, "#"); r.RequestURI == "/" || (i > 0 && r.RequestURI[0:i] == "/") {
		remoteip := r.RemoteAddr[0:strings.LastIndex(r.RemoteAddr, ":")]
		zhttp.TrackAddr(remoteip)
		if session, _ := zhttp.ReadRequestCookie(r); session != nil {
			session.Refresh()
			zhttp.Track(session.Username, remoteip)
		}
	}

	log.Debug(r.Method, " ", r.Host, r.RequestURI)

	fileserver.ServeHTTP(w, r)
})
