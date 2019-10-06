package game

import "ztaylor.me/cast"

type Request struct {
	Username string
	URI      string
	Data     cast.JSON
}

func (r *Request) String() string {
	return r.Username + ":" + r.URI
}
