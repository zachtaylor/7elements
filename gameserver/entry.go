package gameserver

import (
	"github.com/zachtaylor/7elements/deck"
	"taylz.io/http/websocket"
)

// Entry contains user pointers and info per game participant
type Entry struct {
	Deck   *deck.T
	Writer websocket.Writer
}
