package gameelements

import (
	"7elements.ztaylor.me/elements"
	"strconv"
)

type GameElements map[elements.Element][]bool

func New() GameElements {
	return GameElements{}
}

func (ge GameElements) Append(element elements.Element) {
	if set := ge[element]; set == nil {
		ge[element] = []bool{false}
	} else {
		ge[element] = append(set, false)
	}
}

func (ge GameElements) Copy() GameElements {
	ge2 := GameElements{}
	for element, set := range ge {
		set2 := make([]bool, len(set))
		for i := 0; i < len(set); i++ {
			set2[i] = set[i]
		}
		ge2[element] = set2
	}
	return ge
}

func (ge GameElements) Active() elements.Stack {
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

func (ge GameElements) Reactivate() {
	for _, set := range ge {
		for i := 0; i < len(set); i++ {
			set[i] = true
		}
	}
}

func (ge GameElements) Deactivate(stack elements.Stack) {
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

func (ge GameElements) TestElementCount(element elements.Element, count int) bool {
	if stack := ge.Active(); stack[element] < count {
		return false
	}
	return true
}

func (ge GameElements) TestStack(stack elements.Stack) bool {
	return ge.Active().Test(stack)
}

// func (ge GameElements)
