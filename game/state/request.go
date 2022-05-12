package state

import "taylz.io/http/websocket"

type Request struct {
	Username string
	websocket.Message
}

func NewRequest(username, uri string, json map[string]any) *Request {
	return &Request{
		Username: username,
		Message: websocket.Message{
			URI:  uri,
			Data: json,
		},
	}
}

func (r *Request) String() string {
	return r.Username + ":" + r.URI
}
