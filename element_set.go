package vii

import (
	"strings"

	"ztaylor.me/cast"
)

type ElementSet map[Element][]bool

func (set ElementSet) GetActive() ElementMap {
	m := ElementMap{}
	for el, stack := range set {
		for _, ok := range stack {
			if ok {
				m[el]++
			}
		}
	}
	return m
}

func (set ElementSet) Append(e Element) {
	set[e] = append(set[e], true)
}

func (set ElementSet) Reactivate() {
	for _, stack := range set {
		for i := 0; i < len(stack); i++ {
			stack[i] = true
		}
	}
}

func (set ElementSet) Deactivate(emp ElementMap) {
	emp = emp.Copy()

	for e, count := range emp {
		if e == ELEMnull {
			continue
		}

		for i, active := range set[e] {
			if active && count > 0 {
				set[e][i] = false
				count--
			}
		}

		if count != 0 {
			panic("game elements deactivate missing element: " + e.Char())
		}
	}

	for e := ELEMwhite; emp[ELEMnull] > 0; e++ {
		if e > ELEMblack {
			panic("game elements deactivate missing generic")
		}

		for i, active := range set[e] {
			if active {
				set[e][i] = false
				emp[ELEMnull]--
				if emp[ELEMnull] < 1 {
					break
				}
			}
		}
	}
}

func (set ElementSet) Json() Json {
	json := Json{}
	for e, stack := range set {
		json[cast.String(int(e))] = stack
	}
	return json
}

func (set ElementSet) String() string {
	sb := strings.Builder{}
	for e, stack := range set {
		for _, ok := range stack {
			if ok {
				sb.WriteString(e.Char())
			} else {
				sb.Write([]byte{e.Char()[0] + 32})
			}
		}
	}
	return sb.String()
}
