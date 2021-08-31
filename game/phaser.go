package game

import "github.com/zachtaylor/7elements/game/seat"

type Phaser interface {
	// String is Stringer
	String() string
	// Seat is the priority seat
	Seat() string
	// Name is the refferential name of the game State
	Name() string
	// GetNext is called by the engine if state.Stack is unavailable
	GetNext(*T) Phaser
	// JSON() create a representation of this State's extra data
	Data() map[string]interface{}
}

// OnActivatePhaser is a Phaser that is triggered by the engine activating a State only the 1st time
type OnActivatePhaser interface {
	// OnActivate is called by the engine exactly once, when the State is mounted the 1st time
	OnActivate(*T) []Phaser
}

// OnConnectPhaser is a Phaser that is triggered when a player agent (re)connects
type OnConnectPhaser interface {
	// OnConnect is called by the engine whenever a seat.T (re)joins, and when the
	// Phaser re-mounts, as indicated by OnConnect(*T, nil)
	OnConnect(*T, *seat.T)
}

// OnDisconnectPhaser is a Phaser that is triggered when a player agent disconnects
type OnDisconnectPhaser interface {
	// OnDisconnect is called by the engine whenever a seat.T disconnects
	OnDisconnect(*T, *seat.T)
}

// OnFinishPhaser is a Phaser that is triggered by the engine finally resolving a State
type OnFinishPhaser interface {
	// OnFinish is called by the engine exactly once, after all responses, or timeout
	OnFinish(*T) []Phaser
}

// OnRequestPhaser is a Phaser that is triggered by the engine when a Request targets the game state ID
type OnRequestPhaser interface {
	// Request is called when a request is sent to this State
	OnRequest(*T, *seat.T, map[string]interface{})
}
