package element

import "taylz.io/types"

// ParseCount reads a string for count code
func ParseCount(str string) (Count, error) {
	c := make(Count)

	for _, char := range str {
		element, _, err := Parse(byte(char))
		if err != nil {
			return nil, types.WrapErr(err, "parse count")
		}
		c[element]++
	}

	return c, nil
}

// ParseKarma reads a string for karma code
func ParseKarma(s string) (Karma, error) {
	k := Karma{}
	for i, c := range s {
		if i == 0 {
			if c != '{' {
				return nil, types.NewErr("element.ParseKarma: " + s)
			}
		} else if i == len(s)-1 {
			if c != '}' {
				return nil, types.NewErr("element.ParseKarma: " + s)
			}
		} else if el, active, err := Parse(byte(c)); err != nil {
			return nil, err
		} else {
			k.Append(el, active)
		}
	}
	return k, nil
}

// Parse decodes an element with activity status
//
// returns `T` encoded by the byte
//
// returns `bool` indicating activity state
//
// returns error if `c` cannot be parsed
func Parse(c byte) (T, bool, error) {
	switch c {
	case 'w':
		return White, false, nil
	case 'W':
		return White, true, nil
	case 'r':
		return Red, false, nil
	case 'R':
		return Red, true, nil
	case 'y':
		return Yellow, false, nil
	case 'Y':
		return Yellow, true, nil
	case 'g':
		return Green, false, nil
	case 'G':
		return Green, true, nil
	case 'b':
		return Blue, false, nil
	case 'B':
		return Blue, true, nil
	case 'v':
		return Violet, false, nil
	case 'V':
		return Violet, true, nil
	case 'a':
		return Black, false, nil
	case 'A':
		return Black, true, nil
	case 'x':
		return Nil, false, nil
	case 'X':
		return Nil, true, nil
	default:
		return Nil, false, types.NewErr("unknown element: " + string(c))
	}
}
