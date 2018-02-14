package games

import (
	"elemen7s.com/elements"
	"strconv"
)

type Elements map[elements.Element][]bool

var elementsOnStringer = map[elements.Element]string{
	elements.White:  "W",
	elements.Blue:   "B",
	elements.Green:  "G",
	elements.Gold:   "Y",
	elements.Red:    "R",
	elements.Indigo: "I",
	elements.Black:  "D",
}
var elementsOffStringer = map[elements.Element]string{
	elements.White:  "w",
	elements.Blue:   "b",
	elements.Green:  "g",
	elements.Gold:   "y",
	elements.Red:    "r",
	elements.Indigo: "i",
	elements.Black:  "d",
}

// var elementsOffStringer = map[elements.Element]string{

// }

func NewElements() Elements {
	return Elements{}
}

func (ge Elements) Append(element elements.Element) {
	if set := ge[element]; set == nil {
		ge[element] = []bool{false}
	} else {
		ge[element] = append(set, false)
	}
}

func (ge Elements) Copy() Elements {
	ge2 := Elements{}
	for element, set := range ge {
		set2 := make([]bool, len(set))
		for i := 0; i < len(set); i++ {
			set2[i] = set[i]
		}
		ge2[element] = set2
	}
	return ge
}

func (ge Elements) Active() elements.Stack {
	stack := elements.Stack{}
	for element, set := range ge {
		for _, active := range set {
			if active {
				stack[element]++
			}
		}
	}
	return stack
}

func (ge Elements) Reactivate() {
	for _, set := range ge {
		for i := 0; i < len(set); i++ {
			set[i] = true
		}
	}
}

func (ge Elements) Deactivate(stack elements.Stack) {
	stack = stack.Copy()

	for element, count := range stack {
		if element == elements.Null {
			continue
		}

		for i, active := range ge.Copy()[element] {
			if active && count > 0 {
				ge[element][i] = false
				count--
			}
		}

		if count != 0 {
			panic("game elements deactivate missing element#" + strconv.FormatInt(int64(count), 10))
		}
	}

	for element := elements.Null; stack[elements.Null] > 0; element++ {
		if element > elements.Black {
			panic("game elements deactivate missing generic")
		}

		for i, active := range ge.Copy()[element] {
			if active && stack[elements.Null] > 0 {
				ge[element][i] = false
				stack[elements.Null]--
			}
		}
	}
}

func (ge Elements) TestElementCount(element elements.Element, count int) bool {
	if stack := ge.Active(); stack[element] < count {
		return false
	}
	return true
}

func (ge Elements) TestStack(stack elements.Stack) bool {
	return ge.Active().Test(stack)
}

func (ge Elements) String() string {
	s := ""
	for element, set := range ge {
		for i := 0; i < len(set); i++ {
			if set[i] {
				s += elementsOnStringer[element]
			} else {
				s += elementsOffStringer[element]
			}
		}
	}
	return s
}
