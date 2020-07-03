package deck

import (
	"github.com/zachtaylor/7elements/card"
	"ztaylor.me/cast"
)

// Prototype is Deck list
type Prototype struct {
	ID    int
	Name  string
	User  string
	Cover int
	Wins  int
	Loss  int
	Cards card.Count
}

// NewPrototype creates a new Deck list
func NewPrototype() *Prototype {
	return &Prototype{
		Cards: make(card.Count),
	}
}

// Count returns the total number of cards listed
func (proto *Prototype) Count() int {
	total := 0
	for _, count := range proto.Cards {
		total += count
	}
	return total
}

// JSON returns a representation of this Prototype as type cast.JSON
func (proto *Prototype) JSON() cast.JSON {
	cardsJSON := cast.JSON{}
	for k, v := range proto.Cards {
		cardsJSON[cast.StringI(k)] = v
	}
	return cast.JSON{
		"id":    proto.ID,
		"name":  proto.Name,
		"cover": "/img/card/" + cast.StringI(proto.Cover) + ".jpg",
		"cards": cardsJSON,
	}
}

// Prototypes is a set of Deck lists
type Prototypes map[int]*Prototype

// JSON returns a representation of these Deck lists as type fmt.Stringer
func (decks Prototypes) JSON() cast.JSON {
	json := cast.JSON{}
	for _, deck := range decks {
		json[cast.StringI(deck.ID)] = deck.JSON()
	}
	return json
}

// PrototypeService provides access to Prototypes
type PrototypeService interface {
	GetUser(user string) (Prototypes, error)
	UpdateName(id, newname string) error
	Insert(p *Prototype) error
	Delete(id string) error
}
