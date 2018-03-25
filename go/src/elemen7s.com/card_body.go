package vii

import (
	"ztaylor.me/js"
)

type CardBody struct {
	Attack int
	Health int
}

func NewCardBody() *CardBody {
	return &CardBody{}
}

func (b *CardBody) Copy() *CardBody {
	if b == nil {
		return nil
	}
	return &CardBody{
		Attack: b.Attack,
		Health: b.Health,
	}
}

func (b *CardBody) Json() js.Object {
	if b == nil {
		return nil
	}
	return js.Object{
		"attack": b.Attack,
		"health": b.Health,
	}
}
