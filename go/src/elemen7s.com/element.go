package vii

type Element byte

const ELEMnull Element = 0
const ELEMwhite Element = 1
const ELEMred Element = 2
const ELEMyellow Element = 3
const ELEMgreen Element = 4
const ELEMblue Element = 5
const ELEMviolet Element = 6
const ELEMblack Element = 7

var Elements = []Element{ELEMnull, ELEMwhite, ELEMred, ELEMyellow, ELEMgreen, ELEMblue, ELEMviolet, ELEMblack}

func (e Element) Char() string {
	switch e {
	case ELEMwhite:
		return "W"
	case ELEMred:
		return "R"
	case ELEMyellow:
		return "Y"
	case ELEMgreen:
		return "G"
	case ELEMblue:
		return "B"
	case ELEMviolet:
		return "V"
	case ELEMblack:
		return "K"
	default:
		return "X"
	}
}

func (e Element) String() string {
	switch e {
	case ELEMwhite:
		return "white"
	case ELEMred:
		return "red"
	case ELEMyellow:
		return "yellow"
	case ELEMgreen:
		return "green"
	case ELEMblue:
		return "blue"
	case ELEMviolet:
		return "violet"
	case ELEMblack:
		return "black"
	default:
		return ""
	}
}
