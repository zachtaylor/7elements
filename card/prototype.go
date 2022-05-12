package card

import (
	"sort"
	"strconv"

	"github.com/zachtaylor/7elements/element"
	"github.com/zachtaylor/7elements/power"
)

// Prototype is a definition of a card
type Prototype struct {
	ID     int
	Kind   Kind
	Name   string
	Text   string
	Costs  element.Count
	Body   *Body
	Powers power.Set
}

// NewPrototype returns a new empty Card Prototype
func NewPrototype() *Prototype {
	return &Prototype{}
}

func (proto *Prototype) String() string {
	if proto == nil {
		return `<nil>`
	}
	return `card.Prototype{` + strconv.FormatInt(int64(proto.ID), 10) + ` ` + proto.Name + `}`
}

// Data returns a representation of this Prototype as type map[string]any
func (proto *Prototype) Data() map[string]any {
	return map[string]any{
		"id":     proto.ID,
		"name":   proto.Name,
		"text":   proto.Text,
		"type":   proto.Kind.String(),
		"powers": proto.Powers.JSON(),
		"costs":  proto.Costs.JSON(),
		"body":   proto.Body.JSON(),
	}
}

// Prototypes is a set of Prototype, mapped by ID number
type Prototypes map[int]*Prototype

// Data returns a representation of these Prototypes as type fmt.Stringer
func (cards Prototypes) Data() []map[string]any {
	json := make([]map[string]any, len(cards))
	keys := make([]int, len(cards))

	var i int
	for k := range cards {
		keys[i] = k
		i++
	}
	sort.Ints(keys)
	var j int
	for _, id := range keys {
		json[j] = cards[id].Data()
		j++
	}
	return json
}
