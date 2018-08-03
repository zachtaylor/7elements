package chat

import (
	"sync"
	"ztaylor.me/keygen"
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

func CreateChannel() Channel {
	mu.Lock()
	key := "all"
	for ; dispatch[key] != nil; key = keygen.NewVal() {
	}
	dispatch[key] = NewChannel(key)
	mu.Unlock()
	return dispatch[key]
}
