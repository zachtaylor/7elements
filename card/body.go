package card

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

// Data returns a representation of this Body as type fmt.Stringer
func (b *Body) Data() map[string]interface{} {
	if b == nil {
		return nil
	}
	return map[string]interface{}{
		"attack": b.Attack,
		"health": b.Health,
	}
}
