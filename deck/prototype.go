package deck

import (
	"strconv"

	"github.com/zachtaylor/7elements/card"
)

// Prototype is Deck list
type Prototype struct {
	ID    int
	Name  string
	User  string
	Cover int
	Cards card.Count
}

// NewPrototype creates a new Deck list
func NewPrototype(user string) *Prototype {
	return &Prototype{
		User:  user,
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

// Data returns a representation of this Prototype as type map[string]any
func (proto *Prototype) Data() map[string]any {
	cardsJSON := map[string]any{}
	size := 0
	for k, v := range proto.Cards {
		cardsJSON[strconv.FormatInt(int64(k), 10)] = v
		size += v
	}
	return map[string]any{
		"id":    proto.ID,
		"name":  proto.Name,
		"size":  size,
		"cover": proto.Cover,
		"cards": cardsJSON,
	}
}

// Prototypes is a set of Deck lists
type Prototypes map[int]*Prototype

// Data returns a representation of these Deck lists as type fmt.Stringer
func (decks Prototypes) Data() map[string]any {
	json := map[string]any{}
	for _, deck := range decks {
		json[strconv.FormatInt(int64(deck.ID), 10)] = deck.Data()
	}
	return json
}

// // PrototypeService provides access to Prototypes
// type PrototypeService interface {
// 	Get(id int) (*Prototype, error)
// 	Insert(p *Prototype) error
// 	Delete(id int) error
// }
