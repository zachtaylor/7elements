package ctypes

type CardType byte

var Null CardType = 0
var Spell CardType = 1
var Body CardType = 2
var Item CardType = 3

var CardTypes = []CardType{Null, Spell, Body, Item}

func (cardtype CardType) String() string {
	if cardtype == Null {
		return "null"
	} else if cardtype == Spell {
		return "spell"
	} else if cardtype == Body {
		return "body"
	} else if cardtype == Item {
		return "item"
	}
	return "error"
}
