package vii

import (
	"fmt"
	"ztaylor.me/js"
)

type Power struct {
	Id    int
	Costs ElementMap
	Target
	UsesTurn bool
	Script   string
}

func NewPower() *Power {
	return &Power{
		Costs: ElementMap{},
	}
}

func (p Power) Copy() *Power {
	return &Power{
		Id:     p.Id,
		Costs:  p.Costs.Copy(),
		Script: p.Script,
	}
}

func (p *Power) Json() js.Object {
	return js.Object{
		"id":     p.Id,
		"costs":  p.Costs,
		"target": p.Target,
		"script": p.Script,
	}
}

func (p *Power) JsonWithText(text *CardText) js.Object {
	json := p.Json()
	json["description"] = text.Powers[p.Id]
	return json
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

func (powers Powers) JsonWithText(text *CardText) js.Object {
	json := js.Object{}
	for id, p := range powers {
		json[fmt.Sprint(id)] = p.JsonWithText(text)
	}
	return json
}
