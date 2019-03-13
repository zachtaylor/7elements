package game

import "github.com/zachtaylor/7elements"

type Request struct {
	Username string
	URI      string
	Data     vii.Json
}

func (r *Request) String() string {
	return r.Username + ":" + r.URI
}
