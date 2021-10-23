package game

import (
	"time"
)

type State struct {
	id     string
	Phase  Phaser
	Timer  time.Duration
	Reacts map[string]string
	Stack  *State
}

func NewState(timeout time.Duration, phase Phaser) *State {
	return &State{
		Phase:  phase,
		Timer:  timeout,
		Reacts: make(map[string]string),
	}
}
func (t *T) NewState(timeout time.Duration, phase Phaser) *State { return NewState(timeout, phase) }

func (t *State) ID() string { return t.id }

func (t *State) String() string {
	return `state.T{#` + t.id + `(` + t.Phase.Name() + `:` + t.Phase.Seat() + `)` + `}`
}

// Data returns a a representation of game state
func (t *State) Data() map[string]interface{} {
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
