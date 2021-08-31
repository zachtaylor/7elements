package power

import "github.com/zachtaylor/7elements/element"

// T is a Power, a container for scripting
type T struct {
	ID       int
	Text     string
	Trigger  string
	UsesTurn bool
	UsesLife bool
	Costs    element.Count
	Target   string
	Script   string
}

// New creates a new Power
func New() *T {
	return &T{
		Costs: element.Count{},
	}
}

// Copy returns a copy of this Power
func (p *T) Copy() *T {
	return &T{
		ID:       p.ID,
		Text:     p.Text,
		Trigger:  p.Trigger,
		UsesTurn: p.UsesTurn,
		UsesLife: p.UsesLife,
		Costs:    p.Costs.Copy(),
		Target:   p.Target,
		Script:   p.Script,
	}
}

// Data returns a representation of this Power as type websocket.MsgData
func (p *T) Data() map[string]interface{} {
	return map[string]interface{}{
		"id":       p.ID,
		"text":     p.Text,
		"costs":    p.Costs.JSON(),
		"trigger":  p.Trigger,
		"target":   p.Target,
		"usesturn": p.UsesTurn,
		"useslife": p.UsesLife,
	}
}
