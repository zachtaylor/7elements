package chat

import (
	"sync"
)

var dispatch = make(map[string]Channel)
var mu sync.Mutex

func GetChannel(name string) Channel {
	mu.Lock()
	if c := dispatch[name]; c != nil {
		mu.Unlock()
		return c
	} else {
		c := NewChannel(name)
		dispatch[name] = c
		mu.Unlock()
		return c
	}
}
