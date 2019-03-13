package vii

type Body struct {
	Attack int
	Health int
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

func (b *Body) Json() Json {
	if b == nil {
		return nil
	}
	return Json{
		"attack": b.Attack,
		"health": b.Health,
	}
}
