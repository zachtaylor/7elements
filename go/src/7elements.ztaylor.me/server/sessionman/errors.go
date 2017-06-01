package sessionman

import (
	"errors"
)

const pkglog = "sessionman: "

var Errors = struct {
	NoCookie         error
	CookieParse      error
	InvalidSessionId error
}{errors.New(pkglog + "no cookie"), errors.New(pkglog + "cookie parse"), errors.New(pkglog + "invalid session id")}
