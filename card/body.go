package card

import "fmt"

// Body contains stats for Body Cards
type Body struct {
	Life   int
	Attack int
}

func (b *Body) Copy() *Body {
	if b == nil {
		return nil
	}
	return &Body{
		Life:   b.Life,
		Attack: b.Attack,
	}
}

func (b *Body) String() string {
	if b == nil {
		return "<nil>"
	}
	return fmt.Sprintf("Body{♥: %d, ♣: %d}", b.Life, b.Attack)
}

func (b *Body) JSON() map[string]any {
	if b == nil {
		return nil
	}
	return map[string]any{
		"life":   b.Life,
		"attack": b.Attack,
	}
}
