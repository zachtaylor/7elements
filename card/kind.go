package card

import "errors"

var ErrInvalidKind = errors.New("invalid kind")

// Kind is an enum to classify Cards
type Kind byte

const (
	Nil Kind = iota
	Spell
	Being
	Item
)

// All is a []T used to get a Card Type from the corresponding enum value
var All = []Kind{0, Spell, Being, Item}

func (t Kind) String() string {
	switch t {
	case Spell:
		return "spell"
	case Being:
		return "being"
	case Item:
		return "item"
	default:
		return "error"
	}
}
