package server

import (
	"net/http"
	"regexp"
)

type route struct {
	pattern *regexp.Regexp
	handler http.Handler
}

type RegexpHandler []*route

var Router = make(RegexpHandler, 0)

func Handler(pattern *regexp.Regexp, handler http.Handler) {
	Router = append(Router, &route{pattern, handler})
}

func HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	Router = append(Router, &route{regexp.MustCompile(pattern), http.HandlerFunc(handler)})
}

func (rh *RegexpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range *rh {
		if route.pattern.MatchString(r.URL.Path) {
			route.handler.ServeHTTP(w, r)
			return
		}
	}
	http.NotFound(w, r)
}
