package game

import "ztaylor.me/cast"

type State struct {
	id     string
	Timer  cast.Duration
	Reacts map[string]string
	R      Stater
	Stack  *State
}

func (s *State) ID() string {
	return s.id
}

func (s *State) Name() string {
	return s.R.Name()
}

func (s *State) String() string {
	return `game.State{` + s.Print() + `}`
}

// Print returns a detailed compressed string representation
func (s *State) Print() string {
	return s.R.Name() + `#` + s.id + `(` + s.R.Seat() + `)`
}

// JSON returns a a representation of game state
func (s *State) JSON() cast.JSON {
	reactsJSON := cast.JSON{}
	for k, v := range s.Reacts {
		reactsJSON[k] = v
	}
	return cast.JSON{
		"id":     s.ID(),
		"name":   s.R.Name(),
		"seat":   s.R.Seat(),
		"data":   s.R.JSON(),
		"timer":  int(s.Timer.Seconds()),
		"reacts": reactsJSON,
	}
}

// Stater is a behavior of a State
type Stater interface {
	// Name is the refferential name of the game State
	Name() string
	// Seat is the priority holder
	Seat() string
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
