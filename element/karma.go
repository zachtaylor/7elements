package element

import (
	"errors"
	"strconv"
)

// Karma is a set of active and inactive elements
type Karma map[T][]bool

// length returns the total of active and inactive elements
func (k Karma) length() (i int) {
	for _, stack := range k {
		i += len(stack)
	}
	return
}

// Active returns the Elements that are currently active
func (k Karma) Active() Count {
	c := Count{}
	for el, stack := range k {
		for _, ok := range stack {
			if ok {
				c[el]++
			}
		}
	}
	return c
}

// Append adds an Element to this Karma
func (k Karma) Append(e T, active bool) {
	k[e] = append(k[e], active)
}

// Reactivate resets all elements counts to on
func (k Karma) Reactivate() {
	for _, stack := range k {
		for i := 0; i < len(stack); i++ {
			stack[i] = true
		}
	}
}

var (
	ErrKarmaAmbiguous    = errors.New("karma ambiguous")
	ErrKarmaInsufficient = errors.New("karma insufficient")
)

// Deactivate sets the specified Elements in this Karma to unactive
func (k Karma) Deactivate(c Count) error {
	if c[Nil] > 0 {
		return ErrKarmaAmbiguous
	}
	for e, count := range c {
		for _, active := range k[e] {
			if count < 1 {
				break
			} else if active {
				count--
			}
		}
		if count > 0 {
			return ErrKarmaInsufficient
		}
	}

	for e, count := range c {
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

// String returns this Karma as a code representation
func (k Karma) String() string {
	length := k.length()
	code := make([]byte, length+2)
	code[0] = '{'
	code[length+1] = '}'
	i := 1
	for _, e := range Index {
		for _, ok := range k[e] {
			if ok {
				code[i] = e.ToChar()
			} else {
				code[i] = e.ToChar() + 32
			}
			i++
		}
	}
	return string(code)
}

// JSON returns a JSON representation of this Karma
func (k Karma) JSON() map[string]any {
	json := map[string]any{}
	for e, stack := range k {
		json[strconv.FormatInt(int64(e), 10)] = stack
	}
	return json
}
