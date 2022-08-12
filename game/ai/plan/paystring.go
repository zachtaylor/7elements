package plan

import (
	"github.com/zachtaylor/7elements/element"
	"github.com/zachtaylor/7elements/game/ai/view"
)

// PayString allows the AI to choose which elements to pay for an elements cost
func PayString(view view.T, costs element.Count) string {
	pay := make(element.Count)
	active := view.Self.T.Karma.Active()

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
