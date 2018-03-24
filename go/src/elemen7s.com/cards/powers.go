package cards

import (
	"elemen7s.com"
	"fmt"
	"ztaylor.me/js"
)

type Powers map[int]*Power

type Power struct {
	Id       int
	Costs    vii.ElementMap
	UsesTurn bool
	Script   string
}

func NewPowers() Powers {
	return Powers{}
}

func NewPower() *Power {
	return &Power{
		Costs: vii.ElementMap{},
	}
}

func (p Powers) Copy() Powers {
	cp := NewPowers()
	for k, v := range p {
		cp[k] = v.Copy()
	}
	return cp
}

func (p Power) Copy() *Power {
	return &Power{
		Id:       p.Id,
		Costs:    p.Costs.Copy(),
		UsesTurn: p.UsesTurn,
		Script:   p.Script,
	}
}

func (powers Powers) Json() js.Object {
	json := js.Object{}
	for id, p := range powers {
		json[fmt.Sprint(id)] = p.Json()
	}
	return json
}

func (p *Power) Json() js.Object {
	return js.Object{
		"id":       p.Id,
		"costs":    p.Costs.Copy(),
		"usesturn": p.UsesTurn,
		"script":   p.Script,
	}
}

func (powers Powers) JsonWithText(text *vii.CardText) js.Object {
	json := js.Object{}
	for id, p := range powers {
		json[fmt.Sprint(id)] = p.JsonWithText(text)
	}
	return json
}

func (p *Power) JsonWithText(text *vii.CardText) js.Object {
	return js.Object{
		"id":          p.Id,
		"costs":       p.Costs.Copy(),
		"usesturn":    p.UsesTurn,
		"script":      p.Script,
		"description": text.Powers[p.Id],
	}
}
