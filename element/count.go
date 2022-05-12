package element

import (
	"strconv"

	"taylz.io/types"
)

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
	nilcount := 0

	for element := Nil; element <= Black; element++ {
		if element == Nil {
			nilcount += c[element]
		} else if c[element] > c2[element] {
			nilcount += c[element] - c2[element]
		} else if c[element] < c2[element] {
			return false
		}
	}

	return c2[Nil] <= nilcount
}

// Remove will reduce the individual values of this count
func (c Count) Remove(c2 Count) error {
	if c2[Nil] > 0 {
		return types.NewErr("cannot remove nil element")
	}

	for element, count := range c2 {
		if element == Nil {
			continue
		}

		if c[element] < count {
			return types.NewErr("requires more element: " + element.String())
		}
		c[element] -= count
	}

	return nil
}

// JSON returns a representation of this count as type map[string]any
func (c Count) JSON() map[string]any {
	json := map[string]any{}
	for k, v := range c {
		json[strconv.FormatInt(int64(k), 10)] = v
	}
	return json
}

func (c Count) String() string {
	var buf types.StringBuilder
	buf.Grow(c.Total())
	for e := Nil; e <= Black; e++ {
		if c[e] < 1 {
			continue
		}
		char := e.ToChar()
		for i := 0; i < c[e]; i++ {
			buf.WriteByte(char)
		}
	}
	return buf.String()
}
