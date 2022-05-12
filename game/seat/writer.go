package seat

import "taylz.io/http/websocket"

type Writer interface {
	Name() string
	Done() <-chan struct{}
	WriteMessage(*websocket.Message) error
	WriteMessageData([]byte) error
}
