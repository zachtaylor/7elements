package vii

import "ztaylor.me/cast"

// ElementMap is used to convey a count of elements
type ElementMap map[Element]int

// Copy returns a copy of the map
func (m ElementMap) Copy() ElementMap {
	m2 := ElementMap{}
	for element, count := range m {
		m2[element] = count
	}
	return m2
}

// Count returns the total count
func (m ElementMap) Count() (x int) {
	for _, i := range m {
		x += i
	}
	return
}

// Test returns whether the set m2 is contained by set m
//
// this means every count of element in m2 is counted in m, and
// that m2[ELEMnil] is counted by non-duplicate counts in m
func (m ElementMap) Test(m2 ElementMap) bool {
	m2 = m2.Copy()

	for element, count := range m {
		for i := 0; i < count; i++ {
			if m2[element] > 0 {
				m2[element]--
			} else {
				m2[ELEMnil]--
			}
		}
	}

	for _, count := range m2 {
		if count > 0 {
			return false
		}
	}

	return true
}

func (m ElementMap) Remove(m2 ElementMap) {
	m2 = m2.Copy()

	for element, count := range m2 {
		if element == ELEMnil {
			continue
		}

		m[element] -= count
	}

	for element := ELEMnil; m2[ELEMnil] > 0; element++ {
		if element > ELEMblack {
			panic("RemoveElements missing count for any type")
		}

		if m[element] < m2[ELEMnil] {
			m2[ELEMnil] -= m[element]
			delete(m, element)
		} else {
			m[element] -= m2[ELEMnil]
			delete(m2, ELEMnil)
		}
	}
}

func (m ElementMap) JSON() cast.JSON {
	json := cast.JSON{}
	for k, v := range m {
		json[cast.StringI(int(k))] = v
	}
	return json
}
