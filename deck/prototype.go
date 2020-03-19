package deck

import (
	"strings"

	"ztaylor.me/cast"
)

// Prototype is Deck list
type Prototype struct {
	ID      int
	Name    string
	CoverID int
	Cards   map[int]int
}

// NewPrototype creates a new Deck list
func NewPrototype() *Prototype {
	return &Prototype{
		Cards: make(map[int]int),
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
		cardsJSON[cast.StringI(int(k))] = v
	}
	return cast.JSON{
		"id":    proto.ID,
		"name":  proto.Name,
		"cover": "/img/card/" + cast.StringI(proto.CoverID) + ".jpg",
		"cards": cardsJSON,
	}
}

// Prototypes is a set of Deck lists
type Prototypes map[int]*Prototype

// JSON returns a representation of these Deck lists as type fmt.Stringer
func (decks Prototypes) JSON() cast.IStringer {
	json := make([]string, 0)
	keys := make([]int, len(decks))
	var i int
	for k := range decks {
		keys[i] = k
		i++
	}
	cast.SortInts(keys)
	for _, k := range keys {
		json = append(json, decks[k].JSON().String())
	}
	return cast.Stringer(`[` + strings.Join(json, ",") + `]`)
}

// PrototypeService provides access to Prototypes
type PrototypeService interface {
	GetAll() (Prototypes, error)
	Get(int) (*Prototype, error)
}
