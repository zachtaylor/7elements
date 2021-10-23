package plan

import (
	"github.com/zachtaylor/7elements/element"
	"github.com/zachtaylor/7elements/game/seat"
)

// PayString allows the AI to choose which elements to pay for an elements cost
func PayString(seat *seat.T, costs element.Count) string {
	pay := make(element.Count)
	active := seat.Karma.Active()

	for el, count := range costs {
		if el == element.Nil {
			continue
		}
		active[el] -= count
		pay[el] += count
	}

	for i := 0; i < costs[element.Nil]; i++ {
		var el element.T
		var count int
		for e, c := range active {
			if c > count {
				el = e
				c = count
			}
		}
		active[el]--
		pay[el]++
	}

	return pay.String()
}
