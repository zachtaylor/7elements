package card

import (
	"fmt"

	"github.com/zachtaylor/7elements/element"
)

// Set is a multiplicity of Cards, mapped by ID
type Set map[string]*T

func (set Set) Has(id string) (ok bool) {
	_, ok = set[id]
	return
}

// Devotion returns the ElementMap describing the devotion of this group of cards
func (set Set) Devotion() element.Count {
	devo := element.Count{}
	for _, c := range set {
		for e, count := range c.Proto.Costs {
			devo[e] += count
		}
	}
	return devo
}

func (set Set) String() string {
	return `card.Set` + set.Print()
}

func (set Set) Print() string {
	collapse := make([]int, len(set))
	i := 0
	for _, c := range set {
		collapse[i] = c.Proto.ID
		i++
	}
	return fmt.Sprintf("%v", collapse)
}

func (set Set) Keys() []string {
	i, keys := 0, make([]string, len(set))
	for k := range set {
		keys[i] = k
		i++
	}
	return keys
}

// // JSON returns a representation of a set of game cards
// func (set Set) JSON() []string {
// 	data :=
// 	for cid, c := range set {
// 		data[cid] = c.Proto.ID
// 	}
// 	return data
// }
