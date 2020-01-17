package vii

import "ztaylor.me/cast"

// Karma represents elements that may be active or reactivated later
type Karma map[Element][]bool

// ParseKarma reads a string for karma code
func ParseKarma(s string) (Karma, error) {
	k := Karma{}
	for i, c := range s {
		if i == 0 {
			if c != '{' {
				return nil, cast.NewError(nil, "vii.ParseKarma: ", s)
			}
		} else if i == len(s)-1 {
			if c != '}' {
				return nil, cast.NewError(nil, "vii.ParseKarma: ", s)
			}
		} else if el, active, err := ParseElement(byte(c)); err != nil {
			return nil, err
		} else {
			k.Append(el, active)
		}
	}
	return k, nil
}

// Active returns the ElementMap describing the Elements that are Active in this Karma
func (k Karma) Active() ElementMap {
	m := ElementMap{}
	for el, stack := range k {
		for _, ok := range stack {
			if ok {
				m[el]++
			}
		}
	}
	return m
}

// Append adds an Element to this Karma
func (k Karma) Append(e Element, active bool) {
	k[e] = append(k[e], active)
}

// Add adds a new Element to this Karma
func (k Karma) Add(e Element) {
	k.Append(e, false)
}

// Reactivate sets all Elements to Active in this Karma
func (k Karma) Reactivate() {
	for _, stack := range k {
		for i := 0; i < len(stack); i++ {
			stack[i] = true
		}
	}
}

// length returns the number of active and inactive elements in this Karma
func (k Karma) length() int {
	count := 0
	for _, stack := range k {
		count += len(stack)
	}
	return count
}

// Deactivate sets the specified Elements in this Karma to Deactive
//
// returns `ErrMissing` when ElementMap contains `ELEMnil`
//
// returns `ErrKarma` when ElementMap is not contained by this Karma
func (k Karma) Deactivate(m ElementMap) error {
	if m[ELEMnil] > 0 {
		return cast.NewError(nil, `vii.Karma.Deactivate`)
	}
	ok := true
	for e, count := range m {
		for _, active := range k[e] {
			if count > 0 {
				if active {
					count--
				}
			} else {
				break
			}
		}
		if count != 0 {
			ok = false
		}
	}
	if !ok {
		return ErrKarma
	}
	for e, count := range m {
		for i, active := range k[e] {
			if count > 0 {
				if active {
					k[e][i] = false
					count--
				}
			} else {
				break
			}
		}
	}
	return nil
}

// JSON returns a JSON representation of this Karma
func (k Karma) JSON() cast.JSON {
	json := cast.JSON{}
	for e, stack := range k {
		json[cast.StringI(int(e))] = stack
	}
	return json
}

// String returns this Karma as a code representation
func (k Karma) String() string {
	length := k.length()
	code := make([]byte, length+2)
	code[0] = '{'
	code[length+1] = '}'
	i := 1
	for _, e := range Elements {
		for _, ok := range k[e] {
			if ok {
				code[i] = e.Char()
			} else {
				code[i] = e.Char() + 32
			}
			i++
		}
	}
	return cast.StringBytes(code)
}
