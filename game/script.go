package game

// ScriptFunc sets a func pointer for game code
//
// When called by the engine, scope can refer differently
//
// - me may be *game.Token or *game.Card
//
// - args=nil when Power.Target="self"
//
// - returns Staters which create new States to stack
type ScriptFunc func(g *T, s *Seat, me interface{}, args []interface{}) ([]Stater, error)

// Scripts is the injection point
var Scripts = make(map[string]ScriptFunc)
