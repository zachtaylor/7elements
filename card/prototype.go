package card

import (
	"strings"

	"github.com/zachtaylor/7elements/element"
	"github.com/zachtaylor/7elements/power"
	"ztaylor.me/cast"
)

// Prototype is a definition of a card
type Prototype struct {
	ID     int
	Name   string
	Text   string
	Type   Type
	Image  string
	Costs  element.Count
	Body   *Body
	Powers power.Set
}

// NewPrototype returns a new empty Card Prototype
func NewPrototype() *Prototype {
	return &Prototype{
		Costs:  element.Count{},
		Powers: power.NewSet(),
	}
}

// JSON returns a representation of this Prototype as type cast.JSON
func (c *Prototype) JSON() cast.JSON {
	return cast.JSON{
		"id":     c.ID,
		"image":  c.Image,
		"name":   c.Name,
		"text":   c.Text,
		"type":   c.Type.String(),
		"powers": c.Powers.JSON(),
		"costs":  c.Costs.JSON(),
		"body":   c.Body.JSON(),
	}
}

func (c *Prototype) String() string {
	return cast.StringN(`{`, c.ID, ` `, c.Name, `}`)
}

// Prototypes is a set of Prototype, mapped by ID number
type Prototypes map[int]*Prototype

// JSON returns a representation of these Prototypes as type fmt.Stringer
func (cards Prototypes) JSON() cast.IStringer {
	json := make([]string, 0)
	keys := make([]int, len(cards))

	var i int
	for k := range cards {
		keys[i] = k
		i++
	}
	cast.SortInts(keys)
	for _, id := range keys {
		json = append(json, cards[id].JSON().String())
	}
	return cast.Stringer(`[` + strings.Join(json, ",") + `]`)
}

// PrototypeService is used to acquire Card Prototypes
type PrototypeService interface {
	Start() error
	Get(cardid int) (*Prototype, error)
	GetAll() Prototypes
}
