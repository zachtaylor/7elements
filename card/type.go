package card

// Type is an enum to classify Cards
type Type byte

// SpellType is a Type that idenfities Spell Cards
const SpellType Type = 1

// BodyType is a Type that identifies Body Cards
const BodyType Type = 2

// ItemType is a Type that identifies Item Cards
const ItemType Type = 3

// Types is a []Type used to get a Card Type from the corresponding enum value
var Types = []Type{0, SpellType, BodyType, ItemType}

func (t Type) String() string {
	switch t {
	case SpellType:
		return "spell"
	case BodyType:
		return "being"
	case ItemType:
		return "item"
	default:
		return "error"
	}
}
