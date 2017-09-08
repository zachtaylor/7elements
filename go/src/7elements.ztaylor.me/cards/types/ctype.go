package ctypes

type CardType byte

var Null CardType = 0
var Spell CardType = 1
var Body CardType = 2
var Item CardType = 3

var CardTypes = []CardType{Null, Spell, Body, Item}

func (cardtype CardType) String() string {
	if cardtype == Null {
		return "NULL"
	} else if cardtype == Spell {
		return "SPELL"
	} else if cardtype == Body {
		return "BODY"
	} else if cardtype == Item {
		return "ITEM"
	}
	return "error"
}
