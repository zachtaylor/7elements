// Package element provides 7 Elements' typing system
package element

// T is enum of the titular variety of thing
type T byte

// Nil is an unspecific element id
const Nil T = 0

// White is an element id
const White T = 1

// Red is an element id
const Red T = 2

// Yellow is an element id
const Yellow T = 3

// Green is an element id
const Green T = 4

// Blue is an element id
const Blue T = 5

// Violet is an element id
const Violet T = 6

// Black is an element id
const Black T = 7

// Index is used to get an element from the corresponding enum value
var Index = [8]T{Nil, White, Red, Yellow, Green, Blue, Violet, Black}

func (e T) String() string {
	switch e {
	case White:
		return "white"
	case Red:
		return "red"
	case Yellow:
		return "yellow"
	case Green:
		return "green"
	case Blue:
		return "blue"
	case Violet:
		return "violet"
	case Black:
		return "black"
	default:
		return "nil"
	}
}

// ToChar returns the character byte used to encode this element as a capital letter
func (e T) ToChar() byte {
	switch e {
	case White:
		return 'W'
	case Red:
		return 'R'
	case Yellow:
		return 'Y'
	case Green:
		return 'G'
	case Blue:
		return 'B'
	case Violet:
		return 'V'
	case Black:
		return 'A'
	default:
		return 'X'
	}
}
