package cards

import (
	"ztaylor.me/js"
)

type Body struct {
	Attack int
	Health int
}

func (body *Body) Copy() *Body {
	return &Body{
		Attack: body.Attack,
		Health: body.Health,
	}
}

func (b *Body) Json() js.Object {
	if b == nil {
		return nil
	}
	return js.Object{
		"attack": b.Attack,
		"health": b.Health,
	}
}
