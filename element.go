package vii

import "ztaylor.me/cast"

// Element is enum of the titular variety of thing
type Element byte

const ELEMnil Element = 0
const ELEMwhite Element = 1
const ELEMred Element = 2
const ELEMyellow Element = 3
const ELEMgreen Element = 4
const ELEMblue Element = 5
const ELEMviolet Element = 6
const ELEMblack Element = 7

// Elements is a []Element which can indexed by Element number, inverse of Element enum value
var Elements = []Element{ELEMnil, ELEMwhite, ELEMred, ELEMyellow, ELEMgreen, ELEMblue, ELEMviolet, ELEMblack}

// Char returns the character used to encode this Element as a capital letter
func (e Element) Char() byte {
	switch e {
	case ELEMwhite:
		return 'W'
	case ELEMred:
		return 'R'
	case ELEMyellow:
		return 'Y'
	case ELEMgreen:
		return 'G'
	case ELEMblue:
		return 'B'
	case ELEMviolet:
		return 'V'
	case ELEMblack:
		return 'A'
	default:
		return 'X'
	}
}

// ParseElement parses a byte for `Element` encoding, rough inverse of `Element.Char()`
//
// returns `Element` encoded by the byte
//
// returns "activate" bool
//
// returns error!=nil if c fails to parse
func ParseElement(c byte) (Element, bool, error) {
	switch c {
	case 'w':
		return ELEMwhite, false, nil
	case 'W':
		return ELEMwhite, true, nil
	case 'r':
		return ELEMred, false, nil
	case 'R':
		return ELEMred, true, nil
	case 'y':
		return ELEMyellow, false, nil
	case 'Y':
		return ELEMyellow, true, nil
	case 'g':
		return ELEMgreen, false, nil
	case 'G':
		return ELEMgreen, true, nil
	case 'b':
		return ELEMblue, false, nil
	case 'B':
		return ELEMblue, true, nil
	case 'v':
		return ELEMviolet, false, nil
	case 'V':
		return ELEMviolet, true, nil
	case 'a':
		return ELEMblack, false, nil
	case 'A':
		return ELEMblack, true, nil
	case 'x':
		return ELEMnil, false, nil
	case 'X':
		return ELEMnil, true, nil
	default:
		return ELEMnil, false, cast.NewError(nil, "vii.ParseElement: ", c)
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
