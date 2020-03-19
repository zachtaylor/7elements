package card

import (
	"fmt"

	"github.com/zachtaylor/7elements/element"
	"ztaylor.me/cast"
)

// Set is a multiplicity of Cards, mapped by ID
type Set map[string]*T

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

// JSON returns a representation of a set of game cards
func (set Set) JSON() cast.JSON {
	data := cast.JSON{}
	for cid, c := range set {
		data[cid] = c.JSON()
	}
	return data
}
