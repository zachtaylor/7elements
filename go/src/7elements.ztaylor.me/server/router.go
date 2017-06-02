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

func addhandler(pattern *regexp.Regexp, handler http.Handler) {
	Router = append(Router, &route{pattern, handler})
}

func Handler(pattern string, handler http.Handler) {
	addhandler(regexp.MustCompile(pattern), handler)
}

func HandleFunc(pattern string, handler func(http.ResponseWriter, *http.Request)) {
	Handler(pattern, http.HandlerFunc(handler))
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
