package elements

type Element byte

var Null Element = 0
var White Element = 1
var Blue Element = 2
var Green Element = 3
var Gold Element = 4
var Red Element = 5
var Indigo Element = 6
var Black Element = 7

var Elements = []Element{Null, White, Blue, Green, Gold, Red, Indigo, Black}

func (e Element) String() string {
	switch e {
	case Null:
		return ""
	case White:
		return "white"
	case Blue:
		return "blue"
	case Green:
		return "green"
	case Gold:
		return "gold"
	case Red:
		return "red"
	case Indigo:
		return "indigo"
	case Black:
		return "black"
	default:
		return "error"
	}
}
