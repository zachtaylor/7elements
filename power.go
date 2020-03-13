package vii

import (
	"github.com/zachtaylor/7elements/element"
	"ztaylor.me/cast"
)

type Power struct {
	Id       int
	Text     string
	Trigger  string
	UsesTurn bool
	UsesKill bool
	Costs    element.Count
	Target   string
	Script   string
}

func NewPower() *Power {
	return &Power{
		Costs: element.Count{},
	}
}

func (p *Power) Copy() *Power {
	return &Power{
		Id:       p.Id,
		Text:     p.Text,
		Trigger:  p.Trigger,
		Costs:    p.Costs.Copy(),
		Target:   p.Target,
		UsesTurn: p.UsesTurn,
		UsesKill: p.UsesKill,
		Script:   p.Script,
	}
}

func (p *Power) JSON() cast.JSON {
	return cast.JSON{
		"id":       p.Id,
		"text":     p.Text,
		"costs":    p.Costs.JSON(),
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

func (powers Powers) GetTrigger(name string) []*Power {
	ps := make([]*Power, 0)
	for _, p := range powers {
		if p.Trigger == name {
			ps = append(ps, p)
		}
	}
	return ps
}

func (powers Powers) JSON() []cast.JSON {
	if powers == nil {
		return nil
	}
	json := make([]cast.JSON, 0)
	for _, power := range powers {
		json = append(json, power.JSON())
	}
	return json
}
