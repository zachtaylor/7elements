package card

import "ztaylor.me/cast"

// Body contains stats for Body Cards
type Body struct {
	Attack int
	Health int
}

// Copy returns a copy of this Body
func (b *Body) Copy() *Body {
	if b == nil {
		return nil
	}
	return &Body{
		Attack: b.Attack,
		Health: b.Health,
	}
}

// JSON returns a representation of this Body as type fmt.Stringer
func (b *Body) JSON() cast.IStringer {
	if b == nil {
		return cast.Stringer(`null`)
	}
	return cast.JSON{
		"attack": b.Attack,
		"health": b.Health,
	}
}
