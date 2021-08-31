package game

import (
	"time"

	"taylz.io/keygen"
)

type State struct {
	id     string
	Phase  Phaser
	Timer  time.Duration
	Reacts map[string]string
	Stack  *State
}

func NewState(timeout time.Duration, phase Phaser) *State {
	id := keygen.New(4)

	return &State{
		id:     id,
		Phase:  phase,
		Timer:  timeout,
		Reacts: make(map[string]string),
	}
}

func (t *State) ID() string { return t.id }

func (t *State) String() string {
	return `state.T{#` + t.id + `(` + t.Phase.Name() + `:` + t.Phase.Seat() + `)` + `}`
}

// MessageData returns a a representation of game state
func (t *State) MessageData() map[string]interface{} {
	if t == nil {
		return nil
	}
	// reactsJSON := map[string]interface{}{}
	// for k, v := range t.Reacts {
	// 	reactsJSON[k] = v
	// }
	var stack interface{}
	if t.Stack != nil {
		stack = t.Stack.id
	}
	return map[string]interface{}{
		"id":     t.ID(),
		"seat":   t.Phase.Seat(),
		"name":   t.Phase.Name(),
		"data":   t.Phase.Data(),
		"timer":  int(t.Timer.Seconds()),
		"reacts": t.Reacts, // reactsJSON,
		"stack":  stack,
	}
}
