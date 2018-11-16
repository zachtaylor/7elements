package vii

import (
	"fmt"

	"ztaylor.me/js"
)

type Power struct {
	Id       int
	Text     string
	Trigger  string
	UsesTurn bool
	UsesKill bool
	Costs    ElementMap
	Target   string
	Script   string
}

func NewPower() *Power {
	return &Power{
		Costs: ElementMap{},
	}
}

func (p Power) Copy() *Power {
	return &Power{
		Id:       p.Id,
		Text:     p.Text,
		Costs:    p.Costs.Copy(),
		Target:   p.Target,
		UsesTurn: p.UsesTurn,
		UsesKill: p.UsesKill,
		Script:   p.Script,
	}
}

func (p *Power) Json() js.Object {
	return js.Object{
		"id":       p.Id,
		"text":     p.Text,
		"costs":    p.Costs,
		"trigger":  p.Trigger,
		"target":   p.Target,
		"usesturn": p.UsesTurn,
		"useskill": p.UsesKill,
	}
}

type Powers map[int]*Power

func NewPowers() Powers {
	return Powers{}
}

func (p Powers) Copy() Powers {
	cp := NewPowers()
	for k, v := range p {
		cp[k] = v.Copy()
	}
	return cp
}

func (powers Powers) Json() js.Object {
	json := js.Object{}
	for id, p := range powers {
		json[fmt.Sprint(id)] = p.Json()
	}
	return json
}
