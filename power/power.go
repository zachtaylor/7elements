package power

import (
	"github.com/zachtaylor/7elements/element"
	"ztaylor.me/cast"
)

// T is a Power, a container for scripting
type T struct {
	ID       int
	Text     string
	Trigger  string
	UsesTurn bool
	UsesKill bool
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
		UsesKill: p.UsesKill,
		Costs:    p.Costs.Copy(),
		Target:   p.Target,
		Script:   p.Script,
	}
}

// JSON returns a representation of this Power as type cast.JSON
func (p *T) JSON() cast.JSON {
	return cast.JSON{
		"id":       p.ID,
		"text":     p.Text,
		"costs":    p.Costs.JSON(),
		"trigger":  p.Trigger,
		"target":   p.Target,
		"usesturn": p.UsesTurn,
		"useskill": p.UsesKill,
	}
}
