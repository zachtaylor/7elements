package cssbuilder

import (
	"net/http"
)

var Handler = http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	if content == "" {
		CreateContent()
	}

	w.Header().Set("Content-Type", "text/css")
	w.Write([]byte(content))
})
