package game

import "taylz.io/http/websocket"

type Request struct {
	Username string
	URI      string
	Data     websocket.MsgData
}

func (r *Request) String() string {
	return r.Username + ":" + r.URI
}
