package game

import "ztaylor.me/cast"

// Stater is a behavior of a State
type Stater interface {
	// Seat is the priority seat
	Seat() string
	// Name is the refferential name of the game State
	Name() string
	// GetNext is called by the engine if state.Stack is unavailable
	GetNext(*T) Stater
	// JSON() create a representation of this State's extra data
	JSON() cast.JSON
}

// ActivateStater is a Stater that is triggered by the engine activating a State only the 1st time
type ActivateStater interface {
	// OnActivate is called by the engine exactly once, when the State is mounted the 1st time
	OnActivate(*T) []Stater
}

// ConnectStater is a Stater that is triggered when a player agent (re)connects
type ConnectStater interface {
	// OnConnect is called by the engine whenever a Seat (re)joins, and when the
	// Stater re-mounts, as indicated by OnConnect(*T, nil)
	OnConnect(*T, *Seat)
}

// DisconnectStater is a Stater that is triggered when a player agent disconnects
type DisconnectStater interface {
	// OnDisconnect is called by the engine whenever a Seat disconnects
	OnDisconnect(*T, *Seat)
}

// FinishStater is a Stater that is triggered by the engine finally resolving a State
type FinishStater interface {
	// Finish is called by the engine exactly once, after all responses, or timeout
	Finish(*T) []Stater
}

// RequestStater is a Stater that is triggered by the engine when a Request targets the game state ID
type RequestStater interface {
	// Request is called when a request is sent to this State
	Request(*T, *Seat, cast.JSON)
}
