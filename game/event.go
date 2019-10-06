package game

import "ztaylor.me/cast"

// Event is a behavior of a State
type Event interface {
	// Name is the refferential name of the game event
	Name() string
	// Seat is the priority holder
	Seat() string
	// GetNext is called by the engine if state.Stack is unavailable
	GetNext(*T) Event
	// JSON() create a representation of this GameState extra data
	JSON() cast.JSON
}

// ActivateEventer is an Event that is triggered by the engine activating an event for the 1st time
type ActivateEventer interface {
	// OnActivate is called by the engine when the event is mounted the 1st time
	OnActivate(*T) []Event
}

// ConnectEventer is an Event that is triggered when a player agent (re)connects
type ConnectEventer interface {
	// OnConnect is called by the engine whenever a Seat (re)joins, and when the
	// event re-mounts, and indicated by OnConnect(*T, nil)
	OnConnect(*T, *Seat)
}

// DisconnectEventer is an Event that is triggered when a player agent disconnects
type DisconnectEventer interface {
	// OnDisconnect is called by the engine whenever a Seat disconnects
	OnDisconnect(*T, *Seat)
}

// FinishEventer is an Event that is triggered by the engine finally resolving an event
type FinishEventer interface {
	// Finish is called by the engine exactly once, after all passes
	Finish(*T) []Event
}

// RequestEventer is an Event that is triggered by the engine when a Request targets the game state ID
type RequestEventer interface {
	// Request is called when a request is sent to this Event
	Request(*T, *Seat, cast.JSON)
}

// ZEvent type checks Event and all "Eventer types"
type ZEvent interface {
	Event
	ActivateEventer
	ConnectEventer
	FinishEventer
	RequestEventer
}
