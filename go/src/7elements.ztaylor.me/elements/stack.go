package elements

type Stack map[Element]int

func (stack Stack) Copy() Stack {
	stack2 := Stack{}
	for element, count := range stack {
		stack2[element] = count
	}
	return stack2
}

func (stack Stack) Test(stack2 Stack) bool {
	stack2 = stack2.Copy()

	for element, count := range stack {
		for i := 0; i < count; i++ {
			if stack[element] > 0 {
				stack2[element]--
			} else {
				stack2[Null]--
			}
		}
	}

	for _, count := range stack2 {
		if count > 0 {
			return false
		}
	}

	return true
}

func (stack Stack) Remove(stack2 Stack) {
	stack2 = stack2.Copy()

	for element, count := range stack2 {
		if element == Null {
			continue
		}

		stack[element] -= count
	}

	for element := Null; stack2[Null] > 0; element++ {
		if element > Black {
			panic("RemoveElements missing count for any type")
		}

		if stack[element] < stack2[Null] {
			stack2[Null] -= stack[element]
			delete(stack, element)
		} else {
			stack[element] -= stack2[Null]
			delete(stack2, Null)
		}
	}
}
