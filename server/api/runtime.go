package api

import (
	"net/http"
	"sync"

	vii "github.com/zachtaylor/7elements"
	"github.com/zachtaylor/7elements/chat"
	"github.com/zachtaylor/7elements/game"
	"ztaylor.me/http/session"
)

type Runtime struct {
	Root       *vii.Runtime
	Games      game.Service
	Chat       chat.Service
	Salt       string
	Sessions   session.Service
	FileSystem http.FileSystem
	Ping       *Ping
}

// Ping measures the connected websockets
type Ping struct {
	sync.Mutex
	count int
}

func (p *Ping) Add() {
	p.Lock()
	p.count++
	p.Unlock()
}
func (p *Ping) Get() int {
	return p.count
}
func (p *Ping) Remove() {
	p.Lock()
	p.count--
	p.Unlock()
}

// func NewRuntime(root *vii.Runtime)
