package chat

import (
	"sync"
	"ztaylor.me/http"
)

var dispatch = make(map[string]Channel)
var mu sync.Mutex

func GetChannel(name string) Channel {
	mu.Lock()
	if c := dispatch[name]; c != nil {
		mu.Unlock()
		return c
	} else {
		c := &channel{
			name:    name,
			sockets: make(map[string]*http.Socket),
			history: make([]*Message, CHANNEL_BUFFER_SIZE),
		}
		dispatch[name] = c
		mu.Unlock()
		return c
	}
}
