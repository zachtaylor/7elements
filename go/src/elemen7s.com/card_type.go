package vii

type CardType int

const CTYPspell CardType = 1
const CTYPbody CardType = 2
const CTYPitem CardType = 3

var CardTypes = []CardType{0, CTYPspell, CTYPbody, CTYPitem}

func (cardtype CardType) String() string {
	switch cardtype {
	case CTYPspell:
		return "spell"
	case CTYPbody:
		return "body"
	case CTYPitem:
		return "item"
	default:
		return "error"
	}
}
