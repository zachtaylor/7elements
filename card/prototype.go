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
	Name   string
	Text   string
	Type   Type
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

// Data returns a representation of this Prototype as type map[string]interface{}
func (proto *Prototype) Data() map[string]interface{} {
	return map[string]interface{}{
		"id":     proto.ID,
		"name":   proto.Name,
		"text":   proto.Text,
		"type":   proto.Type.String(),
		"powers": proto.Powers.Data(),
		"costs":  proto.Costs.JSON(),
		"body":   proto.Body.Data(),
	}
}

func (proto *Prototype) String() string {
	return `{` + strconv.FormatInt(int64(proto.ID), 10) + ` ` + proto.Name + `}`
}

// Prototypes is a set of Prototype, mapped by ID number
type Prototypes map[int]*Prototype

// Data returns a representation of these Prototypes as type fmt.Stringer
func (cards Prototypes) Data() []map[string]interface{} {
	json := make([]map[string]interface{}, len(cards))
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
