package game

import "github.com/zachtaylor/7elements"

// Event is a behavior of a State
type Event interface {
	// Name is the refferential name of the game event
	Name() string
	// GetNext is called by the engine after GetStack, if available
	GetNext(*T) *State
	// Json() create a representation of this GameState extra data
	Json(*T) vii.Json
}

// ActivateEventer is an Event that is triggered by the engine (re)activating an event
type ActivateEventer interface {
	// OnActivate is called by the engine when the event timer (re)starts
	OnActivate(*T)
}

// ConnectEventer is an Event that is triggered when a player agent (re)connects
type ConnectEventer interface {
	// OnConnect is called by the engine whenever a Seat (re)joins
	OnConnect(*T, *Seat)
}

// RequestEventer is an Event that is triggered by the engine when a Request targets the game state ID
type RequestEventer interface {
	// Request is called when a request is sent to this Event
	Request(*T, *Seat, vii.Json)
}

// FinishEventer is an Event that is triggered by the engine finally resolving an event
type FinishEventer interface {
	// Finish is called by the engine exactly once, after all passes
	Finish(*T)
}

// StackEventer is an Event that requires unwind handling
type StackEventer interface {
	// GetStack returns the stacked event
	GetStack(*T) *State
}

// ZEvent type checks Event and all "Eventer types"
type ZEvent interface {
	Event
	ActivateEventer
	ConnectEventer
	FinishEventer
	RequestEventer
	StackEventer
}
