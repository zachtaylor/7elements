package vii

import "ztaylor.me/cast"

type Body struct {
	Attack int
	Health int
}

func (b *Body) JSON() cast.IStringer {
	if b == nil {
		return cast.Stringer(`null`)
	}
	return cast.JSON{
		"attack": b.Attack,
		"health": b.Health,
	}
}

func (b *Body) Copy() *Body {
	if b == nil {
		return nil
	}
	return &Body{
		Attack: b.Attack,
		Health: b.Health,
	}
}
