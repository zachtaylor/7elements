package vii

type Element byte

const NullElement Element = 0
const White Element = 1
const Red Element = 2
const Yellow Element = 3
const Green Element = 4
const Blue Element = 5
const Violet Element = 6
const Black Element = 7

var Elements = []Element{NullElement, White, Red, Yellow, Green, Blue, Violet, Black}

func (e Element) Char() string {
	switch e {
	case NullElement:
		return "1"
	case White:
		return "w"
	case Red:
		return "r"
	case Yellow:
		return "y"
	case Green:
		return "g"
	case Blue:
		return "b"
	case Violet:
		return "v"
	case Black:
		return "k"
	default:
		return "x"
	}
}

func (e Element) String() string {
	switch e {
	case NullElement:
		return ""
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
		return "error"
	}
}
