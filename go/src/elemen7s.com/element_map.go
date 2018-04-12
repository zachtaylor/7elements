package vii

type ElementMap map[Element]int

func (m ElementMap) Copy() ElementMap {
	m2 := ElementMap{}
	for element, count := range m {
		m2[element] = count
	}
	return m2
}

func (m ElementMap) Test(m2 ElementMap) bool {
	m2 = m2.Copy()

	for element, count := range m {
		for i := 0; i < count; i++ {
			if m2[element] > 0 {
				m2[element]--
			} else {
				m2[ELEMnull]--
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
		if element == ELEMnull {
			continue
		}

		m[element] -= count
	}

	for element := ELEMnull; m2[ELEMnull] > 0; element++ {
		if element > ELEMblack {
			panic("RemoveElements missing count for any type")
		}

		if m[element] < m2[ELEMnull] {
			m2[ELEMnull] -= m[element]
			delete(m, element)
		} else {
			m[element] -= m2[ELEMnull]
			delete(m2, ELEMnull)
		}
	}
}
