package game

// type State struct {
// 	id     string
// 	Phase  Phaser
// 	Timer  time.Duration
// 	Reacts map[string]string
// 	Stack  *State
// }

// func NewState(timeout time.Duration, phase Phaser) *State {
// 	return &State{
// 		Phase:  phase,
// 		Timer:  timeout,
// 		Reacts: make(map[string]string),
// 	}
// }
// func (t *T) NewState(timeout time.Duration, phase Phaser) *State { return NewState(timeout, phase) }

// func (t *State) ID() string { return t.id }

// func (t *State) String() string {
// 	return `state.T{#` + t.id + `(` + t.Phase.Name() + `:` + t.Phase.Seat() + `)` + `}`
// }

// // JSON returns a a representation of game state data as map[string]any
// func (t *State) JSON() map[string]any {
// 	if t == nil {
// 		return nil
// 	}
// 	// reactsJSON := map[string]any{}
// 	// for k, v := range t.Reacts {
// 	// 	reactsJSON[k] = v
// 	// }
// 	var stack interface{}
// 	if t.Stack != nil {
// 		stack = t.Stack.id
// 	}
// 	return map[string]any{
// 		"id":     t.ID(),
// 		"seat":   t.Phase.Seat(),
// 		"name":   t.Phase.Name(),
// 		"data":   t.Phase.JSON(),
// 		"timer":  int(t.Timer.Seconds()),
// 		"reacts": t.Reacts, // reactsJSON,
// 		"stack":  stack,
// 	}
// }
