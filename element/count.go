package element

import "ztaylor.me/cast"

// Count is a count of elements
type Count map[T]int

// Copy returns a copy of this `Count`
func (c Count) Copy() Count {
	c2 := Count{}
	for element, count := range c {
		c2[element] = count
	}
	return c2
}

// Total returns the sum of all element counts
func (c Count) Total() (x int) {
	for _, i := range c {
		x += i
	}
	return
}

// Test returns whether another count is contained by this count
//
// this means that for every non-nil element `e`, `c2[e] <= c[e]`
//
// nil elements in `c2` are satisfied by remaining element counts of `c`
func (c Count) Test(c2 Count) bool {
	c2 = c2.Copy()

	for element, count := range c {
		for i := 0; i < count; i++ {
			if c2[element] > 0 {
				c2[element]--
			} else {
				c2[Nil]--
			}
		}
	}

	for _, count := range c2 {
		if count > 0 {
			return false
		}
	}

	return true
}

// Remove will reduce the individual values of this count
func (c Count) Remove(c2 Count) {
	c2 = c2.Copy()

	for element, count := range c2 {
		if element == Nil {
			continue
		}

		c[element] -= count
	}

	for element := Nil; c2[Nil] > 0; element++ {
		if element > Black {
			panic("RemoveElements missing count for any type")
		}

		if c[element] < c2[Nil] {
			c2[Nil] -= c[element]
			delete(c, element)
		} else {
			c[element] -= c2[Nil]
			delete(c2, Nil)
		}
	}
}

// JSON returns a representation of this count as type `cast.JSON`
func (c Count) JSON() cast.JSON {
	json := cast.JSON{}
	for k, v := range c {
		json[cast.StringI(int(k))] = v
	}
	return json
}
