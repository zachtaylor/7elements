// Package notify provides a simple cache to place user notices
package notify // import "github.com/zachtaylor/7tcg/notify"

import (
	"sync"
	"time"
)

var cache = make(map[string][]*Notification)
var mu sync.Mutex

type Notification struct {
	Key     string
	Message string
}

func Send(username, key, msg string) {
	mu.Lock()
	n := &Notification{key, msg}
	if ns := cache[username]; ns != nil {
		cache[username] = append(ns, n)
	} else {
		cache[username] = []*Notification{n}
	}
	mu.Unlock()
}

func Get(username string) []*Notification {
	mu.Lock()
	r := make([]*Notification, len(cache[username]))
	if ns := cache[username]; ns != nil {
		copy(r, ns)
	}
	mu.Unlock()
	return r
}
